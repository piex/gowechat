package gowechat

import (
	"net/url"

	"github.com/golang/glog"
)

// Logout 登出微信
func (w *WeChat) Logout() {
	v := url.Values{}
	v.Add("redirect", "1")
	v.Add("type", "0")
	v.Add("lang", "zh_CN")
	v.Add("skey", w.loginRes.Skey)

	uri := logoutURI + "?" + v.Encode()

	_, err := w.post(uri, nil)
	if err != nil {
		glog.Warning("logout maybe.")
	}

	glog.Info("logout success.")
}
