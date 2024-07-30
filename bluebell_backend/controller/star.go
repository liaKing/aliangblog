package controller

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/models"
	"bluebell_backend/pkg/snowflake"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func Star(c *gin.Context) {
	var star models.Star
	if err := c.ShouldBindJSON(&star); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	// 获取作者ID，当前请求的UserID
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}

	// 生成点赞ID
	starID, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("snowflake.GetID() failed", zap.Error(err))
		return
	}

	star.StartId = starID
	star.UserId = userID
	err = mysql.CreateStar(&star)
	if err != nil {
		zap.L().Error("logic.CreateStar failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	err = mysql.AddStarOfPost(star.ParentId)

	if err != nil {
		zap.L().Error("logic.AddStar failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)

}
func DeleteStarByUserIdAndPostId(c *gin.Context) {
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	postId, _ := c.GetQuery("postId")

	err = mysql.DeleteStarByUserIdAndPostId(userID, postId)

	err = mysql.MinusStarOfPost(postId)
	ResponseSuccess(c, postId)
}

func IsStarByParentId(c *gin.Context) {
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	postId, _ := c.GetQuery("postId")
	rest, _ := mysql.IsStarByParentId(userID, postId)

	ResponseSuccess(c, rest)
}
