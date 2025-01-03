package kafka

import (
	config2 "GoDistributedAppDevelop/demo10_kafka/config"
	"context"
	"fmt"
	"github.com/IBM/sarama"
)

// KafkaProducer 封装了 sarama.SyncProducer，用于发送消息到 Kafka
type KafkaProducer struct {
	Producer sarama.SyncProducer // sarama.SyncProducer 实例，用于同步发送消息
}

// KafkaConsumer 封装了 sarama.Consumer 和 sarama.ConsumerGroup，用于从 Kafka 消费消息
type KafkaConsumer struct {
	Consumer      sarama.Consumer      // sarama.Consumer 实例，用于消费消息
	ConsumerGroup sarama.ConsumerGroup // sarama.ConsumerGroup 实例，用于消费者组消费模式
}

// 我们封装下，定义结构 KafkaTool，包含了 Kafka 集群的 broker 地址、配置以及生产者和消费者实例
type KafkaTool struct {
	// kafka集群的broker地址列表
	Brokers  []string
	Config   *sarama.Config
	Producer *KafkaProducer
	Consumer *KafkaConsumer
}

func NewKafkaTool() *KafkaTool {
	// 从配置文件中读取 Kafka broker 地址和端口，并拼接成地址字符串
	bros := config2.CONFIG.Kafka.Host + ":" + config2.CONFIG.Kafka.Port
	// 初始化 KafkaTool 实例
	kt := &KafkaTool{
		Brokers:  []string{bros},
		Producer: &KafkaProducer{},
		Consumer: &KafkaConsumer{},
	}
	//配置对象
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 设置生产者等待所有副本确认
	config.Producer.Compression = sarama.CompressionSnappy    // 设置消息压缩方式为 Snappy
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 设置分区器为随机分区器
	config.Producer.Return.Successes = true                   // 设置返回成功发送的消息
	config.Producer.Return.Errors = true                      // 设置返回发送失败的消息
	kt.Config = config                                        // 将配置对象赋值给 KafkaTool 实例

	// 创建并初始化 sarama.SyncProducer 实例
	_producer, err := sarama.NewSyncProducer(kt.Brokers, config)
	if err != nil {
		panic(err)
	}
	kt.Producer.Producer = _producer

	// 创建并初始化 sarama.Consumer 实例
	consumer, err := sarama.NewConsumer(kt.Brokers, config)
	if err != nil {
		panic(err)
	}
	kt.Consumer.Consumer = consumer
	return kt
}

// AddConsumerHander 添加一个消费者处理器到 KafkaTool 实例
// groupId: 消费者组 ID
// topics: 需要消费的主题列表
// handler: 处理消息的回调函数
func (kt *KafkaTool) AddConsumerHander(groupId string, topics []string, handler sarama.ConsumerGroupHandler) {
	consumerGroup, err := sarama.NewConsumerGroup(kt.Brokers, groupId, kt.Config)
	if err != nil {
		panic(err)
	}
	kt.Consumer.ConsumerGroup = consumerGroup
	for {
		// 使用消费者组消费指定主题的消息，并调用 handler 处理消息
		err = consumerGroup.Consume(context.Background(), topics, handler)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// ProduceMsg 使用 KafkaTool 实例的生产者发送一条消息到指定的主题
// topic: 目标主题
// message: 要发送的消息内容
func (kt *KafkaTool) ProduceMsg(topic string, message string) {
	// 创建一个 sarama.ProducerMessage 实例，包含主题和消息内容
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message), // 将消息内容编码为字符串
	}
	// 使用生产者发送消息，并获取消息的分区和偏移量
	partition, offset, err := kt.Producer.Producer.SendMessage(msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("生产消息:", msg)
	// 输出发送成功的分区和偏移量信息
	fmt.Printf("Partition: %d, Offset: %d\n", partition, offset)
}
