package v1

import (
	"GoDistributedAppDevelop/demo07_dstri_locks/redis_locks_server/internal/logics"
	"github.com/gin-gonic/gin"
	"github.com/jiebozeng/golangutils/convert"
)

// @Tags	APP接口/抢红包
// @Summary	抢红包
// @Param   param query int64 user_id int64 redpack_id
// @Router  /{user_id}/{redpack_id} [get]
func GetRedpack(c *gin.Context) {
	userId := c.Param("user_id")
	redpackId := c.Param("redpack_id")
	redpackLgc := &logics.RedPackLgc{}
	redId, amount, err := redpackLgc.GradRedPack(convert.ToInt64(userId), convert.ToInt64(redpackId))
	if err != nil {
		c.JSON(-1, gin.H{
			"redpackId": -1,
			"amount":    0,
			"err":       err.Error(),
		})
		return
	}
	c.JSON(0, gin.H{
		"redpackId": redId,
		"amount":    amount,
		"err":       nil,
	})
}
