package logic

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/dao/redis"
	"bluebell_backend/models"
	"bluebell_backend/pkg/snowflake"
	"fmt"
	"go.uber.org/zap"
)

func CreatePost(post *models.Post) (err error) {
	// 生成帖子ID
	postID, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("snowflake.GetID() failed", zap.Error(err))
		return
	}

	post.PostID = postID
	// 创建帖子
	if err := mysql.CreatePost(post); err != nil {
		zap.L().Error("mysql.CreatePost(&post) failed", zap.Error(err))
		return err
	}
	//fmt.Sprint（post.CommunityID）"用于将帖子的CommunityID变量的值转换为字符串格式。这通常用于Go编程中将字符串连接或以格式化方式打印变量
	community, err := mysql.GetCommunityNameByID(fmt.Sprint(post.CommunityID))
	if err != nil {
		zap.L().Error("mysql.GetCommunityNameByID failed", zap.Error(err))
		return err
	}
	//可能是缓存操作
	if err := redis.CreatePost(
		fmt.Sprint(post.PostID),
		fmt.Sprint(post.AuthorId),
		post.Title,
		TruncateByWords(post.Content, 120),
		community.CommunityName); err != nil {
		zap.L().Error("redis.CreatePost failed", zap.Error(err))
		return err
	}
	return

}

func GetPost(postID string) (post *models.ApiPostDetail, err error) {
	post, err = mysql.GetPostByID(postID)
	if err != nil {
		post, _ = mysql.GetPostAuditByID(postID)
	}
	user, err := mysql.GetUserByID(fmt.Sprint(post.AuthorId))
	if err != nil {
		zap.L().Error("mysql.GetUserByID() failed", zap.String("author_id", fmt.Sprint(post.AuthorId)), zap.Error(err))
		return
	}
	post.AuthorName = user.UserName
	community, err := mysql.GetCommunityByID(fmt.Sprint(post.CommunityID))
	if err != nil {
		zap.L().Error("mysql.GetCommunityByID() failed", zap.String("community_id", fmt.Sprint(post.CommunityID)), zap.Error(err))
		return
	}
	post.CommunityName = community.CommunityName
	return post, nil
}

func GetPostList2() (data []*models.ApiPostDetail, err error) {
	postList, err := mysql.GetPostList()
	if err != nil {
		fmt.Println(err)
		return
	}
	data = make([]*models.ApiPostDetail, 0, len(postList))
	for _, post := range postList {
		user, err := mysql.GetUserByID(fmt.Sprint(post.AuthorId))
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed", zap.String("author_id", fmt.Sprint(post.AuthorId)), zap.Error(err))
			continue
		}
		post.AuthorName = user.UserName
		community, err := mysql.GetCommunityByID(fmt.Sprint(post.CommunityID))
		if err != nil {
			zap.L().Error("mysql.GetCommunityByID() failed", zap.String("community_id", fmt.Sprint(post.CommunityID)), zap.Error(err))
			continue
		}
		post.CommunityName = community.CommunityName
		data = append(data, post)
	}
	return
}
