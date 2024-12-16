package v1

import (
	"GoDistributedAppDevelop/demo04_microser_comm_restful/user_server/internal/logics"
	"github.com/gin-gonic/gin"
	"github.com/jiebozeng/golangutils/convert"
)

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
