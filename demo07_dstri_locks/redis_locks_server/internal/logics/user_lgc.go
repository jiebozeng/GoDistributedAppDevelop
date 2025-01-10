package logics

import (
	"GoDistributedAppDevelop/demo07_dstri_locks/redis_locks_server/internal/models"
	"GoDistributedAppDevelop/demo07_dstri_locks/redis_locks_server/pkg/mysqldb"
)

type User_lgc struct {
}

// 逻辑实现
func (u *User_lgc) GetUserByUid(userId int64) (*models.User, error) {
	user := &models.User{}
	query := mysqldb.Mysql.Model(user)
	//查询单个用户
	query = query.Where("user_id = ?", userId)
	err := query.Find(&user).Limit(1).Error
	return user, err
}

// 逻辑实现，分页查询
func (u *User_lgc) GetUserList(pageNum, pageSize int64) (userList []*models.User, err error) {
	//分页查询
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 2 {
		pageSize = 2
	}
	if pageSize > 1000 {
		pageSize = 1000
	}

	user := &models.User{}
	query := mysqldb.Mysql.Model(user)
	err = query.Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&userList).Error
	return userList, err
}
