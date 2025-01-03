package main

import (
	"GoDistributedAppDevelop/demo10_kafka/config"
	"GoDistributedAppDevelop/demo10_kafka/msg_datas"
	"GoDistributedAppDevelop/demo10_kafka/pkg/kafka"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func main() {
	config.Init()
	kt := kafka.NewKafkaTool()
	for i := 1; i < 11; i++ {
		pmsg := &msg_datas.ProduceMsg{
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
		bmsg := &msg_datas.BaseMsg{
			MsgId:   msg_datas.Msg_hello_id,
			MsgData: pmsgStr,
		}
		jsonBytes, err := json.Marshal(bmsg)
		kt.ProduceMsg(msg_datas.Topic_hello, string(jsonBytes))
		time.Sleep(time.Second * 3)
	}
}
