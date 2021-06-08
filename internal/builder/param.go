package builder

import (
    "time"
    
    "github.com/houseme/mipush/internal/miconst"
)

// Params .
type Params struct {
    AppSecret string `json:"appSecret"`
    MiUrl    string `json:"miUrl"`
    MiEnv    miconst.RequestEnv `json:"miEnv"`
    TimeOut  time.Duration `json:"timeOut"`
}
