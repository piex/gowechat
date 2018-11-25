package gowechat

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang/glog"
)

// statusNotify 报告自己的状态
func (w *WeChat) notify() error {
	uri := fmt.Sprintf("%s?lang=zh_CN&pass_ticket=%s", statusnotifyURI, w.loginRes.PassTicket)
	params := make(map[string]interface{})
	params["BaseRequest"] = w.baseRequest
	params["Code"] = 3
	params["FromUserName"] = w.user.UserName
	params["ToUserName"] = w.user.UserName
	params["ClientMsgId"] = int(time.Now().Unix())

	data, err := w.post(uri, params)
	if err != nil {
		glog.Error(err)
		return err
	}

	var res StatusNotifyResult

	if err := json.Unmarshal(data, &res); err != nil {
		glog.Error("unmarshal fail", err)
		return err
	}

	return nil
}
