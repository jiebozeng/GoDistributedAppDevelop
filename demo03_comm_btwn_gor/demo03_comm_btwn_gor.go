package main

import (
	"fmt"
	"github.com/jiebozeng/golangutils/convert"
	"github.com/jiebozeng/golangutils/mathutils"
	"time"
)

// 商品库存基本结构
type ProductSku struct {
	ProductId  int64  //商品id
	ProductSku string //商品sku码
	StockNum   int64  //商品库存
}

func NewProductSku(proId int64, productSku string, stockNum int64) *ProductSku {
	ret := &ProductSku{
		ProductId:  proId,
		ProductSku: productSku,
		StockNum:   stockNum,
	}
	return ret
}
func (p *ProductSku) Print() {
	fmt.Println("商品id:", p.ProductId, "商品Sku码:", p.ProductSku, "商品库存:", p.StockNum)
}

// 下订单，扣减库存 通信消息的结构体
type NotifySubInventory struct {
	ProductId  int64  //商品id
	ProductSku string //商品的sku码 例如：T00101
	SubNum     int64  //要扣除的库存数量
}

func NewNotifySubInventory(proId int64, productSku string, subNum int64) *NotifySubInventory {
	ret := &NotifySubInventory{
		ProductId:  proId,
		ProductSku: productSku,
		SubNum:     subNum,
	}
	return ret
}

// 下订单
func CreateOrder(ch chan NotifySubInventory) {
	//每隔5秒随机商品下单，总下单50单停止
	count := 0
	for {
		//创建订单，下单逻辑
		//......
		if count > 50 {
			fmt.Println("模拟下单结束！")
			break
		}
		//通知减库存
		i := mathutils.RandInterval(0, 9)
		proId := int64(1000 + i)
		proSku := "T0010" + convert.ToString(i)
		notifyMsg := *NewNotifySubInventory(proId, proSku, 1)
		ch <- notifyMsg
		count++
		time.Sleep(time.Second * 2)
	}
}

// 减少库存
func SubInventory(ch chan NotifySubInventory, proSkuAr []*ProductSku) {
	for {
		select {
		case msg := <-ch:
			fmt.Println("收到减库存的消息")
			for _, proSku := range proSkuAr {
				if msg.SubNum > 0 && proSku.ProductId == msg.ProductId && proSku.ProductSku == msg.ProductSku {
					proSku.StockNum -= msg.SubNum
					proSku.Print()
				}
			}
		}
	}
}

func main() {
	//商品库存信息
	fmt.Println("初始化商品库存的数据：")
	productSkuAr := make([]*ProductSku, 10)
	for i := int64(0); i < 10; i++ {
		proSku := NewProductSku(1000+i, "T0010"+fmt.Sprintf("%d", i), 100)
		productSkuAr[i] = proSku
	}
	//输出打印
	for _, r := range productSkuAr {
		r.Print()
	}
	//初始化消息通知通道chan
	ch := make(chan NotifySubInventory)
	//开启一个协程去模拟下单
	go CreateOrder(ch)
	//开启一个协程去减库存
	go SubInventory(ch, productSkuAr)

	//等待终端输入，避免协程没执行，就退出
	fmt.Scanln()
}
