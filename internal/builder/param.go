/**
 * @Project: go-mi-push
 * @Author: houseme
 * @Description:
 * @File: CommonParam
 * @Version: 1.0.0
 * @Date: 2020/1/30 16:09
 *
 */
package builder

import (
    "time"
    
    "github.com/houseme/mipush/internal/miconst"
)

// Params .
type Params struct {
    AppSecret string
    MiUrl    string
    MiEnv    miconst.RequestEnv
    TimeOut  time.Duration
}
