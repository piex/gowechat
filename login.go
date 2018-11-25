package gowechat

import (
	"encoding/xml"
	"strconv"

	"github.com/golang/glog"
)

func (w *WeChat) login() error {
	res, err := w.get(w.redirectURI + "&fun=new")
	if err != nil {
		glog.Error("login fail", err)
	}

	var lr LoginResult
	if err = xml.Unmarshal(res, &lr); err != nil {
		glog.Error("unmarshal fail", err)
		return err
	}

	w.loginRes = lr

	Uin, err := strconv.ParseInt(lr.Wxuin, 10, 64)
	if err != nil {
		glog.Error("redirct fail", err)
		return err
	}

	w.baseRequest["Uin"] = Uin
	w.baseRequest["Sid"] = lr.Wxsid
	w.baseRequest["Skey"] = lr.Skey
	w.baseRequest["DeviceID"] = w.deviceID
	return nil
}
