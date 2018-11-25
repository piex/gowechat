package gowechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"
)

// get 发送 GET 请求
func (w *WeChat) get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	req.Header.Add("Referer", origin)
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Connection", "close")

	res, err := w.client.Do(req)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		return ioutil.ReadAll(res.Body)
	}

	return nil, fmt.Errorf(fmt.Sprintf("%v", res.StatusCode))
}

// 发送 POST 请求
func (w *WeChat) post(url string, params map[string]interface{}) ([]byte, error) {

	data, err := json.Marshal(params)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	body := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Referer", origin)
	req.Header.Add("User-Agent", userAgent)

	resp, err := w.client.Do(req)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
