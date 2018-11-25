package gowechat

import (
	"net/url"

	"github.com/golang/glog"
)

func (w *WeChat) sendText() error {
	v := url.Values{}
	v.Add("pass_ticket", w.loginRes.PassTicket)
	v.Add("lang", "zh_CN")

	uri := sendmsgURI + "?" + v.Encode()

	msg := make(map[string]interface{})
	msg["Type"] = 1
	msg["FromUserName"] = w.user.UserName

	body := make(map[string]interface{})
	body["BaseRequest"] = w.baseRequest
	body["Scene"] = 0

	for send := range w.api.sendMsg {

		glog.Info("[#] send new text：", send.Message)
		clientMsgID := getClientMsgID()

		msg["Content"] = send.Message
		msg["LocalID"] = clientMsgID
		msg["ClientMsgId"] = clientMsgID
		msg["ToUserName"] = send.To

		body["Msg"] = msg

		_, err := w.post(uri, body)
		if err != nil {
			return err
		}

	}

	return nil
}
