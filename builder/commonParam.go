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
    "github.com/houseme/mipush/miconst"
)

type CommonParams struct {
    AppSecret string
    MiUrl     string
    HttpType  miconst.HttpType
    MiEnv     miconst.RequestEnv
}
