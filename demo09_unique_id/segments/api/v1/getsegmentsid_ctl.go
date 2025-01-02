package v1

import (
	"GoDistributedAppDevelop/demo09_unique_id/segments/internal/logics"
	"github.com/gin-gonic/gin"
	"github.com/jiebozeng/golangutils/convert"
)

func GetSegIds(c *gin.Context) {
	bizType := c.Param("biz_type")
	version := c.Param("version")
	userLgc := &logics.User_lgc{}
	user, err := userLgc.GetUserByUid(convert.ToInt64(userId))
	if err != nil {
		c.JSON(-1, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(0, gin.H{
		"user": user,
	})
}
