package gowechat

import (
	"github.com/golang/glog"
)

// start 启动
func start() (*WeChat, error) {
	glog.Info("[*] wechat start...")
	wx, err := New()
	if err != nil {
		return nil, err
	}

	run("[*] get uuid ...", wx.getUUID)
	run("[*] show qrcode ...", wx.showQrcode)
	run("[*] check login ...", wx.checkLogin)
	run("[*] get conf ... ", wx.getConf)
	run("[*] login ...", wx.login)
	run("[*] init wechat ...", wx.wxinit)
	run("[*] open status notify ...", wx.notify)
	state = StateLogin
	// TODO updateContacts
	run("[*] get contact ...", wx.getContact)
	// run("[*] send text ... ", wx.sendText)
	go wx.sendText()
	run("[*] sync polling ... ", wx.syncPolling)
	// TODO checkPolling

	return wx, nil
}

// Start 启动
func Start() *API {
	go start()
	return wechatAPI
}
