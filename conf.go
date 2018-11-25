package gowechat

import (
	"fmt"
	"regexp"

	"github.com/golang/glog"
)

var (
	userAgent          = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36"
	host               = "wx.qq.com"                                                                       // 微信主域名
	origin             = "https://" + host                                                                 // 完整微信主域名
	loginDomain        = "login.weixin.qq.com"                                                             // 登陆域名
	fileDomain         = "file.wx.qq.com"                                                                  // 文件主域名
	pushDomain         = "webpush.weixin.qq.com"                                                           // 推送域名
	qrcodeURI          = "https://" + loginDomain + "/qrcode"                                              // 获取二维码
	baseURI            = "https://" + loginDomain + "/cgi-bin/mmwebwx-bin"                                 // 微信 base uri
	jsloginURI         = "https://" + loginDomain + "/jslogin?appid=wx782c26e4c19acffb&fun=new&lang=zh-CN" // 获取 uuid
	loginURI           = baseURI + "/login"                                                                // 确认二维码扫面状态
	initURI            = baseURI + "/webwxinit"                                                            // weixin init
	statusnotifyURI    = baseURI + "/webwxstatusnotify"                                                    // 手机状态同志
	getcontactURI      = baseURI + "/webwxgetcontact"                                                      // 获取所有联系人
	batchgetcontactAPI = baseURI + "/webwxbatchgetcontact"                                                 // 批量获取联系人
	statreportURI      = baseURI + "/webwxstatreport"                                                      // 报告状态
	synccheckURI       = baseURI + "/synccheck"                                                            // 检查新信息
	syncURI            = baseURI + "/webwxsync"                                                            // 同步新信息
	logoutURI          = baseURI + "/webwxlogout"                                                          // 登出
	sendmsgURI         = baseURI + "/webwxsendmsg"                                                         // 发送文本
	sendimgURI         = baseURI + "/webwxsendmsgimg"                                                      // 发送图片
	sendvideoURI       = baseURI + "/webwxsendvideomsg"                                                    // 发送视频
	sendappmsgURI      = baseURI + "/webwxsendappmsg"                                                      // 发送收藏表情
	sendemoticonURI    = baseURI + "/webwxsendemoticon"                                                    // 发送表情
)

func (w *WeChat) getConf() error {
	reg, _ := regexp.Compile(`(wx|web)(\d?).(qq|wechat).com`)

	host = reg.FindString(string(w.redirectURI)) // 微信主域名
	glog.Info("[*] Host is：", host)

	if host == "" {
		return fmt.Errorf("get conf fail %s", host)
	}

	origin = "https://" + host                                                                     // 完整微信主域名
	loginDomain = "login." + host                                                                  // 登陆域名
	fileDomain = "file." + host                                                                    // 文件主域名
	pushDomain = "webpush." + host                                                                 // 推送域名
	qrcodeURI = "https://" + loginDomain + "/qrcode"                                               // 获取二维码
	baseURI = "https://" + loginDomain + "/cgi-bin/mmwebwx-bin"                                    // 微信 base uri
	jsloginURI = "https://" + loginDomain + "/jslogin?appid=wx782c26e4c19acffb&fun=new&lang=zh-CN" // 获取 uuid
	loginURI = baseURI + "/login"                                                                  // 确认二维码扫面状态
	initURI = baseURI + "/webwxinit"                                                               // weixin init
	statusnotifyURI = baseURI + "/webwxstatusnotify"                                               // 手机状态同志
	getcontactURI = baseURI + "/webwxgetcontact"                                                   // 获取所有联系人
	batchgetcontactAPI = baseURI + "/webwxbatchgetcontact"                                         // 批量获取联系人
	statreportURI = baseURI + "/webwxstatreport"                                                   // 报告状态
	synccheckURI = baseURI + "/synccheck"                                                          // 检查新信息
	syncURI = baseURI + "/webwxsync"                                                               // 同步新信息
	logoutURI = baseURI + "/webwxlogout"                                                           // 登出
	sendmsgURI = baseURI + "/webwxsendmsg"                                                         // 发送文本
	sendimgURI = baseURI + "/webwxsendmsgimg"                                                      // 发送图片
	sendvideoURI = baseURI + "/webwxsendvideomsg"                                                  // 发送视频
	sendappmsgURI = baseURI + "/webwxsendappmsg"                                                   // 发送收藏表情
	sendemoticonURI = baseURI + "/webwxsendemoticon"                                               // 发送表情

	return nil
}
