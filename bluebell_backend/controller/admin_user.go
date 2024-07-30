package controller

import (
	"bluebell_backend/dao/mysql"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUserList(c *gin.Context) {
	pageStr, ok := c.GetQuery("page")
	if !ok {
		pageStr = "1"
	}
	pageNum, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		pageNum = 1
	}
	users, err := mysql.GetUserList(pageNum)

	fmt.Println(len(users))
	ResponseSuccess(c, users)
}

func DeleteUserByUserId(c *gin.Context) {
	userId, _ := c.GetQuery("userId")

	mysql.DeleteUserByUserId(userId)

	ResponseSuccess(c, userId)
}

func RecoverUserByUserName(c *gin.Context) {
	userName, _ := c.GetQuery("userName")

	mysql.RecoverUserByUserName(userName)

	ResponseSuccess(c, userName)
}
