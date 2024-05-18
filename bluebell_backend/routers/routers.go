package routers

import (
	"bluebell_backend/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	//r := gin.New()
	//r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//gin.Default()是Gin框架提供的一个方法，用于创建一个具有默认中间件的Gin实例。
	//默认中间件包括日志记录、恢复恐慌和错误处理等功能，可以帮助简化服务器端的开发。
	//通过调用gin.Default()方法，可以快速创建一个具有常用中间件的Gin实例，以便构建和运行HTTP服务
	r := gin.Default()
	v1 := r.Group("/api/v1")
	v1.POST("/login", controller.LoginHandler)               //登录
	v1.POST("/signup", controller.SignUpHandler)             //注册
	v1.GET("/refresh_token", controller.RefreshTokenHandler) //刷新token

	v1.Use(controller.JWTAuthMiddleware()) //添加检查中间件
	{
		v1.GET("/community", controller.CommunityHandler)           //获取全部社区信息列表
		v1.GET("/community/:id", controller.CommunityDetailHandler) //根据id获取社区详情

		v1.POST("/post", controller.CreatePostHandler)                    //创建帖子
		v1.GET("/post/:id", controller.PostDetailHandler)                 //获取贴子详情
		v1.GET("/post", controller.PostListHandler)                       //分页获取帖子列表
		v1.GET("/posts/getPostByTitle", controller.GetPostListByTitle)    //模糊分页获取帖子列表
		v1.GET("/posts/getPostAudit", controller.GetPostAuditListByTitle) //模糊分页获取帖子列表
		v1.PUT("/posts/PostAuditSuccess", controller.PostAuditSuccess)    //博客审核通过
		v1.PUT("/posts/PostAuditFail", controller.PostAuditFail)          //博客审核通过

		//管理员用户管理
		v1.GET("/admin/userList", controller.GetUserList)
		v1.PUT("/admin/deleteUser", controller.DeleteUserByUserId)
		v1.PUT("/admin/recoverUser", controller.RecoverUserByUserName)

		v1.GET("/posts2", controller.PostList2Handler) //获取前两个帖子（包含用户name和communityName）

		v1.POST("/vote", controller.VoteHandler) //投票

		v1.POST("/comment", controller.CommentHandler)    //评论
		v1.GET("/comment", controller.CommentListHandler) //根据id获取评论列表

		//点赞
		v1.POST("/star", controller.Star)                          //评论
		v1.DELETE("/star", controller.DeleteStarByUserIdAndPostId) //根据id获取评论列表
		v1.GET("/star", controller.IsStarByParentId)
		//收藏
		v1.POST("/collect", controller.Collect)                    //评论
		v1.DELETE("/collect", controller.DeleteCollectByCollectId) //根据id获取评论列表
		v1.GET("/collect", controller.IsCollectByParentId)
		v1.GET("/collectList", controller.GetCollectListByuserId)

		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
