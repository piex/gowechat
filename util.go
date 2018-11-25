package gowechat

import (
	"math/rand"
	"strconv"
	"time"
)

func timestamp() string {
	return strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
}

// 生成客户端 id
func getClientMsgID() string {
	return timestamp() + "0" + strconv.Itoa(rand.Int())[3:6]
}
