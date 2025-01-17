package logics

import (
	"GoDistributedAppDevelop/demo06_dstri_tran/tran_server/internal/model"
	"GoDistributedAppDevelop/demo06_dstri_tran/tran_server/pkg/mysqldb"
	"GoDistributedAppDevelop/demo06_dstri_tran/tran_server/pkg/redisdb"
	"GoDistributedAppDevelop/demo06_dstri_tran/tran_server/pkg/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type OrderLgc struct {
}

// 下订单
func (o *OrderLgc) PlaceAnOrder(userId int64, productId int64, num int64) (orderId int64, err error) {
	// 1. 获取商品信息
	product, err := o.GetProductById(productId)
	if err != nil {
		return 0, err
	}
	if product.InvenNum < num {
		//库存不足
		return 0, fmt.Errorf("库存不足")
	}
	//创建订单
	order := &model.Order{
		UserId:    userId,
		ProductId: productId,
		Num:       num,
		Status:    1,
		Amount:    utils.RoundToTwoDecimals(product.Price * float64(num)),
	}
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	//开始事务
	tx := mysqldb.Mysql.Begin()
	if tx.Error != nil {
		//开始事务失败
		return 0, err
	}
	err = tx.Create(order).Error
	if err != nil {
		//创建订单失败
		tx.Rollback()
		return 0, err
	}
	//把order 序列化后，发送到redis stream 消息队列
	orderMarshal, _ := json.Marshal(order)
	_, err = redisdb.RedisDb.XAdd(context.TODO(), &redis.XAddArgs{
		Stream: "order_stream",
		Values: map[string]interface{}{
			"order": string(orderMarshal),
		},
	}).Result()
	if err != nil {
		//发送消息队列失败，事务回滚
		tx.Rollback()
		return 0, err
	} else {
		//发送消息队列成功 事务提交
		tx.Commit()
	}
	return int64(order.ID), nil
}

// 根据商品id获取商品信息
func (o *OrderLgc) GetProductById(productId int64) (*model.Product, error) {
	product := &model.Product{}
	err := mysqldb.Mysql.Model(product).Where("id = ?", productId).Find(&product).Error
	return product, err
}
