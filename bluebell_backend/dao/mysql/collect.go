package mysql

import (
	"bluebell_backend/models"
	"fmt"
	"go.uber.org/zap"
)

func CreateCollect(collect *models.Collect) (err error) {
	sqlStr := `insert into collect_link(
	collect_id, user_id, parent_id)
	values(?,?,?)`
	_, err = db.Exec(sqlStr, collect.CollectId, collect.UserId,
		collect.ParentId)
	if err != nil {
		zap.L().Error("insert post failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

func AddCollectOfPost(postId string) (err error) {
	_, err = db.Exec("UPDATE post SET collect_number = collect_number + 1 WHERE post_id = ?", postId)
	if err != nil {
		fmt.Println("更新数据错误:", err)
		return
	}

	return
}

func IsCollectByParentId(userId, postId string) (rest bool, err error) {
	sqlStr := `SELECT * FROM collect_link where user_id = ? and parent_id = ?;`
	var collect models.Collect
	err = db.Get(&collect, sqlStr, userId, postId)
	if err != nil {
		rest = false
		return
	}
	rest = true
	return
}

func DeleteCollectByCollectId(userId, postId string) (err error) {
	var star *models.Star
	sqlStr := `delete from collect_link where user_id = ? and parent_id = ?;`
	err = db.Select(&star, sqlStr, userId, postId)
	return
}
func MinusCollectOfPost(postId string) (err error) {
	_, err = db.Exec("UPDATE post SET collect_number = collect_number - 1 WHERE post_id = ?", postId)
	if err != nil {
		fmt.Println("更新数据错误:", err)
		return
	}

	return
}
