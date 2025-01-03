package main

import (
	"GoDistributedAppDevelop/demo10_kafka/config"
	"GoDistributedAppDevelop/demo10_kafka/msg_datas"
	"GoDistributedAppDevelop/demo10_kafka/pkg/kafka"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
)

// HelloConsumerGroupHandler 实现了 sarama.ConsumerGroupHandler 接口
type HelloConsumerGroupHandler struct {
}

// Setup 在消费者组创建或重新平衡时调用
func (h *HelloConsumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	fmt.Println("Consumer group session started")
	return nil
}

// Cleanup 在消费者组关闭或重新平衡结束时调用
func (h *HelloConsumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	fmt.Println("Consumer group session ended")
	return nil
}

// ConsumeClaim 处理每个分区的消息
func (h *HelloConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Printf("Message claimed: topic = %s,timestamp = %v\n", message.Topic, message.Timestamp)
		//解包
		baseMsg := &msg_datas.BaseMsg{}
		err := json.Unmarshal(message.Value, baseMsg)
		if err != nil {
			fmt.Println(err)
		} else {
			switch baseMsg.MsgId {
			case msg_datas.Msg_hello_id: //对相应消息id做出相应的处理
				helloMsg := &msg_datas.ProduceMsg{}
				msgDataBytes, _ := base64.StdEncoding.DecodeString(baseMsg.MsgData)
				err = json.Unmarshal(msgDataBytes, helloMsg)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("helloMsg：", helloMsg.ToUid, helloMsg.SendUid, helloMsg.MsgType, helloMsg.Content)
				}
			}
		}
		// 标记消息已处理
		session.MarkMessage(message, "")
	}
	return nil
}

func main() {
	config.Init()
	kt := kafka.NewKafkaTool()

	// 定义消费者组 ID 和要消费的主题
	groupId := "my-consumer-group"
	topics := []string{msg_datas.Topic_hello}

	// 创建消费者组处理器实例
	handler := &HelloConsumerGroupHandler{}

	// 添加消费者处理器
	kt.AddConsumerHander(groupId, topics, handler)
}
