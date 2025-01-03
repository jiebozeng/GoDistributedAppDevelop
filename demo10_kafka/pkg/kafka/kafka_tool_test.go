package kafka

import (
	"GoDistributedAppDevelop/demo10_kafka/config"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"strconv"
	"sync"
	"testing"
	"time"
)

const Ka_produce_msg_id int64 = 1001

type BaseMsg struct {
	MsgId   int64  `json:"msg_id" bson:"msg_id"`
	MsgData string `json:"msg_data" bson:"msg_data"`
}

type ProduceMsg struct {
	ToUid   string `json:"to_uid" bson:"to_uid"`
	SendUid string `json:"send_uid" bson:"send_uid"`
	MsgType int    `json:"msg_type" bson:"msg_type"`
	Content string `json:"content" bson:"content"`
}

// 注意 用test测试的时候  config.Init路径要修改
func TestKafKaInit(t *testing.T) {
	config.Init()
	fmt.Println(config.CONFIG.Kafka.Host, config.CONFIG.Kafka.Port, 999)
	kt := NewKafkaTool()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 1; i < 11; i++ {
			time.Sleep(time.Second * 3)

			pmsg := &ProduceMsg{
				ToUid:   "1",
				MsgType: 1,
				Content: "hello world " + strconv.Itoa(i),
			}

			pmsgBytes, err := json.Marshal(pmsg)
			if err != nil {
				fmt.Println(err)
				continue
			}
			pmsgStr := base64.StdEncoding.EncodeToString(pmsgBytes)
			bmsg := &BaseMsg{
				MsgId:   Ka_produce_msg_id,
				MsgData: pmsgStr,
			}
			jsonBytes, err := json.Marshal(bmsg)
			kt.ProduceMsg("he_msg", string(jsonBytes))
		}
		wg.Done()
	}()
	//kt.AddConsumerHander(handler.GroupId, handler.Topic, handler)
	wg.Wait()
}
