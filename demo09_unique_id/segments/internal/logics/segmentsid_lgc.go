package logics

import (
	"GoDistributedAppDevelop/demo08_partition/internal/models"
	"GoDistributedAppDevelop/demo08_partition/pkg/mysqldb"
	"github.com/jiebozeng/golangutils/convert"
)

type User_lgc struct {
}

// 逻辑实现
func (u *User_lgc) GetUserByUid(userId int64) (*models.User, error) {
	user := &models.User{}
	tableName := "user_" + convert.ToString(userId%10+1)
	query := mysqldb.Mysql.Table(tableName).Model(user)
	//查询单个用户
	query = query.Where("user_id = ?", userId)
	err := query.Find(&user).Limit(1).Error
	return user, err
}
