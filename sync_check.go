package gowechat

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

// strSyncKey 拼接 syncKey
func (w *WeChat) strSyncKey() string {
	kvs := []string{}
	for _, list := range w.syncKey.List {
		kvs = append(kvs, strconv.Itoa(list.Key)+"_"+strconv.Itoa(list.Val))
	}

	return strings.Join(kvs, "|")
}

// syncCheck 定时检查是否有新信息
func (w *WeChat) syncCheck() (int, error) {

	v := url.Values{}
	v.Add("r", timestamp())
	v.Add("sid", w.loginRes.Wxsid)
	v.Add("uin", strconv.FormatInt(w.baseRequest["Uin"].(int64), 10))
	v.Add("skey", w.baseRequest["Skey"].(string))
	v.Add("deviceid", w.deviceID)
	v.Add("synckey", w.strSyncKey())
	v.Add("_", timestamp())

	uri := synccheckURI + "?" + v.Encode()

	data, err := w.get(uri)
	if err != nil {
		glog.Error("sync check fail", err)
		return 0, err
	}

	reg := regexp.MustCompile(`window.synccheck={retcode:"(\d+)",selector:"(\d+)"}`)
	codes := reg.FindStringSubmatch(string(data))
	if len(codes) > 2 {
		retcode, _ := strconv.Atoi(codes[1])
		selector, _ := strconv.Atoi(codes[2])

		switch retcode {

		case 1100: // 手机上退出网页版微信
			errMsg := "[*] logout with phone, bye."
			glog.Error(errMsg)
			return 0, fmt.Errorf(errMsg)

		case 1101: // 在其他地方登录网页版微信
			errMsg := "[*] login wechat at other palce, bye"
			glog.Error(errMsg)
			return 0, fmt.Errorf(errMsg)

		case 0: // 同步成功
			return selector, nil
		}

	}
	return 0, fmt.Errorf("sync check fail")
}
