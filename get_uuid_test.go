package gowechat

import "testing"

func TestGetUUID(t *testing.T) {
	w, err := New()
	if err != nil {
		t.Error("[*] New WeChat Error", err)
	}
	err = w.getUUID()
	if err != nil {
		t.Error("[*] Get UUID Error", err)
	}
	t.Log("[*] GET UUID SUCCESS：", w.uuid)
}
