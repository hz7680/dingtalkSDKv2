package model

type CallBackRequest struct {
	Signature string `form:"signature"`
	TimeStamp string `form:"timestamp"`
	Nonce     string `form:"nonce"`
	Encrypt   string `json:"encrypt"`
}

type CallBackResponse struct {
	MsgSignature string `json:"msg_signature"`
	TimeStamp    string `json:"timeStamp"`
	Nonce        string `json:"nonce"`
	Encrypt      string `json:"encrypt"`
}

type CheckUrlEvent struct {
	EventType string `json:"EventType"`
}