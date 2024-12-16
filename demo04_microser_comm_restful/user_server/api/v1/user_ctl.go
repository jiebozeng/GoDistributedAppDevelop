package v1

import (
	"GoDistributedAppDevelop/demo04_microser_comm_restful/user_server/internal/logics"
	"github.com/gin-gonic/gin"
	"github.com/jiebozeng/golangutils/convert"
)

// @Tags	APP接口/获取单个用户信息
// @Summary	获取单个用户信息
// @Param   param query int "user_id"
// @Router  /info/{user_id} [get]
func GetUser(c *gin.Context) {
	userId := c.Param("user_id")
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

// @Tags	APP接口/获取用户列表
// @Summary	获取用户列表 分页
// @Param   param query int "page_num" "page_size"
// @Router  /list [get]
func UserList(c *gin.Context) {
	pageNum := c.Param("page_num")
	pageSize := c.Param("page_size")
	userLgc := &logics.User_lgc{}
	userList, err := userLgc.GetUserList(convert.ToInt64(pageNum), convert.ToInt64(pageSize))
	if err != nil {
		c.JSON(-1, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(0, gin.H{
		"user": userList,
	})
}
