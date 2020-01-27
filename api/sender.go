package api

import (
    "encoding/json"
    "errors"
    "fmt"
    "net/url"
    
    "github.com/houseme/mipush/builder"
    "github.com/houseme/mipush/miconst"
    "github.com/houseme/mipush/result"
    "github.com/houseme/mipush/util"
)

//通过RegId群推
func SendMessageByRegIds(appSecret string, builder *builder.Builder) (*result.Result, error) {
    
    if builder.RegistrationId == "" {
        return nil, errors.New("registration_id is required")
    }
    
    form, _ := messageToForm(builder)
    form.Add("registration_id", builder.RegistrationId) //追加registration_id
    
    res, err := util.DoPost(miconst.MessageRegIdURL, appSecret, form)
    if err != nil {
        return nil, err
    }
    
    var result = &result.Result{}
    err = json.Unmarshal(res, result)
    if err != nil {
        return result, err
    }
    
    return result, nil
}

//通过Alias群推
func SendMessageByRegAliasIds(appSecret string, builder *builder.Builder) (*result.Result, error) {
    
    if builder.Alias == "" {
        return nil, errors.New("alias is required")
    }
    
    form, _ := messageToForm(builder)
    form.Add("alias", builder.Alias) //追加alias
    
    res, err := util.DoPost(miconst.MessageAliasURL, appSecret, form)
    if err != nil {
        return nil, err
    }
    
    var result = &result.Result{}
    err = json.Unmarshal(res, result)
    if err != nil {
        return result, err
    }
    
    return result, nil
}

//给所有人发送消息
func SendMessageAll(appSecret string, builder *builder.Builder) (*result.Result, error) {
    
    form, _ := messageToForm(builder)
    res, err := util.DoPost(miconst.MessageAllURL, appSecret, form)
    var result = &result.Result{}
    err = json.Unmarshal(res, result)
    if err != nil {
        return result, err
    }
    
    return result, nil
}

//消息转表单，小米推送接口使用form表单提交
func messageToForm(builder *builder.Builder) (*url.Values, error) {
    
    form := &url.Values{}
    
    form.Add("restricted_package_name", builder.RestrictedPackageName)
    form.Add("payload", builder.Payload)
    form.Add("title", builder.Title)
    form.Add("description", builder.Description)
    form.Add("notify_type", fmt.Sprintf("%d", builder.NotifyType))
    form.Add("pass_through", fmt.Sprintf("%d", builder.PassThrough))
    
    if builder.NotifyID > 0 {
        form.Add("notify_id", fmt.Sprintf("%d", builder.NotifyID))
    }
    
    if builder.TimeToLive > 0 {
        form.Add("time_to_live", fmt.Sprintf("%d", builder.TimeToLive))
    }
    
    if builder.TimeToSend > 0 {
        form.Add("time_to_send", fmt.Sprintf("%d", builder.TimeToSend))
    }
    
    if builder.Extra != nil {
        for k, v := range builder.Extra {
            form.Add(fmt.Sprintf("extra.%s", k), v)
        }
    }
    
    return form, nil
}
