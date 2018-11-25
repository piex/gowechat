package gowechat

// getNickName 获取昵称
func (w *WeChat) getNickName(userName string) string {
	if v, ok := w.contacts[userName]; ok {
		return v.NickName
	}

	return userName
}
