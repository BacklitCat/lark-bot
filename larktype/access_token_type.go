package larktype

type AppAccessTokenReq struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type AppAccessTokenResp struct {
	Code           int    `json:"code"`
	Msg            string `json:"msg"`
	AppAccessToken string `json:"app_access_token"`
	Expire         int    `json:"expire"`
}

type TenantAccessTokenReq struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type TenantAccessTokenResp struct {
	Code              int    `json:"code"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int64  `json:"expire"`
}

type MessagesReq struct {
	ReceiveID string `json:"receive_id"`
	MsgType   string `json:"msg_type"`
	Content   string `json:"content"`
	UUID      string `json:"uuid"`
}

type MessagesResp struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Data Message `json:"data"`
}

type Message struct {
	Code           string      `json:"code"`
	RootID         string      `json:"root_id"`
	ParentID       string      `json:"parent_id"`
	MsgType        string      `json:"msg_type"`
	CreateTime     string      `json:"create_time"`
	UpdateTime     string      `json:"update_time"`
	Deleted        bool        `json:"deleted"`
	Updated        bool        `json:"updated"`
	ChatID         string      `json:"chat_id"`
	Sender         sender      `json:"sender"`
	Body           MessageBody `json:"body"`
	Mentions       []Mentions  `json:"mentions"`
	UpperMessageId string      `json:"upper_message_id"`
}

type sender struct {
	ID         string `json:"id"`
	IDType     string `json:"id_type"`
	SenderType string `json:"sender_type"`
	TenantKey  string `json:"tenant_key"`
}

type MessageBody struct {
	Content string `json:"content"`
}

type Mentions struct {
	Key       string `json:"key"`
	ID        string `json:"id"`
	IDType    string `json:"id_type"`
	Name      string `json:"name"`
	TenantKey string `json:"tenant_key"`
}
