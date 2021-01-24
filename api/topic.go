/**
 * @Project: mipush
 * @Author: qun
 * @Description:
 * @File: topic
 * @Version: 1.0.0
 * @Date: 2020/1/31 12:27
 *
 */
package api

import (
    "github.com/houseme/mipush/internal/builder"
    "github.com/houseme/mipush/internal/http"
    "github.com/houseme/mipush/internal/json"
    "github.com/houseme/mipush/internal/miconst"
    "github.com/houseme/mipush/internal/result"
)
// Topic .
var Topic = new(apiTopic)

type apiTopic struct {
}

// SubscribeTopicByRegId 注册topic
func (a *apiTopic) SubscribeTopicByRegId(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    params.MiUrl = miconst.MessageAllURL
    res, err := http.Client.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    return json.ToJson(res)
}

// BatchSubscribeTopicByRegIds 批量注册主题
func (a *apiTopic) BatchSubscribeTopicByRegIds(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    params.MiUrl = miconst.MessageAllURL
    res, err := http.Client.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    return json.ToJson(res)
}

// SubscribeTopicByAlias 注册topic
func (a *apiTopic) SubscribeTopicByAlias(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    params.MiUrl = miconst.MessageAllURL
    res, err := http.Client.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    return json.ToJson(res)
}

// BatchSubscribeTopicByAlias 批量注册主题
func (a *apiTopic) BatchSubscribeTopicByAlias(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    params.MiUrl = miconst.MessageAllURL
    res, err := http.Client.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    return json.ToJson(res)
}

// UnSubscribeTopicByRegId 通过regId取消注册topic
func (a *apiTopic) UnSubscribeTopicByRegId(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    params.MiUrl = miconst.MessageAllURL
    res, err := http.Client.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    return json.ToJson(res)
}

// BatchUnSubscribeTopicByRegIds 通过regId批量取消注册主题
func (a *apiTopic) BatchUnSubscribeTopicByRegIds(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    params.MiUrl = miconst.MessageAllURL
    res, err := http.Client.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    return json.ToJson(res)
}

// UnSubscribeTopicByAlias 通过Alias取消注册topic
func (a *apiTopic) UnSubscribeTopicByAlias(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    params.MiUrl = miconst.MessageAllURL
    res, err := http.Client.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    return json.ToJson(res)
}

// BatchUnSubscribeTopicByAlias 通过Alias批量取消主题
func (a *apiTopic) BatchUnSubscribeTopicByAlias(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    params.MiUrl = miconst.MessageAllURL
    res, err := http.Client.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    return json.ToJson(res)
}
