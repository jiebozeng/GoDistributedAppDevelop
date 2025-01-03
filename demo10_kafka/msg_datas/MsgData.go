package msg_datas

const Msg_hello_id int64 = 1001        // 消息id
const Topic_hello string = "hello_msg" // 消息主题

// 消息结构
type BaseMsg struct {
	MsgId   int64  `json:"msg_id" bson:"msg_id"`
	MsgData string `json:"msg_data" bson:"msg_data"` // 消息具体内容
}

// 要生产的消息结构体
type ProduceMsg struct {
	ToUid   string `json:"to_uid" bson:"to_uid"`     // 接收者id
	SendUid string `json:"send_uid" bson:"send_uid"` // 发送者id
	MsgType int    `json:"msg_type" bson:"msg_type"` // 消息类型
	Content string `json:"content" bson:"content"`   // 消息内容
}
