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
    "encoding/json"
    
    "github.com/houseme/mipush/internal/result"
)

// ToJson byte 转化为json 对象
func ToJson(res []byte) (*result.Result, error) {
    var r *result.Result
    err := json.Unmarshal(res, &r)
    if err != nil {
        return r, err
    }
    return r, nil
}
