package gowechat

import (
	"net/url"
	"regexp"

	"github.com/golang/glog"
)

// 检查二维码扫描和登陆状态
func (w *WeChat) checkLogin() error {

	v := url.Values{}
	v.Add("tip", "0")
	v.Add("uuid", w.uuid)
	v.Add("loginicon", "true")
	v.Add("_", timestamp())

	uri := loginURI + "?" + v.Encode()

	data, err := w.get(uri)
	if err != nil {
		glog.Error(err)
		return err
	}

	reg := regexp.MustCompile(`window.code=(\d+);`)
	codes := reg.FindStringSubmatch(string(data))

	if len(codes) > 1 {
		code := codes[1]
		switch code {
		case "201": // 扫码成功
			glog.Info("[*] scan code success.")
			w.checkLogin()
			break
		case "200": // 登陆成功
			glog.Info("[*] login success, wait to redirect.")
			reg := regexp.MustCompile(`window.redirect_uri="(\S+?)";`)
			redirctURIs := reg.FindStringSubmatch(string(data))

			if len(redirctURIs) > 1 {
				redirctURI := redirctURIs[1] + "&fun=new&version=v2"
				w.redirectURI = redirctURI
				reg = regexp.MustCompile(`/`)
				baseURIs := reg.FindAllStringIndex(redirctURI, -1)
				w.baseURI = redirctURI[:baseURIs[len(baseURIs)-1][0]]
			}
			break
		case "408": // 登陆超时
			w.checkLogin()
		}
	}

	return nil
}
