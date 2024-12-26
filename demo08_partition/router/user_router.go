package router

import (
	v1 "GoDistributedAppDevelop/demo08_partition/api/v1"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.Engine) {
	// 用户路由组
	user := r.Group("/user")
	{
		//获取单个用户信息
		user.GET("/info/:user_id", v1.GetUser)
	}
}
