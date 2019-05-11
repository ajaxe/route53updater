package shared

// Payload object containing ip information
type Payload struct {
	Nonce   string `json:"nonce"`
	IP      string `json:"ip"`
	HashKey string `json:"hashKey"`
}
