// * @Project: go-push-api
// * @Author: qun
// * @Description:
// * @File: topic
// * @Version: 1.0.0
// * @Date: 2020/1/31 12:27

package api

import (
	"github.com/houseme/go-push-api/internal/helper"
	"github.com/houseme/go-push-api/internal/mi"
	"github.com/houseme/go-push-api/internal/mi/builder"
)

// Topic .
var Topic = apiTopic{}

type apiTopic struct {
}

// SubscribeTopicByRegId 注册topic
func (a *apiTopic) SubscribeTopicByRegId(builder *builder.Builder, params mi.Params) (*mi.Result, error) {
	params.MiUrl = mi.MessageAllURL
	res, err := helper.Client.DoPost(builder, params)
	if err != nil {
		return nil, err
	}
	return mi.ToJSON(res)
}

// BatchSubscribeTopicByRegIds 批量注册主题
func (a *apiTopic) BatchSubscribeTopicByRegIds(builder *builder.Builder, params mi.Params) (*mi.Result, error) {
	params.MiUrl = mi.MessageAllURL
	res, err := helper.Client.DoPost(builder, params)
	if err != nil {
		return nil, err
	}
	return mi.ToJSON(res)
}

// SubscribeTopicByAlias 注册topic
func (a *apiTopic) SubscribeTopicByAlias(builder *builder.Builder, params mi.Params) (*mi.Result, error) {
	params.MiUrl = mi.MessageAllURL
	res, err := helper.Client.DoPost(builder, params)
	if err != nil {
		return nil, err
	}
	return mi.ToJSON(res)
}

// BatchSubscribeTopicByAlias 批量注册主题
func (a *apiTopic) BatchSubscribeTopicByAlias(builder *builder.Builder, params mi.Params) (*mi.Result, error) {
	params.MiUrl = mi.MessageAllURL
	res, err := helper.Client.DoPost(builder, params)
	if err != nil {
		return nil, err
	}
	return mi.ToJSON(res)
}

// UnSubscribeTopicByRegId 通过regId取消注册topic
func (a *apiTopic) UnSubscribeTopicByRegId(builder *builder.Builder, params mi.Params) (*mi.Result, error) {
	params.MiUrl = mi.MessageAllURL
	res, err := helper.Client.DoPost(builder, params)
	if err != nil {
		return nil, err
	}
	return mi.ToJSON(res)
}

// BatchUnSubscribeTopicByRegIds 通过regId批量取消注册主题
func (a *apiTopic) BatchUnSubscribeTopicByRegIds(builder *builder.Builder, params mi.Params) (*mi.Result, error) {
	params.MiUrl = mi.MessageAllURL
	res, err := helper.Client.DoPost(builder, params)
	if err != nil {
		return nil, err
	}
	return mi.ToJSON(res)
}

// UnSubscribeTopicByAlias 通过Alias取消注册topic
func (a *apiTopic) UnSubscribeTopicByAlias(builder *builder.Builder, params mi.Params) (*mi.Result, error) {
	params.MiUrl = mi.MessageAllURL
	res, err := helper.Client.DoPost(builder, params)
	if err != nil {
		return nil, err
	}
	return mi.ToJSON(res)
}

// BatchUnSubscribeTopicByAlias 通过Alias批量取消主题
func (a *apiTopic) BatchUnSubscribeTopicByAlias(builder *builder.Builder, params mi.Params) (*mi.Result, error) {
	params.MiUrl = mi.MessageAllURL
	res, err := helper.Client.DoPost(builder, params)
	if err != nil {
		return nil, err
	}
	return mi.ToJSON(res)
}
