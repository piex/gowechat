package gowechat

import (
	"fmt"

	"github.com/golang/glog"
)

// 报告状态
func (w *WeChat) startReport(arg interface{}) error {

	var text string

	if arg != nil {
		text = arg.(string)
	} else {
		text = fmt.Sprintf("{'type': '[action-record]', 'data': { 'actions': [{ 'type': 'click', 'action': '发送框', 'time': %v }] } }", timestamp())
	}

	report := Report{
		Text: text,
		Type: 1,
	}

	reports := []Report{report}

	uri := fmt.Sprintf("%s??fun=new&lang=zh_CN&pass_ticket=%s", statreportURI, w.loginRes.PassTicket)
	params := make(map[string]interface{})
	params["BaseRequest"] = w.baseRequest
	params["Count"] = len(reports)
	params["List"] = reports

	_, err := w.post(uri, params)

	if err != nil {
		glog.Error("start report fail", err)
		return err
	}
	return nil
}
