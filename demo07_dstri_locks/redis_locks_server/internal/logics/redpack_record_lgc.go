package logics

import (
	"GoDistributedAppDevelop/demo07_dstri_locks/redis_locks_server/internal/models"
	"GoDistributedAppDevelop/demo07_dstri_locks/redis_locks_server/pkg/mysqldb"
)

type RedpackRecordLgc struct {
}

// 插入红包记录
func (r *RedpackRecordLgc) InsertRedPackRecord(record *models.RedPackRecord) error {
	return mysqldb.Mysql.Create(record).Error
}
