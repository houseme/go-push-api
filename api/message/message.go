package message

import (
    "errors"
    
    "github.com/json-iterator/go"
    
    "github.com/houseme/mipush/builder"
    "github.com/houseme/mipush/miconst"
    "github.com/houseme/mipush/result"
    "github.com/houseme/mipush/util/http"
)

//通过RegId群推
func SendMessageByRegIds(builder *builder.Builder, params builder.CommonParams) (*result.Result, error) {
    
    if builder.RegistrationId == "" {
        return nil, errors.New("registration_id is required")
    }
    params.MiUrl = miconst.MessageRegIdURL
    res, err := http.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    
    var result = &result.Result{}
    var json = jsoniter.ConfigCompatibleWithStandardLibrary
    err = json.Unmarshal(res, result)
    if err != nil {
        return result, err
    }
    
    return result, nil
}

//通过Alias群推
func SendMessageByRegAliasIds(builder *builder.Builder, params builder.CommonParams) (*result.Result, error) {
    
    if builder.Alias == "" {
        return nil, errors.New("alias is required")
    }
    params.MiUrl = miconst.MessageAliasURL
    res, err := http.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    
    var result = &result.Result{}
    var json = jsoniter.ConfigCompatibleWithStandardLibrary
    err = json.Unmarshal(res, result)
    if err != nil {
        return result, err
    }
    
    return result, nil
}

//给所有人发送消息
func SendMessageAll(builder *builder.Builder, params builder.CommonParams) (*result.Result, error) {
    params.MiUrl = miconst.MessageAllURL
    res, err := http.DoPost(builder, params)
    var result = &result.Result{}
    var json = jsoniter.ConfigCompatibleWithStandardLibrary
    err = json.Unmarshal(res, result)
    if err != nil {
        return result, err
    }
    
    return result, nil
}
