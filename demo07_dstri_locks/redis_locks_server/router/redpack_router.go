package router

import (
	v1 "GoDistributedAppDevelop/demo07_dstri_locks/redis_locks_server/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRedpackRouter(r *gin.Engine) {
	// 红包路由组
	redpack := r.Group("/red")
	{
		//获取单个用户信息
		redpack.GET("/:user_id/:redpack_id", v1.GetRedpack)
	}
}
