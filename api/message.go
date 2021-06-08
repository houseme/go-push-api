package api

import (
    "errors"
    
    "github.com/houseme/mipush/internal/builder"
    "github.com/houseme/mipush/internal/http"
    "github.com/houseme/mipush/internal/json"
    "github.com/houseme/mipush/internal/miconst"
    "github.com/houseme/mipush/internal/result"
)

var Message = apiMessage{}

type apiMessage struct {
}

// SendMessageByRegIds 通过RegId群推
func (a *apiMessage) SendMessageByRegIds(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    
    if len(builder.RegistrationId) == 0 {
        return nil, errors.New("registration_id is required")
    }
    params.MiUrl = miconst.MessageRegIdURL
    res, err := http.Client.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    
    return json.ToJson(res)
}

// SendMessageByRegAliasIds 通过Alias群推
func (a *apiMessage) SendMessageByRegAliasIds(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    
    if len(builder.Alias) == 0 {
        return nil, errors.New("alias is required")
    }
    params.MiUrl = miconst.MessageAliasURL
    res, err := http.Client.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    
    return json.ToJson(res)
}

// SendMessageAll 给所有人发送消息
func (a *apiMessage) SendMessageAll(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    params.MiUrl = miconst.MessageAllURL
    res, err := http.Client.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    return json.ToJson(res)
}

// SendMessageByRegUserAccounts 通过userAccount 发送消息
func (a *apiMessage) SendMessageByRegUserAccounts(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    params.MiUrl = miconst.MessageAccountURL
    res, err := http.Client.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    return json.ToJson(res)
}

// SendMessageByTopic 通过Topic 发送消息
func (a *apiMessage) SendMessageByTopic(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    params.MiUrl = miconst.MessageTopicURL
    res, err := http.Client.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    return json.ToJson(res)
}

// SendMessageByMultiTopic 通过MultiTopic 发送消息
func (a *apiMessage) SendMessageByMultiTopic(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    params.MiUrl = miconst.MessageTopicOpURL
    res, err := http.Client.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    return json.ToJson(res)
}

// SendMessage .
func (a *apiMessage) SendMessage(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    switch params.AccountType {
    case miconst.RegIdAccountType:
        return a.SendMessageByRegIds(builder, params)
    case miconst.AliasAccountType:
        return a.SendMessageByRegAliasIds(builder, params)
    case miconst.UserAccountType:
        return a.SendMessageByRegUserAccounts(builder, params)
    case miconst.TopicAccountType:
        return a.SendMessageByTopic(builder, params)
    default:
        return a.SendMessageAll(builder, params)
    }
}
