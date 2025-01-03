package router

import (
	v1 "GoDistributedAppDevelop/demo09_unique_id/segments/api/v1"
	"github.com/gin-gonic/gin"
)

func InitSegmentsIdRouter(r *gin.Engine) {
	// 获取号码段路由组
	seg := r.Group("/segmentsid")
	{
		//获取号码段
		seg.GET("/:biz_type", v1.GetSegIds)
	}
}
