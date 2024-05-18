package mysql

import (
	"bluebell_backend/models"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

// CreatePost 创建帖子
func CreatePost(post *models.Post) (err error) {
	sqlStr := `insert into post_audit(
	post_id, title, content, author_id, community_id)
	values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, post.PostID, post.Title,
		post.Content, post.AuthorId, post.CommunityID)
	if err != nil {
		zap.L().Error("insert post failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

// GetPostByID
func GetPostByID(idStr string) (post *models.ApiPostDetail, err error) {
	post = new(models.ApiPostDetail)
	sqlStr := `select post_id, title, content, author_id, community_id, created_time,star_number,comment_number,collect_number
	from post
	where post_id = ?`
	err = db.Get(post, sqlStr, idStr)
	if err == sql.ErrNoRows {
		err = ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query post failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	return
}

func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	where post_id in (?)`
	// 动态填充id
	query, args, err := sqlx.In(sqlStr, ids)
	if err != nil {
		return
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
func GetPostListByTitle(title string, page int64) (postList []*models.Post, err error) {

	sqlStr := `SELECT * FROM post WHERE title LIKE ? LIMIT 10 OFFSET ?;`
	offset := (page - 1) * 10

	postList = make([]*models.Post, 0, 10)
	err = db.Select(&postList, sqlStr, "%"+title+"%", offset)

	return
}

func GetPostList() (posts []*models.ApiPostDetail, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, created_time
	from post
	limit 10
	`
	posts = make([]*models.ApiPostDetail, 0, 10)
	err = db.Select(&posts, sqlStr)
	return
}

func GetPostAuditListByTitle(page int64) (postList []*models.Post, err error) {
	sqlStr := `SELECT * FROM post_audit LIMIT 10 OFFSET ?;`
	offset := (page - 1) * 10

	postList = make([]*models.Post, 0, 10)
	err = db.Select(&postList, sqlStr, offset)
	return
}

func PostAuditSuccess(postId string) (post *models.Post, err error) {
	sqlStr := `select * from post_audit where post_id = ?;`
	post = new(models.Post)
	err = db.Get(post, sqlStr, postId)
	//删除审核表中数据
	sqlStr = `delete from post_audit where post_id = ?;`
	err = db.Select(&post, sqlStr, postId)
	//将数据添加到博客中
	sqlStr = `insert into post(
	post_id, title, content, author_id, community_id,status)
	values(?,?,?,?,?,?);`
	_, err = db.Exec(sqlStr, post.PostID, post.Title,
		post.Content, post.AuthorId, post.CommunityID, post.Status)

	if err != nil {
		zap.L().Error("insert post failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}

	return
}

func PostAuditFail(postId string) (err error) {
	_, err = db.Exec("UPDATE post_audit SET published = 1 WHERE post_id = ?", postId)
	if err != nil {
		fmt.Println("更新数据错误:", err)
		return
	}

	return
}
