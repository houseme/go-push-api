package http

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
    "time"
    
    "github.com/houseme/mipush/internal/builder"
    "github.com/houseme/mipush/internal/miconst"
)

var Client = new(httpClient)

type httpClient struct {
}

// baseHost 请求的 服务域名 默认正式环境
func (hc *httpClient) baseHost(env miconst.RequestEnv) string {
    if env == miconst.SandBoxRequestEnv {
        return miconst.SandboxHost
    }
    return miconst.OfficialHost
}

// DoPost Net Do post request
func (hc *httpClient) DoPost(builder *builder.Builder, params builder.Params) ([]byte, error) {
    form, _ := hc.messageToForm(builder)
    
    req, err := http.NewRequest("POST", fmt.Sprintf("%s", hc.baseHost(params.MiEnv)+params.MiUrl),
        strings.NewReader(form.Encode()))
    
    if err != nil {
        return nil, err
    }
    
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
    req.Header.Add("Authorization", fmt.Sprintf("key=%s", params.AppSecret))
    
    var client = &http.Client{
        Timeout: 5 * time.Second,
    }
    
    res, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return nil, err
    }
    res.Body.Close()
    
    return body, nil
}

// DoGet Net Do get request
func (hc *httpClient) DoGet(builder *builder.Builder, params builder.Params) ([]byte, error) {
    form, _ := hc.messageToForm(builder)
    
    req, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", hc.baseHost(params.MiEnv)+params.MiUrl, form.Encode()), nil)
    
    if err != nil {
        return nil, err
    }
    
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
    req.Header.Set("Authorization", fmt.Sprintf("key=%s", params.AppSecret))
    
    var client = &http.Client{
        Timeout: 5 * time.Second,
    }
    
    res, err := client.Do(req)
    body, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    
    return body, nil
}

// messageToForm 消息转表单，小米推送接口使用form表单提交
func (hc *httpClient) messageToForm(builder *builder.Builder) (*url.Values, error) {
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
