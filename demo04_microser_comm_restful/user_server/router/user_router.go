package router

import (
	v1 "GoDistributedAppDevelop/demo04_microser_comm_restful/user_server/api/v1"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.Engine) {
	// 用户路由组
	user := r.Group("/user")
	{
		//获取用户列表，分页
		user.GET("/list", v1.UserList)
		//获取单个用户信息
		user.GET("/info/:user_id", v1.GetUser)
	}
}
