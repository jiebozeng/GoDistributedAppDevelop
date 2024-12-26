package mysqldb

import (
	"GoDistributedAppDevelop/demo08_partition/config"
	"GoDistributedAppDevelop/demo08_partition/pkg/logs"
	"database/sql"
	//"github.com/go-sql-driver/mysql"

	//"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var db *sql.DB

var Mysql *gorm.DB

func InitMysql() {
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.CONFIG.Mysql.User,
		config.CONFIG.Mysql.Password,
		config.CONFIG.Mysql.Host,
		config.CONFIG.Mysql.Port,
		config.CONFIG.Mysql.Database,
	)
	c := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	//if config.App.Env == config.ConfigEnvDebug {
	//	c.Logger = logger.Default.LogMode(logger.Info)
	//}

	Mysql, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	}), c)
	if err != nil {
		logs.ZapLogger.Error("连接mysql数据库失败：" + err.Error())
	}
	// 设置连接池
	var MysqlDB *sql.DB
	MysqlDB, err = Mysql.DB()
	if err != nil {
		return
	}
	MysqlDB.SetConnMaxLifetime(5 * time.Minute)
	MysqlDB.SetConnMaxIdleTime(2 * time.Minute)
	MysqlDB.SetMaxIdleConns(config.CONFIG.Mysql.MaxIdleConns)
	MysqlDB.SetMaxOpenConns(config.CONFIG.Mysql.MaxOpenConns)
}

func CloseMysql() {
	MysqlDB, err := Mysql.DB()
	if err != nil {
		logs.ZapLogger.Error("关闭mysql数据库失败：" + err.Error())
		return
	}
	err = MysqlDB.Close()
	if err != nil {
		logs.ZapLogger.Error("关闭mysql数据库失败：" + err.Error())
		return
	}
}
