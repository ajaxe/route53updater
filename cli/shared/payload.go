package shared

type Payload struct {
	Nonce   string `json:"nonce"`
	IP      string
	HashKey string
}
