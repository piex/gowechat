package gowechat

import (
	// "golang.org/x/tools/go/packages"
	"flag"
	"net/http"

	"github.com/golang/glog"
)

func init() {
	flag.Parse()

	glog.Info("[*] wechat init.")
}

// WeChat WeChat struct
type WeChat struct {
	uuid        string                 // uuid
	deviceID    string                 // device id
	client      *http.Client           // weixin server http client
	redirectURI string                 // redirect uri
	baseURI     string                 // base URI
	baseRequest map[string]interface{} // base request
	loginRes    LoginResult            // login response
	user        User                   // user info
	syncKey     SyncKey                // 同步数据的 key
	contacts    map[string]Contact     // 联系人
	api         *API                   // 暴露出来的 api
}
