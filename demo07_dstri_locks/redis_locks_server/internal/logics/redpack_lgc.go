package logics

import (
	"GoDistributedAppDevelop/demo07_dstri_locks/redis_locks_server/internal/models"
	"GoDistributedAppDevelop/demo07_dstri_locks/redis_locks_server/pkg/mysqldb"
	"GoDistributedAppDevelop/demo07_dstri_locks/redis_locks_server/pkg/redisdb"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type RedPackLgc struct {
}

// 抢红包 返回红包id 抢到的金额
// 没抢到的话 返回的红包id为-1
func (r *RedPackLgc) GradRedPack(userId int64, redpackId int64) (redId int64, amount int64, err error) {
	ctx := context.Background()
	lockKey := fmt.Sprintf("lock:redpack:%d", redpackId)
	lockTimeout := 5 * time.Second

	// 尝试获取分布式锁
	lockAcquired, err := redisdb.RedisDb.SetNX(ctx, lockKey, "1", lockTimeout).Result()
	if err != nil {
		return -1, 0, err
	}
	if !lockAcquired {
		// 未能获取锁，可能红包正在被其他用户处理
		return -1, 0, errors.New("红包正在被其他用户处理，请稍后再试")
	}
	defer redisdb.RedisDb.Del(ctx, lockKey) // 确保锁被释放

	// 获取红包信息
	redpackJSON, err := redisdb.RedisDb.Get(ctx, fmt.Sprintf("redpack:%d", redpackId)).Result()
	if err != nil {
		if err == redis.Nil {
			return -1, 0, errors.New("红包不存在或已过期")
		}
		return -1, 0, err
	}

	redpack := &models.Redpack{}
	if err := json.Unmarshal([]byte(redpackJSON), redpack); err != nil {
		return -1, 0, err
	}

	// 检查红包状态
	if redpack.Status != 1 || redpack.Num <= redpack.ProNum {
		return -1, 0, errors.New("红包已抢完或无效")
	}

	// 计算红包金额（这里简化为总金额/20，可以根据需求调整）
	if redpack.Amount <= 5 {
		amount = redpack.Amount
	} else {
		amount = redpack.Amount / 20
		if amount <= 0 {
			amount = 1 // 最小金额为1
		}
	}

	// 更新已抢数量和剩余金额
	redpack.ProNum++
	redpack.Amount -= amount
	if redpack.Amount == 0 {
		redpack.Status = 5 //已领完
		r.UpdateRedPackStatus(redpackId, 5)
	}
	updatedRedpackJSON, err := json.Marshal(redpack)
	if err != nil {
		return -1, 0, err
	}

	if _, err := redisdb.RedisDb.Set(ctx, fmt.Sprintf("redpack:%d", redpackId), updatedRedpackJSON, 0).Result(); err != nil {
		return -1, 0, err
	}
	//抢到红包 插入记录
	userLgc := &User_lgc{}
	user, _ := userLgc.GetUserByUid(userId)
	record := &models.RedPackRecord{
		UserId:    int64(user.ID),
		UserName:  user.UserName,
		RedPackId: redpackId,
		Amount:    amount,
	}
	record.CreatedAt = time.Now()
	record.UpdatedAt = time.Now()
	redpackRecordLgc := &RedpackRecordLgc{}
	redpackRecordLgc.InsertRedPackRecord(record)
	return redpackId, amount, nil
}

// 获取单个红包信息
func (r *RedPackLgc) GetRedPackById(redpackId int64) (*models.Redpack, error) {
	redpack := &models.Redpack{}
	query := mysqldb.Mysql.Model(redpack)
	//查询红包信息
	query = query.Where("id = ? and status=? ", redpackId, 1)
	err := query.Find(&redpack).Order("id desc").Limit(1).Error
	return redpack, err
}

// 更新领取数量
func (r *RedPackLgc) UpdateRedPackProNum(redpackId int64) error {
	query := mysqldb.Mysql.Model(&models.Redpack{})
	return query.Where("id = ?", redpackId).Update("prop_num", gorm.Expr("prop_num + 1")).Error
}

// 更新状态
func (r *RedPackLgc) UpdateRedPackStatus(redpackId int64, status int) error {
	query := mysqldb.Mysql.Model(&models.Redpack{})
	return query.Where("id = ?", redpackId).Update("status", status).Error
}
