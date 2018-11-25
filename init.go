package gowechat

import (
	"encoding/json"
	"fmt"

	"github.com/golang/glog"
)

// web wx init
func (w *WeChat) wxinit() error {
	uri := fmt.Sprintf("%s?pass_ticket=%s&skey=%s&r=%s&lang=zh_CN", initURI, w.loginRes.PassTicket, w.loginRes.Skey, timestamp())

	params := make(map[string]interface{})
	params["BaseRequest"] = w.baseRequest

	data, err := w.post(uri, params)
	if err != nil {
		glog.Error("weixin init fail", err)
		return err
	}

	var res InitResult

	if err := json.Unmarshal(data, &res); err != nil {
		glog.Error("unmarshal fail", err)
		return err
	}

	w.user = res.User
	w.api.User = res.User
	w.syncKey = res.SyncKey

	w.updateContacts(res.ContactList)

	return nil
}
