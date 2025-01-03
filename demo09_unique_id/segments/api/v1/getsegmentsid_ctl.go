package v1

import (
	"GoDistributedAppDevelop/demo09_unique_id/segments/internal/logics"
	"github.com/gin-gonic/gin"
	"github.com/jiebozeng/golangutils/convert"
)

// @Tags	APP接口/基于数据库获取号码段
// @Summary	基于数据库获取号码段 [minId,maxId] 都可用
// @Param   param query int "biz_type"
// @Router  /segmentsid/biz_type [get]
func GetSegIds(c *gin.Context) {
	bizType := c.Param("biz_type")
	segLgc := &logics.Segmentsid_lgc{}
	minId, maxId, err := segLgc.GetSegmentsIds(convert.ToInt64(bizType))
	if err != nil {
		c.JSON(-1, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(0, gin.H{
		"min_id":   minId,
		"max_id":   maxId,
		"biz_type": bizType,
	})
}
