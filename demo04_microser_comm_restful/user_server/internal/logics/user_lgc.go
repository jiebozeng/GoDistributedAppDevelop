package logics

import (
	"GoDistributedAppDevelop/demo04_microser_comm_restful/user_server/internal/models"
	"GoDistributedAppDevelop/demo04_microser_comm_restful/user_server/pkg/mysqldb"
)

type User_lgc struct {
}

func (u *User_lgc) GetUserByUid(userId int64) (*models.User, error) {
	user := &models.User{}
	query := mysqldb.Mysql.Model(user)
	query = query.Where("user_id = ?", userId)
	err := query.Find(&user).Limit(1).Error
	return user, err
}

func (u *User_lgc) GetUserList(pageNum, pageSize int64) (userList []*models.User, err error) {
	//分页查询
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 10 {
		pageSize = 10
	}
	if pageSize > 1000 {
		pageSize = 1000
	}

	user := &models.User{}
	query := mysqldb.Mysql.Model(user)
	err = query.Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&userList).Error
	return userList, err
}
