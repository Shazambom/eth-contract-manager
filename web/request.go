package web

type Event struct {
	Token string `json:"token"`
	IP string `json:"ip"`
	UserAgent string `json:"user_agent"`
	Address string `json:"address"`
	NumAvatars int `json:"num_avatars"`
	TransactionNumber int `json:"transaction_number"`
	Nonce string `json:"nonce"`
}
