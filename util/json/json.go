/**
 * @Project: mipush
 * @Author: qun
 * @Description:
 * @File: json
 * @Version: 1.0.0
 * @Date: 2020/2/3 09:06
 *
 */
package json

import (
    "github.com/json-iterator/go"
    
    "github.com/houseme/mipush/result"
)

// byte 转化为json 对象
func ToJson(res []byte, result *result.Result) (*result.Result, error) {
    var json = jsoniter.ConfigCompatibleWithStandardLibrary
    err := json.Unmarshal(res, result)
    if err != nil {
        return result, err
    }
    return result, nil
}
