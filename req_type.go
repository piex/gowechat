package gowechat

// BatchContacts 批量查询联系人
type BatchContacts struct {
	BaseResponse BaseResponse `json:"BaseResponse"`
	Count        int          `json:"Count"`
	ContactList  []Contact    `json:"ContactList"`
}

// Report 报告状态
type Report struct {
	Text string `json:"Text"`
	Type int    `json:"Type"`
}

// TODO
// ReportText 报告状态内容
type ReportText struct {
}
