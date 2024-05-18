package controller

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/models"
	"bluebell_backend/pkg/snowflake"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Collect(c *gin.Context) {
	var collect models.Collect
	if err := c.ShouldBindJSON(&collect); err != nil {
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
	collectId, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("snowflake.GetID() failed", zap.Error(err))
		return
	}

	collect.CollectId = collectId
	collect.UserId = userID
	err = mysql.CreateCollect(&collect)
	if err != nil {
		zap.L().Error("logic.CreateStar failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	err = mysql.AddCollectOfPost(collect.ParentId)

	if err != nil {
		zap.L().Error("logic.AddStar failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, collect)
}

func DeleteCollectByCollectId(c *gin.Context) {
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	postId, _ := c.GetQuery("postId")

	err = mysql.DeleteCollectByCollectId(userID, postId)

	err = mysql.MinusCollectOfPost(postId)
	ResponseSuccess(c, postId)
}

func IsCollectByParentId(c *gin.Context) {
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	postId, _ := c.GetQuery("postId")
	rest, _ := mysql.IsCollectByParentId(userID, postId)
	ResponseSuccess(c, rest)
}
func GetCollectListByuserId(c *gin.Context) {

}
