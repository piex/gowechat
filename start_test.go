package gowechat

import (
	"fmt"
	"strings"
	"testing"

	"github.com/golang/glog"
)

func TestStart(t *testing.T) {
	wx := Start()

	for msg := range wx.MsgChan {
		for _, m := range msg.AddMsgList {

			m.Content = strings.Replace(m.Content, "&lt;", "<", -1)
			m.Content = strings.Replace(m.Content, "&gt;", ">", -1)

			switch m.MsgType {
			case 1:
				if m.FromUserName[:2] == "@@" { // 群消息
					fmt.Printf("[群消息] %s", m.Content)
					wx.SendMessage(m.Content, m.FromUserName)
				} else {
					// 普通消息
					fmt.Printf("[好友消息] %s", m.Content)
					wx.SendMessage(m.Content, m.FromUserName)
				}
			}
		}

		glog.Flush()
		t.Log("start success")
	}
}
