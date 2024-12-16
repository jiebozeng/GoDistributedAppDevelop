package router

import (
	v1 "GoDistributedAppDevelop/demo04_microser_comm_restful/user_server/api/v1"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.GET("/list", v1.UserList)
		user.GET("/info/:user_id", v1.GetUser)
	}
}
