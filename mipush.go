package mipush

import (
    "github.com/houseme/mipush/api"
    "github.com/houseme/mipush/internal/builder"
    "github.com/houseme/mipush/internal/result"
)

// SendMessage 发送消息
func SendMessage(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    return api.Message.SendMessage(builder, params)
}

// RevokeMessage 消息退回
func RevokeMessage(builder *builder.Builder, params builder.Params) {

}

// DealTopic 处理topic
func DealTopic() {

}
