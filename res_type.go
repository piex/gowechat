package gowechat

// LoginResult 登录接口返回值
type LoginResult struct {
	Skey       string `xml:"skey"`
	Wxsid      string `xml:"wxsid"`
	Wxuin      string `xml:"wxuin"`
	PassTicket string `xml:"pass_ticket"`
}

// InitResult init 请求返回值
type InitResult struct {
	BaseResponse        BaseResponse     `json:"BaseResponse"`
	Count               int              `json:"Count"`
	ContactList         []Contact        `json:"ContactList"`
	SyncKey             SyncKey          `json:"SyncKey"`
	User                User             `json:"User"`
	ChatSet             string           `json:"ChatSet"`
	SKey                string           `json:"SKey"`
	ClientVersion       int              `json:"ClientVersion"`
	SystemTime          int              `json:"SystemTime"`
	GrayScale           int              `json:"GrayScale"`
	InviteStartCount    int              `json:"InviteStartCount"`
	MPSubscribeMsgCount int              `json:"MPSubscribeMsgCount"`
	MPSubscribeMsgList  []MPSubscribeMsg `json:"MPSubscribeMsgList"`
	ClickReportInterval int              `json:"ClickReportInterval"`
}

// BaseResponse 基础返回参数
type BaseResponse struct {
	Ret    int    `json:"Ret"`
	ErrMsg string `json:"ErrMsg"`
}

// SyncKeyList TODO
type SyncKeyList struct {
	Key int `json:"Key"`
	Val int `json:"Val"`
}

// SyncKey 同步的 key
type SyncKey struct {
	Count int           `json:"Count"`
	List  []SyncKeyList `json:"List"`
}

// Contact 聊天接口返回值
type Contact struct {
	Uin              int      `json:"Uin"`
	UserName         string   `json:"UserName"`
	NickName         string   `json:"NickName"`
	HeadImgURL       string   `json:"HeadImgUrl"`
	ContactFlag      int      `json:"ContactFlag"`
	MemberCount      int      `json:"MemberCount"`
	MemberList       []Member `json:"MemberList"`
	RemarkName       string   `json:"RemarkName"`
	HideInputBarFlag int      `json:"HideInputBarFlag"`
	Sex              int      `json:"Sex"`
	Signature        string   `json:"Signature"`
	VerifyFlag       int      `json:"VerifyFlag"`
	OwnerUin         int      `json:"OwnerUin"`
	PYInitial        string   `json:"PYInitial"`
	PYQuanPin        string   `json:"PYQuanPin"`
	RemarkPYInitial  string   `json:"RemarkPYInitial"`
	RemarkPYQuanPin  string   `json:"RemarkPYQuanPin"`
	StarFriend       int      `json:"StarFriend"`
	AppAccountFlag   int      `json:"AppAccountFlag"`
	Statues          int      `json:"Statues"`
	AttrStatus       int      `json:"AttrStatus"`
	Province         string   `json:"Province"`
	City             string   `json:"City"`
	Alias            string   `json:"Alias"`
	SnsFlag          int      `json:"SnsFlag"`
	UniFriend        int      `json:"UniFriend"`
	DisplayName      string   `json:"DisplayName"`
	ChatRoomID       int      `json:"ChatRoomId"`
	KeyWord          string   `json:"KeyWord"`
	EncryChatRoomID  string   `json:"EncryChatRoomId"`
	IsOwner          int      `json:"IsOwner"`
}

// User 用户相关信息
type User struct {
	Uin               int    `json:"Uin"`
	UserName          string `json:"UserName"`
	NickName          string `json:"NickName"`
	HeadImgURL        string `json:"HeadImgUrl"`
	RemarkName        string `json:"RemarkName"`
	PYInitial         string `json:"PYInitial"`
	PYQuanPin         string `json:"PYQuanPin"`
	RemarkPYInitial   string `json:"RemarkPYInitial"`
	RemarkPYQuanPin   string `json:"RemarkPYQuanPin"`
	HideInputBarFlag  int    `json:"HideInputBarFlag"`
	StarFriend        int    `json:"StarFriend"`
	Sex               int    `json:"Sex"`
	Signature         string `json:"Signature"`
	AppAccountFlag    int    `json:"AppAccountFlag"`
	VerifyFlag        int    `json:"VerifyFlag"`
	ContactFlag       int    `json:"ContactFlag"`
	WebWxPluginSwitch int    `json:"WebWxPluginSwitch"`
	HeadImgFlag       int    `json:"HeadImgFlag"`
	SnsFlag           int    `json:"SnsFlag"`
}

// MPSubscribeMsg TODO
type MPSubscribeMsg struct {
	UserName       string      `json:"UserName"`
	MPArticleCount int         `json:"MPArticleCount"`
	MPArticleList  []MPArticle `json:"MPArticleList"`
	Time           int         `json:"Time"`
	NickName       string      `json:"NickName"`
}

// Member 群组成员对象
type Member struct {
	Uin             int    `json:"Uin"`
	UserName        string `json:"UserName"`
	NickName        string `json:"NickName"`
	AttrStatus      int    `json:"AttrStatus"`
	PYInitial       string `json:"PYInitial"`
	PYQuanPin       string `json:"PYQuanPin"`
	RemarkPYInitial string `json:"RemarkPYInitial"`
	RemarkPYQuanPin string `json:"RemarkPYQuanPin"`
	MemberStatus    int    `json:"MemberStatus"`
	DisplayName     string `json:"DisplayName"`
	KeyWord         string `json:"KeyWord"`
}

// MPArticle TODO
type MPArticle struct {
	Title  string `json:"Title"`
	Digest string `json:"Digest"`
	Cover  string `json:"Cover"`
	URL    string `json:"Url"`
}

// StatusNotifyResult status notify 状态返回值
type StatusNotifyResult struct {
	BaseResponse BaseResponse `json:"BaseResponse"`
	MsgID        string       `json:"MsgID"`
}

// Contacts 多条联系人
type Contacts struct {
	BaseResponse BaseResponse `json:"BaseResponse"`
	MemberCount  int          `json:"MemberCount"`
	MemberList   []Contact    `json:"MemberList"`
	Seq          int          `json:"Seq"`
}

// Message 消息
type Message struct {
	BaseResponse           BaseResponse  `json:"BaseResponse"`
	AddMsgCount            int           `json:"AddMsgCount"`
	AddMsgList             []AddMsg      `json:"AddMsgList"`
	ModContactCount        int           `json:"ModContactCount"`
	ModContactList         []Contact     `json:"ModContactList"`
	DelContactCount        int           `json:"DelContactCount"`
	DelContactList         []interface{} `json:"DelContactList"`
	ModChatRoomMemberCount int           `json:"ModChatRoomMemberCount"`
	ModChatRoomMemberList  []interface{} `json:"ModChatRoomMemberList"`
	Profile                Profile       `json:"Profile"`
	ContinueFlag           int           `json:"ContinueFlag"`
	SyncKey                SyncKey       `json:"SyncKey"`
	SKey                   string        `json:"SKey"`
	SyncCheckKey           SyncKey       `json:"SyncCheckKey"`
}

// AddMsg TODO
type AddMsg struct {
	MsgID                string        `json:"MsgId"`
	FromUserName         string        `json:"FromUserName"`
	ToUserName           string        `json:"ToUserName"`
	MsgType              int           `json:"MsgType"`
	Content              string        `json:"Content"`
	Status               int           `json:"Status"`
	ImgStatus            int           `json:"ImgStatus"`
	CreateTime           int           `json:"CreateTime"`
	VoiceLength          int           `json:"VoiceLength"`
	PlayLength           int           `json:"PlayLength"`
	FileName             string        `json:"FileName"`
	FileSize             string        `json:"FileSize"`
	MediaID              string        `json:"MediaId"`
	URL                  string        `json:"Url"`
	AppMsgType           int           `json:"AppMsgType"`
	StatusNotifyCode     int           `json:"StatusNotifyCode"`
	StatusNotifyUserName string        `json:"StatusNotifyUserName"`
	RecommendInfo        RecommendInfo `json:"RecommendInfo"`
	ForwardFlag          int           `json:"ForwardFlag"`
	AppInfo              AppInfo       `json:"AppInfo"`
	HasProductID         int           `json:"HasProductId"`
	Ticket               string        `json:"Ticket"`
	ImgHeight            int           `json:"ImgHeight"`
	ImgWidth             int           `json:"ImgWidth"`
	SubMsgType           int           `json:"SubMsgType"`
	NewMsgID             int64         `json:"NewMsgId"`
	OriContent           string        `json:"OriContent"`
	EncryFileName        string        `json:"EncryFileName"`
}

// RecommendInfo TODO
type RecommendInfo struct {
	UserName   string `json:"UserName"`
	NickName   string `json:"NickName"`
	QQNum      int    `json:"QQNum"`
	Province   string `json:"Province"`
	City       string `json:"City"`
	Content    string `json:"Content"`
	Signature  string `json:"Signature"`
	Alias      string `json:"Alias"`
	Scene      int    `json:"Scene"`
	VerifyFlag int    `json:"VerifyFlag"`
	AttrStatus int    `json:"AttrStatus"`
	Sex        int    `json:"Sex"`
	Ticket     string `json:"Ticket"`
	OpCode     int    `json:"OpCode"`
}

// AppInfo 应用信息
type AppInfo struct {
	AppID string `json:"AppID"`
	Type  int    `json:"Type"`
}

// Profile TODO
type Profile struct {
	BitFlag           int         `json:"BitFlag"`
	UserName          ProfileBuff `json:"UserName"`
	NickName          ProfileBuff `json:"NickName"`
	BindUin           int         `json:"BindUin"`
	BindEmail         ProfileBuff `json:"BindEmail"`
	BindMobile        ProfileBuff `json:"BindMobile"`
	Status            int         `json:"Status"`
	Sex               int         `json:"Sex"`
	PersonalCard      int         `json:"PersonalCard"`
	Alias             string      `json:"Alias"`
	HeadImgUpdateFlag int         `json:"HeadImgUpdateFlag"`
	HeadImgURL        string      `json:"HeadImgUrl"`
	Signature         string      `json:"Signature"`
}

// ProfileBuff TODO
type ProfileBuff struct {
	Buff string `json:"Buff"`
}
