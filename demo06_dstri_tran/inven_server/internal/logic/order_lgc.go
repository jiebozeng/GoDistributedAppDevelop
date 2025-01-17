package logic

import (
	"GoDistributedAppDevelop/demo06_dstri_tran/inven_server/internal/model"
	"GoDistributedAppDevelop/demo06_dstri_tran/inven_server/pkg/logs"
	"GoDistributedAppDevelop/demo06_dstri_tran/inven_server/pkg/mysqldb"
	"GoDistributedAppDevelop/demo06_dstri_tran/inven_server/pkg/redisdb"
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type OrderLgc struct {
}

func (o *OrderLgc) ListenPlaceOrder() {
	// 监听订单消息
	o.consumeMsgBySteamName("order_stream")
}

// ConsumeMsgBySteamName 消费消息
func (o *OrderLgc) consumeMsgBySteamName(_streamName string) (string, error) {
	streams := make([]string, 1)
	streams[0] = _streamName
	// 配置 XReadArgs，用于 Redis 的 XREAD 命令
	args := &redis.XReadArgs{
		Streams: streams,
		Count:   1,               // 每次读取的消息数量
		Block:   5 * time.Second, // 阻塞等待的时间，如果无消息则等待
		ID:      "$",             // $美元符号，代表消息id redis自动生成
	}
	for {
		// 使用 XREAD 命令从指定的流中读取消息
		reciveMsgs, err := redisdb.RedisDb.XRead(context.TODO(), args).Result()
		if err != nil && err != redis.Nil {
			logs.ZapLogger.Error("接收消息错误 error: %v" + err.Error())
			return "", err
		}
		for _, v := range reciveMsgs {
			streamName := v.Stream
			msg := v.Messages[0] //第一条消息
			msgId := msg.ID
			logs.ZapLogger.Debug("接收到消息: steam name: " + streamName + " msgId：" + msgId)
			order := &model.Order{}
			if err := json.Unmarshal([]byte(msg.Values["order"].(string)), order); err != nil {
				logs.ZapLogger.Error("解析消息错误 error: %v" + err.Error())
				return "", err
			}
			//减库存 gorm 操作数据库
			product := &model.Product{}
			product.ID = uint(order.ProductId)
			mysqldb.Mysql.Model(product).Update("inven_num", gorm.Expr("inven_num - ?", order.Num)).Where("id = ?", order.ProductId).Limit(1)
		}
	}
	return "", errors.New("time out")
}
