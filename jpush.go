package push

import (
	"context"

	"github.com/houseme/go-push-api/internal/jums"
)

// Jums 推送客户端
type Jums struct {
	ctx    context.Context
	client *jums.Jums
}

// SendSms 发送短信
func (j *Jums) SendSms(tempID int, phone, code, signID string) error {
	var adu = []string{phone}
	return j.client.Message().To(jums.ToSms("sms", adu)).Content(jums.MsgSms(signID, tempID, map[string]interface{}{"code": code})).Send()
}

// NewClient 创建推送客户端
func NewClient(ctx context.Context, key, secret string) *Jums {
	return &Jums{
		ctx: ctx,
		client: jums.New(jums.Config{
			Key:    key,
			Secret: secret,
		}),
	}
}

// SendPush 发送推送
func (j *Jums) SendPush(ctx context.Context, client *jums.Jums, title, content, url string, aud []string) error {
	return nil
}
