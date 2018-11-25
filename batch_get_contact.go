package gowechat

import (
	"encoding/json"
	"fmt"

	"github.com/golang/glog"
)

// 批量查找联系人
func (w *WeChat) batchGetContact(contacts []Contact) error {
	uri := fmt.Sprintf("%s?lang=zh_CN&type=ex&pass_ticket=%s&r=%v", batchgetcontactAPI, w.loginRes.PassTicket, timestamp())

	params := make(map[string]interface{})
	params["BaseRequest"] = w.baseRequest
	params["Count"] = len(contacts)
	params["List"] = contacts

	data, err := w.post(uri, params)
	if err != nil {
		glog.Error("get contact fail", err)
		return err
	}

	var batchContacts BatchContacts
	if err := json.Unmarshal(data, &batchContacts); err != nil {
		glog.Error("unmarshal fail", err)
		return err
	}

	w.updateContacts(batchContacts.ContactList)

	return nil
}
