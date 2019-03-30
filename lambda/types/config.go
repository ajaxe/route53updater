package types

// Config base json configuration type
type Config struct {
	Domains []Domain `json:"domains"`
}

type Domain struct {
	EntryName  string `json:"entryName"`
	RecordType string `json:"recordType"`
}

// AWSConfig comprises of AWS access configuration
type AWSConfig struct {
	AccessKeyID string
	SecretKey   string
	Region      string
}
