package gowechat

import (
	"fmt"
	"time"

	"github.com/golang/glog"
)

// State 状态
type State int

const (
	// StateInit init 状态
	StateInit State = iota
	// StateUUID uuid 状态
	StateUUID
	// StateLogin login 状态
	StateLogin
	// StateLogout logout 状态
	StateLogout
)

var syncPollingID = 0

var state State

// 同步失败次数
var syncErrorCount = 0

// 最后一次同步时间
var lastSyncTime = time.Now()

// 同步
func (w *WeChat) syncPolling() error {
	if state != StateLogin {
		return fmt.Errorf("state not login")
	}

	selector, err := w.syncCheck()
	if err != nil { // 同步失败
		if state != StateLogin {
			return err
		}

		syncErrorCount++
		if syncErrorCount > 2 {
			// 多次失败后重启
			glog.Errorf("连续%v次同步失败，5s后尝试重启", syncErrorCount)
			time.Sleep(5 * time.Second)
			// TODO w.reStart()
			Start()
		} else {
			// 两秒后重新尝试心跳连接
			time.Sleep(time.Duration(syncErrorCount) * 2 * time.Second)
			w.syncPolling()
		}

	}

	glog.Info("[*] sync check success, selector：", selector)

	// selector:
	//       a. 0 正常
	//       b. 2 新的消息
	//       c. 7 手机操作了微信
	if selector != 0 {
		msg, err := w.sync()
		if err != nil {
			glog.Error("sync fail", err)
		}
		syncErrorCount = 0
		w.handleSync(msg)
	} else {
		lastSyncTime = time.Now()
	}

	return w.syncPolling()
}

// 心跳轮训检测
func (w *WeChat) checkPolling() {

}
