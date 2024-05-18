package mysql

import (
	"bluebell_backend/models"
	"fmt"
)

func GetUserList(page int64) (users []*models.UserInfo, err error) {
	sqlStr := `select user_id,username,password,create_time,del_flg from user LIMIT 10  OFFSET ?`
	offset := (page - 1) * 10

	users = make([]*models.UserInfo, 0, 10)
	err = db.Select(&users, sqlStr, offset)

	return
}

func DeleteUserByUserId(userId string) (err error) {
	_, err = db.Exec("UPDATE user SET del_flg = 1 WHERE user_id = ?", userId)
	if err != nil {
		fmt.Println("更新数据错误:", err)
		return
	}

	return
}

func RecoverUserByUserName(userName string) (err error) {
	_, err = db.Exec("UPDATE user SET del_flg = 0 WHERE username = ?", userName)
	if err != nil {
		fmt.Println("更新数据错误:", err)
		return
	}

	return
}
