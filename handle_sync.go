package gowechat

import (
	"strings"

	"github.com/golang/glog"
)

func (w *WeChat) handleSync(msg *Message) error {

	w.updateContacts(msg.ModContactList)

	for _, m := range msg.AddMsgList {

		m.Content = strings.Replace(m.Content, "&lt;", "<", -1)
		m.Content = strings.Replace(m.Content, "&gt;", ">", -1)

		switch m.MsgType {
		case 1:
			if m.FromUserName[:2] == "@@" { // 群消息
				glog.Infof("[群消息] %s: %s", w.getNickName(m.FromUserName), m.Content)
			} else {
				// 普通消息
				glog.Infof("[好友消息] %s: %s", w.getNickName(m.FromUserName), m.Content)
			}
		case 51:
			glog.Info("sync ok.")
		}
	}

	w.api.MsgChan <- *msg

	return nil
}
