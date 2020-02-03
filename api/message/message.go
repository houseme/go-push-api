package message

import (
    "errors"
    
    "github.com/houseme/mipush/builder"
    "github.com/houseme/mipush/miconst"
    "github.com/houseme/mipush/result"
    "github.com/houseme/mipush/util/http"
    "github.com/houseme/mipush/util/json"
)

//通过RegId群推
func SendMessageByRegIds(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    
    if builder.RegistrationId == "" {
        return nil, errors.New("registration_id is required")
    }
    params.MiUrl = miconst.MessageRegIdURL
    res, err := http.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    
    var result = &result.Result{}
    return json.ToJson(res, result)
}

//通过Alias群推
func SendMessageByRegAliasIds(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    
    if builder.Alias == "" {
        return nil, errors.New("alias is required")
    }
    params.MiUrl = miconst.MessageAliasURL
    res, err := http.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    
    var result = &result.Result{}
    return json.ToJson(res, result)
}

//给所有人发送消息
func SendMessageAll(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    params.MiUrl = miconst.MessageAllURL
    res, err := http.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    var result = &result.Result{}
    return json.ToJson(res, result)
}

func SendMessageByRegUserAccounts(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    params.MiUrl = miconst.MessageAccountURL
    res, err := http.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    var result = &result.Result{}
    return json.ToJson(res, result)
}
