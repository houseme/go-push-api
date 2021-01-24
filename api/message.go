package api

import (
    "errors"
    
    "github.com/houseme/mipush/internal/builder"
    "github.com/houseme/mipush/internal/http"
    "github.com/houseme/mipush/internal/json"
    "github.com/houseme/mipush/internal/miconst"
    "github.com/houseme/mipush/internal/result"
)

var Message = new(apiMessage)

type apiMessage struct {
}

// SendMessageByRegIds 通过RegId群推
func (a *apiMessage) SendMessageByRegIds(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    
    if builder.RegistrationId == "" {
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
    
    if builder.Alias == "" {
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
