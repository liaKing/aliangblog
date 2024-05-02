package controller

import "github.com/gin-gonic/gin"

// 封装的从上下文获取userId的方法
func getCurrentUserID(c *gin.Context) (userID uint64, err error) {
	_userID, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = _userID.(uint64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
