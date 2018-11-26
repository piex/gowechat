package gowechat

// SendMsg 发送信息
type SendMsg struct {
	To      string // 目标用户
	Message string // 信息
}

// API 对外暴露出去的 api
type API struct {
	User        User         // user info
	MsgChan     chan Message // 新信息信道
	sendMsg     chan SendMsg
	sendEmotion chan SendMsg
}

// SendMessage 发送消息
func (a *API) SendMessage(message, to string) {
	s := SendMsg{
		To:      to,
		Message: message,
	}

	a.sendMsg <- s
}
