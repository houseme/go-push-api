package mipush

import (
    "github.com/houseme/mipush/api/message"
    "github.com/houseme/mipush/builder"
    "github.com/houseme/mipush/result"
)

// 发送消息
func SendMessage(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    return message.SendMessageByRegIds(builder, params)
}

// 消息退回
func RevokeMessage(builder *builder.Builder, params builder.Params) {

}

// 处理topic
func DealTopic() {

}
