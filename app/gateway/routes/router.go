package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"grpc-todolist/app/gateway/internal/handler"
	"grpc-todolist/app/gateway/middleware"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.ErrorMiddleware())
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("my-session", store))
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})
		// 用户服务
		v1.POST("/user/register", handler.UserRegister)
		v1.POST("/user/login", handler.UserLogin)

		// 需要登录保护
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			// 任务模块
			authed.GET("task", handler.GetTaskList)
			authed.GET("task/search", handler.SearchTask)
			authed.POST("task", handler.CreateTask)
			authed.PUT("task", handler.UpdateTask)
			authed.PUT("tasks", handler.UpdateAllTask)
			authed.DELETE("task", handler.DeleteTask)
			authed.DELETE("tasks", handler.DeleteAllTask)
		}
	}
	return ginRouter
}
