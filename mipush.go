package pushapi

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
func RevokeMessage(builder *builder.Builder, params mi.Params) (*mi.Result, error) {
	return api.Message.RevokeMessage(builder, params)
}

// DealTopic 处理topic
func DealTopic(builder *builder.Builder, params mi.Params) error {
	return api.Message.DealTopic(builder, params)
}

// SubscribeTopicByRegID 注册topic
func SubscribeTopicByRegID(builder *builder.Builder, params mi.Params) (*mi.Result, error) {
	return api.Topic.SubscribeTopicByRegID(builder, params)
}

// BatchSubscribeTopicByRegIds 批量注册主题
func BatchSubscribeTopicByRegIds(builder *builder.Builder, params mi.Params) (*mi.Result, error) {
	return api.Topic.BatchSubscribeTopicByRegIds(builder, params)
}

// SubscribeTopicByAlias 通过别名订阅topic
func SubscribeTopicByAlias(builder *builder.Builder, params mi.Params) (*mi.Result, error) {
	return api.Topic.SubscribeTopicByAlias(builder, params)
}

// BatchSubscribeTopicByAlias 批量注册主题
func BatchSubscribeTopicByAlias(builder *builder.Builder, params mi.Params) (*mi.Result, error) {
	return api.Topic.BatchSubscribeTopicByAlias(builder, params)
}

// UnSubscribeTopicByRegID 通过regId取消注册topic
func UnSubscribeTopicByRegID(builder *builder.Builder, params mi.Params) (*mi.Result, error) {
	return api.Topic.UnSubscribeTopicByRegID(builder, params)
}
