package helper

import (
	"encoding/json"

	"github.com/houseme/go-push-api/internal/mi"
)

// ToJSON byte 转化为json 对象
func ToJSON(res []byte) (*mi.Result, error) {
	var r *mi.Result
	err := json.Unmarshal(res, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}
