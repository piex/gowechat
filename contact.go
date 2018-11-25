package gowechat

import (
	"encoding/json"
	"net/url"

	"github.com/golang/glog"
)

// getContact 通讯录中的联系人（包括已保存的群聊）
func (w *WeChat) getContact() error {

	v := url.Values{}
	v.Add("lang", "zh_CN")
	v.Add("seq", "0")
	v.Add("skey", w.loginRes.Skey)
	v.Add("pass_ticket", w.loginRes.PassTicket)
	v.Add("r", timestamp())

	uri := getcontactURI + "?" + v.Encode()

	data, err := w.get(uri)
	if err != nil {
		glog.Error("get contact fail", err)
		return err
	}

	var contacts Contacts
	if err := json.Unmarshal(data, &contacts); err != nil {
		glog.Error("unmarshal fail", err)
		return err
	}

	w.updateContacts(contacts.MemberList)

	return nil
}

// 临时的群聊会话在初始化的接口中可以获取，因此这里也需要更新一遍 contacts，否则后面可能会拿不到某个临时群聊的信息
func (w *WeChat) updateContacts(contacts []Contact) error {

	for _, contact := range contacts {
		if contact.NickName == "" {
			contact.NickName = contact.UserName
		}
		w.contacts[contact.UserName] = contact
	}

	return nil
}
