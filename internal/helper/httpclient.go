package helper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/houseme/go-push-api/internal/mi"
	"github.com/houseme/go-push-api/internal/mi/builder"
)

// Client 客户端
var Client = httpClient{}

type httpClient struct {
}

// baseHost 请求的 服务域名 默认正式环境
func (c *httpClient) baseHost(env mi.RequestEnv) string {
	if env == mi.SandBoxRequestEnv {
		return mi.SandboxHost
	}
	return mi.OfficialHost
}

// DoPost Net Do post request
func (c *httpClient) DoPost(builder *builder.Builder, params mi.Params) ([]byte, error) {
	form := c.messageToForm(builder)

	req, err := http.NewRequest("POST", c.baseHost(params.MiEnv)+params.MiURL, strings.NewReader(form.Encode()))

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
	_ = res.Body.Close()

	return body, nil
}

// DoGet Net Do get request
func (c *httpClient) DoGet(builder *builder.Builder, params mi.Params) ([]byte, error) {
	form := c.messageToForm(builder)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", c.baseHost(params.MiEnv)+params.MiURL, form.Encode()), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Set("Authorization", fmt.Sprintf("key=%s", params.AppSecret))

	var client = &http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// messageToForm 消息转表单，小米推送接口使用form表单提交
func (c *httpClient) messageToForm(builder *builder.Builder) *url.Values {
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

	return form
}
