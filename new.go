package gowechat

import (
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"time"

	"github.com/golang/glog"
)

var wechatAPI = &API{
	MsgChan:     make(chan Message),
	sendMsg:     make(chan SendMsg),
	sendEmotion: make(chan SendMsg),
}

// New 初始化 WeChat
func New() (*WeChat, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		glog.Error("get cookiejar fail", err)
		return nil, err
	}

	client := &http.Client{
		CheckRedirect: nil,
		Jar:           jar,
		Timeout:       60 * time.Second,
	}

	rand.Seed(time.Now().Unix())
	randID := strconv.Itoa(rand.Int())

	return &WeChat{
		client:      client,
		deviceID:    "e" + randID[2:17],
		baseRequest: make(map[string]interface{}),
		contacts:    make(map[string]Contact),
		api:         wechatAPI,
	}, nil
}
