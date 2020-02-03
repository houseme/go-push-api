package http

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
    "time"
    
    "github.com/valyala/fasthttp"
    
    "github.com/houseme/mipush/builder"
    "github.com/houseme/mipush/miconst"
)

//do post request
func DoPost(builder *builder.Builder, params builder.Params) ([]byte, error) {
    if params.HttpType == miconst.FastHttpType {
        return DoPostByFastHttp(builder, params)
    }
    return DoPostByNet(builder, params)
}

//do get request
func DoGet(builder *builder.Builder, params builder.Params) ([]byte, error) {
    if params.HttpType == miconst.FastHttpType {
        return DoGetByFastHttp(builder, params)
    }
    return DoGetByNet(builder, params)
}

//请求的 服务域名 默认正式环境
func baseHost(env miconst.RequestEnv) string {
    if env == miconst.SandBoxRequestEnv {
        return miconst.SandboxHost
    }
    return miconst.OfficialHost
}

//FastHttp Do post request
func DoPostByFastHttp(builder *builder.Builder, params builder.Params) ([]byte, error) {
    
    form, errs := messageToFormForFastHttp(builder)
    if errs != nil {
        fmt.Println("请求失败:", errs.Error())
        return nil, errs
    }
    req := fasthttp.AcquireRequest()
    resp := fasthttp.AcquireResponse()
    defer func() {
        // 用完需要释放资源
        fasthttp.ReleaseResponse(resp)
        fasthttp.ReleaseRequest(req)
    }()
    
    req.Header.SetContentType("application/x-www-form-urlencoded;charset=UTF-8")
    req.Header.SetMethod("GET")
    req.Header.Set("Authorization", fmt.Sprintf("key=%s", params.AppSecret))
    req.SetRequestURI(baseHost(params.MiEnv) + params.MiUrl)
    
    req.SetBody(form.QueryString())
    if params.TimeOut == 0{
        params.TimeOut = miconst.DefaultTimeOut
    }
    
    if err := fasthttp.DoTimeout(req, resp, params.TimeOut); err != nil {
        fmt.Println("请求失败:", err.Error())
        return nil, err
    }
    
    body := resp.Body()
    
    fmt.Println("result:\r\n", string(body))
    return body, nil
}

//FastHttp Do get request
func DoGetByFastHttp(builder *builder.Builder, params builder.Params) ([]byte, error) {
    form, errs := messageToFormForFastHttp(builder)
    if errs != nil {
        fmt.Println("请求失败:", errs.Error())
        return nil, errs
    }
    req := fasthttp.AcquireRequest()
    resp := fasthttp.AcquireResponse()
    defer func() {
        // 用完需要释放资源
        fasthttp.ReleaseResponse(resp)
        fasthttp.ReleaseRequest(req)
    }()
    
    req.Header.SetContentType("application/x-www-form-urlencoded;charset=UTF-8")
    req.Header.SetMethod("GET")
    req.Header.Set("Authorization", fmt.Sprintf("key=%s", params.AppSecret))
    req.SetRequestURI(baseHost(params.MiEnv) + params.MiUrl)
    
    req.SetBody(form.QueryString())
    if params.TimeOut == 0{
        params.TimeOut = miconst.DefaultTimeOut
    }
    if err := fasthttp.DoTimeout(req, resp, params.TimeOut); err != nil {
        fmt.Println("请求失败:", err.Error())
        return nil, err
    }
    
    body := resp.Body()
    
    fmt.Println("result:\r\n", string(body))
    return body, nil
}

//Net Do post request
func DoPostByNet(builder *builder.Builder, params builder.Params) ([]byte, error) {
    form, _ := messageToFormForNet(builder)
    param := ""
    if form != nil {
        param = form.Encode()
    }
    
    //fmt.Println(strings.NewReader(param))
    req, err := http.NewRequest("POST", fmt.Sprintf("%s", baseHost(params.MiEnv)+params.MiUrl), strings.NewReader(param))
    
    if err != nil {
        return nil, err
    }
    
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
    req.Header.Add("Authorization", fmt.Sprintf("key=%s", params.AppSecret))
    
    var client = &http.Client{
        Timeout: 5 * time.Second,
    }
    
    res, err := client.Do(req)
    body, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    
    return body, nil
}

//Net Do get request
func DoGetByNet(builder *builder.Builder, params builder.Params) ([]byte, error) {
    form, _ := messageToFormForNet(builder)
    param := ""
    if form != nil {
        param = form.Encode()
    }
    
    req, err := http.NewRequest("GET",
        fmt.Sprintf("%s?%s", baseHost(params.MiEnv)+params.MiUrl, param),
        nil)
    
    if err != nil {
        return nil, err
    }
    
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
    req.Header.Add("Authorization", fmt.Sprintf("key=%s", params.AppSecret))
    
    var client = &http.Client{
        Timeout: 5 * time.Second,
    }
    
    res, err := client.Do(req)
    body, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    
    return body, nil
}

//消息转表单，小米推送接口使用form表单提交
func messageToFormForNet(builder *builder.Builder) (*url.Values, error) {
    
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

//消息转表单，小米推送接口使用form表单提交
func messageToFormForFastHttp(builder *builder.Builder) (*fasthttp.Args, error) {
    
    form := &fasthttp.Args{}
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
