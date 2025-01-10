package produce_redpack

import (
	"GoDistributedAppDevelop/demo07_dstri_locks/redis_locks_server/internal/models"
	"GoDistributedAppDevelop/demo07_dstri_locks/redis_locks_server/pkg/logs"
	"GoDistributedAppDevelop/demo07_dstri_locks/redis_locks_server/pkg/mysqldb"
	"GoDistributedAppDevelop/demo07_dstri_locks/redis_locks_server/pkg/redisdb"
	"context"
	"encoding/json"
	"github.com/jiebozeng/golangutils/convert"
	"github.com/jiebozeng/golangutils/timer"
	"time"
)

// 生成红包
func ProduceRedpack() {
	tm := timer.NewTimer(time.Second)
	tm.Start()

	tm.AddTimer(time.Second*5, 10, func(t *timer.Timer) {
		//生成红包 插入数据库 并写入Redis
		redpack := &models.Redpack{
			Amount:    100,
			Num:       10,
			ValidTime: -1,
			Status:    1,
			ProNum:    0,
		}
		redpack.CreatedAt = time.Now()
		redpack.UpdatedAt = time.Now()
		err := mysqldb.Mysql.Create(redpack).Error
		if err != nil {
			logs.ZapLogger.Error("生成红包失败" + err.Error())
			panic(err)
		}
		logs.ZapLogger.Info("生成红包成功 id=>" + convert.ToString(redpack.ID))
		//红包信息保存到Redis
		redpackMarshal, _ := json.Marshal(redpack)
		redisdb.RedisDb.Set(context.TODO(), "redpack:"+convert.ToString(redpack.ID), redpackMarshal, 0)
	})
}
