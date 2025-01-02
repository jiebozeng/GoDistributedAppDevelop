package router

import (
	v1 "GoDistributedAppDevelop/demo09_unique_id/segments/api/v1"
	"github.com/gin-gonic/gin"
)

func InitSegmentsIdRouter(r *gin.Engine) {
	// 用户路由组
	user := r.Group("/segmentsid")
	{
		//获取号码段
		user.GET("/:biz_type/:version", v1.GetSegIds)
	}
}
