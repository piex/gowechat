package gowechat

import (
	"net/url"
	"regexp"

	"github.com/golang/glog"
)

func (w *WeChat) sendEmotion() error {
	v := url.Values{}
	v.Add("fun", "sys")
	v.Add("pass_ticket", w.loginRes.PassTicket)
	v.Add("lang", "zh_CN")

	uri := sendemoticonURI + "?" + v.Encode()

	msg := make(map[string]interface{})
	msg["Type"] = 47
	msg["FromUserName"] = w.user.UserName
	msg["EmojiFlag"] = 2

	body := make(map[string]interface{})
	body["BaseRequest"] = w.baseRequest
	body["Scene"] = 0

	reg, _ := regexp.Compile("^@")

	for send := range w.api.sendEmotion {
		glog.Info("[#] send new emotion：", send.Message)
		clientMsgID := getClientMsgID()

		msg["LocalID"] = clientMsgID
		msg["ClientMsgId"] = clientMsgID
		msg["ToUserName"] = send.To

		if reg.Match([]byte(send.Message)) {
			msg["MediaId"] = send.Message
		} else {
			msg["EMoticonMd5"] = send.Message
		}

		body["Msg"] = msg

		_, err := w.post(uri, body)
		if err != nil {
			return err
		}
	}

	return nil
}
