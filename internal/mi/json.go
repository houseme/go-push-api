package mi

import (
	"encoding/json"
)

// ToJSON byte 转化为json 对象
func ToJSON(res []byte) (*Result, error) {
	var r *Result
	err := json.Unmarshal(res, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}
