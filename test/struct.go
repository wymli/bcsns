package main

type SendMsgRequest struct {
	ActionType  string `json:"action_type,omitempty,options=user_chat|room_chat|moments|drift"`  // 消息类型
	ContentType string `json:"content_type,omitempty,options=user_chat|room_chat|moments|drift"` // 内容类型
	SendMsgId   string `json:"send_msg_id,omitempty"`  // 用于服务端去重,客户端去重
	SendTime    int64  `json:"send_time,omitempty"`    // 用于接收端排序,但客户端时间不可信,除非先时间同步
	Token       string `json:"token,omitempty"`        // 用于认证鉴权
	SendUserId  uint64 `json:"send_user_id,omitempty"` // user id
	RoomId      uint64 `json:"room_id,omitempty"`      // room id
	RecvUserId  uint64 `json:"recv_user_id,omitempty"` // user id
	Data        []byte `json:"data,omitempty"`         // payload body
	IsTest      *bool  `json:"is_test,omitempty"`      // 是否测试
}
