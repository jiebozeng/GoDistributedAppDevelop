package sql

import (
	"GoDistributedAppDevelop/demo08_partition/config"
	"GoDistributedAppDevelop/demo08_partition/internal/models"
	"GoDistributedAppDevelop/demo08_partition/pkg/logs"
	"GoDistributedAppDevelop/demo08_partition/pkg/mysqldb"
	"github.com/jiebozeng/golangutils/convert"
	"github.com/jiebozeng/golangutils/timeutils"
	"testing"
)

func TestInitPartitionUsers(t *testing.T) {
	config.Init()
	//初始化log
	logs.InitLogger(config.CONFIG.Logger.LogTypes, config.CONFIG.Logger.Dir, logs.LogEnvType(config.CONFIG.System.Mode), config.CONFIG.Logger.LogMaxAge)
	//初始化数据库链接
	mysqldb.InitMysql()
	defer mysqldb.CloseMysql()
	defer logs.CloseLog()
	nowTime := timeutils.GetNowTime()
	for i := int64(1); i < 100001; i++ {
		user := &models.User{
			UserId:     i,
			UserName:   "user" + convert.ToString(i),
			UserMobile: convert.ToString(13800000000 + i),
			UserEmail:  "user_" + convert.ToString(i) + "@qq.com",
			UserPwd:    "E10ADC3949BA59ABBE56E057F20F883E",
			CreatedAt:  nowTime,
			UpdatedAt:  nowTime,
		}
		tableName := "user_" + convert.ToString(i%10+1)
		ret := mysqldb.Mysql.Table(tableName).Create(user)
		if ret.Error != nil {
			t.Error(ret.Error)
		} else {
			//fmt.Println("插入用户成功:", user)
		}
	}
}
