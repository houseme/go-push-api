package mipush

import (
    "github.com/houseme/mipush/api/message"
    "github.com/houseme/mipush/builder"
    "github.com/houseme/mipush/result"
)

// 发送消息
func SendMessage(builder *builder.Builder, params builder.CommonParams) (*result.Result, error) {
    return message.SendMessageByRegIds(builder, params)
}

func RevokeMessage(builder *builder.Builder, params builder.CommonParams) {

}

func DealTopic() {

}
