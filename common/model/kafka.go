package model

const (
	SendP2PMsgKey   = "send_p2p_msg"
	SendGroupMsgKey = "send_group_msg"
)

type SendP2PMsgKafka struct {
	IncrementID      int64  `json:"incrementID"`
	SourceUID        int64  `json:"sourceUID"`
	TargetUID        int64  `json:"targetUID"`
	MsgID            string `json:"msgID"`
	MsgType          string `json:"msgType"`
	Msg              string `json:"msg"`
	AccessServerAddr string `json:"accessServerAddr"`
	Online           bool   `json:"online"`
}

type SendGroupMsgKafka struct {
	SourceUID int64
	GroupID   int64
	MsgID     string
	Msg       string
}
