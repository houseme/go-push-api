package mipush

import (
	"github.com/houseme/go-push-api/internal/mi"
	"github.com/houseme/go-push-api/internal/mi/api"
	"github.com/houseme/go-push-api/internal/mi/builder"
)

// SendMessage 发送消息
func SendMessage(builder *builder.Builder, params mi.Params) (*mi.Result, error) {
	return api.Message.SendMessage(builder, params)
}

// RevokeMessage 消息退回
func RevokeMessage(builder *builder.Builder, params mi.Params) {

}

// DealTopic 处理topic
func DealTopic() {

}
