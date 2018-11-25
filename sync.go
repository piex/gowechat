package gowechat

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/golang/glog"
)

func (w *WeChat) sync() (*Message, error) {
	v := url.Values{}
	v.Add("lang", "zh_CN")
	v.Add("sid", w.loginRes.Wxsid)
	v.Add("skey", w.loginRes.Skey)
	v.Add("pass_ticket", w.loginRes.PassTicket)

	uri := syncURI + "?" + v.Encode()

	params := make(map[string]interface{})
	params["BaseRequest"] = w.baseRequest
	params["SyncKey"] = w.syncKey
	params["rr"] = ^int(time.Now().Unix())

	data, err := w.post(uri, params)
	if err != nil {
		return nil, err
	}

	var msg Message
	if err := json.Unmarshal(data, &msg); err != nil {
		glog.Error("unmarshal fail", err)
		return nil, err
	}

	if msg.BaseResponse.Ret != 0 {
		glog.Error("ret not equal 0", msg.BaseResponse.ErrMsg)
		return nil, fmt.Errorf(msg.BaseResponse.ErrMsg)
	}

	// TODO
	err = w.updateSyncKey(msg)
	if err != nil {
		glog.Error("update synckey fail", err)
		return nil, err
	}

	return &msg, nil
}

func (w *WeChat) updateSyncKey(msg Message) error {
	w.syncKey = msg.SyncKey

	if msg.SKey != "" {
		w.loginRes.Skey = msg.SKey
	}

	if msg.SyncCheckKey.Count != 0 {
		w.syncKey = msg.SyncCheckKey
	}
	return nil
}
