package gowechat

import (
	"fmt"
	"strings"

	"github.com/golang/glog"
)

// 获取 uuid
func (w *WeChat) getUUID() error {
	if w.uuid != "" {
		glog.Info("[*] already get uuid")
		return nil
	}

	uri := jsloginURI + timestamp()

	data, err := w.get(uri)
	if err != nil {
		glog.Error(err)
		return nil
	}

	res := make(map[string]string)

	datas := strings.Split(string(data), ";")

	for _, d := range datas {
		kvs := strings.Split(d, " = ")
		if len(kvs) == 2 {
			res[strings.Trim(kvs[0], " ")] = strings.Trim(strings.Trim(kvs[1], " "), "\"")
		}
	}

	if res["window.QRLogin.code"] == "200" {
		if uuid, ok := res["window.QRLogin.uuid"]; ok {
			w.uuid = uuid
			glog.Info("[*] get uuid success", uuid)
			return nil
		}
	}

	return fmt.Errorf(string(data))
}
