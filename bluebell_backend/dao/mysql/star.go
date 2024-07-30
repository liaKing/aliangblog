package mysql

import (
	"bluebell_backend/models"
	"fmt"
	"go.uber.org/zap"
)

func CreateStar(star *models.Star) (err error) {
	sqlStr := `insert into star(
	star_id, user_id, parent_id, target_id)
	values(?,?,?,?)`
	_, err = db.Exec(sqlStr, star.StartId, star.UserId,
		star.ParentId, star.ParentId)
	if err != nil {
		zap.L().Error("insert post failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

func AddStarOfPost(postId string) (err error) {
	_, err = db.Exec("UPDATE post SET star_number = star_number + 1 WHERE post_id = ?", postId)
	if err != nil {
		fmt.Println("更新数据错误:", err)
		return
	}

	return
}

func DeleteStarByUserIdAndPostId(userId, postId string) (err error) {
	var star *models.Star
	sqlStr := `delete from star where user_id = ? and parent_id = ?;`
	err = db.Select(&star, sqlStr, userId, postId)
	return
}

func MinusStarOfPost(postId string) (err error) {
	_, err = db.Exec("UPDATE post SET star_number = star_number - 1 WHERE post_id = ?", postId)
	if err != nil {
		fmt.Println("更新数据错误:", err)
		return
	}

	return
}

func IsStarByParentId(userId string, postId string) (rest bool, err error) {
	sqlStr := `SELECT * FROM star where user_id = ? and parent_id = ?;`
	var star models.Star
	err = db.Get(&star, sqlStr, userId, postId)
	if err != nil {
		rest = false
		return
	}
	rest = true
	return
}
