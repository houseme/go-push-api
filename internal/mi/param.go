package mi

import (
	"time"
)

// Params .
type Params struct {
	AppSecret   string        `json:"appSecret"`
	MiUrl       string        `json:"miUrl"`
	MiEnv       RequestEnv    `json:"miEnv"`
	TimeOut     time.Duration `json:"timeOut"`
	AccountType AccountType   `json:"accountType"`
}
