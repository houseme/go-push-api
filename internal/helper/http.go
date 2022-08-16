package helper

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/houseme/go-push-api/internal/mi"
	"github.com/houseme/go-push-api/internal/mi/builder"
)

var (
	hc *httpClient
)

type httpClient struct {
	client *client.Client
}

// Client .http client
func Client() (*httpClient, error) {
	c, err := client.NewClient(
		client.WithDialTimeout(5 * time.Second),
	)
	if err != nil {
		return nil, err
	}
	hc = &httpClient{
		client: c,
	}
	return hc, nil
}

// DoPost .
func (c *httpClient) DoPost(ctx context.Context, builder *builder.Builder, params mi.Params) ([]byte, error) {
	return c.do(ctx, builder, params, consts.MethodPost)
}

// DoGet .
func (c *httpClient) DoGet(ctx context.Context, builder *builder.Builder, params mi.Params) ([]byte, error) {
	return c.do(ctx, builder, params, consts.MethodGet)
}

// Do .
func (c *httpClient) do(ctx context.Context, builder *builder.Builder, params mi.Params, method string) ([]byte, error) {
	form := c.messageToForm(ctx, builder)
	req := &protocol.Request{}
	res := &protocol.Response{}
	req.SetMethod(method)
	req.Header.SetContentTypeBytes([]byte("application/x-www-form-urlencoded;charset=UTF-8"))
	req.Header.Add("Authorization", "key="+params.AppSecret)
	req.SetRequestURI(c.requestURL(ctx, params.MiEnv, params.MiURL))
	req.SetBodyString(form.Encode())
	err := c.client.Do(context.Background(), req, res)
	if err != nil {
		return nil, err
	}
	return res.Body(), nil
}

// requestURL 请求的url
func (c *httpClient) requestURL(_ context.Context, env mi.RequestEnv, url string) string {
	var host = mi.OfficialHost
	if env == mi.SandBoxRequestEnv {
		host = mi.SandboxHost
	}
	return host + url
}

// messageToForm 消息转表单，小米推送接口使用form表单提交
func (c *httpClient) messageToForm(_ context.Context, builder *builder.Builder) *url.Values {
	form := &url.Values{}
	form.Add("restricted_package_name", builder.RestrictedPackageName)
	form.Add("payload", builder.Payload)
	form.Add("title", builder.Title)
	form.Add("description", builder.Description)
	form.Add("notify_type", strconv.FormatInt(builder.NotifyType, 10))
	form.Add("pass_through", strconv.FormatInt(builder.PassThrough, 10))

	if builder.NotifyID > 0 {
		form.Add("notify_id", strconv.FormatInt(builder.NotifyID, 10))
	}

	if builder.TimeToLive > 0 {
		form.Add("time_to_live", strconv.FormatInt(builder.TimeToLive, 10))
	}

	if builder.TimeToSend > 0 {
		form.Add("time_to_send", strconv.FormatInt(builder.TimeToSend, 10))
	}

	if builder.Extra != nil {
		for k, v := range builder.Extra {
			form.Add("extra."+k, v)
		}
	}

	return form
}
