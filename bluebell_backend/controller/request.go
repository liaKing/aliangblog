package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// 封装的从上下文获取userId的方法
func getCurrentUserID(c *gin.Context) (userID string, err error) {
	_userID, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID64, ok := _userID.(uint64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID = strconv.FormatUint(userID64, 10)
	return
}
