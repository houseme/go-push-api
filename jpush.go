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
func (j *Jums) SendSms(ctx context.Context, tempID int, phone, code, signID string) error {
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

// SendEmail 发送邮件
func (j *Jums) SendEmail(ctx context.Context, tempID int, email, code, signID string) error {
	var adu = []string{email}
	return j.client.Message().To(jums.ToEmail("sms", adu)).Content(jums.MsgSms(signID, tempID, map[string]interface{}{"code": code})).Send()
}

// SendWechat 发送微信
func (j *Jums) SendWechat(ctx context.Context, tempID int, openID, code, signID string) error {
	var adu = []string{openID}
	return j.client.Message().To(jums.ToWechatMp("wechat", adu)).Content(jums.MsgSms(signID, tempID, map[string]interface{}{"code": code})).Send()
}

// SendAPP 发送APP
func (j *Jums) SendAPP(ctx context.Context, tempID int, appID, code, signID string) error {
	var adu = []string{appID}
	return j.client.Message().To(jums.ToApp("app", adu)).Content(jums.MsgSms(signID, tempID, map[string]interface{}{"code": code})).Send()
}

// SendWechatLite 发送微信小程序
func (j *Jums) SendWechatLite(ctx context.Context, tempID int, appID, code, signID string) error {
	var adu = []string{appID}
	return j.client.Message().To(jums.ToWechatLite("app", adu)).Content(jums.MsgSms(signID, tempID, map[string]interface{}{"code": code})).Send()
}

// SendWechatWork 发送企业微信
func (j *Jums) SendWechatWork(ctx context.Context, tempID int, appID, code, signID string) error {
	var adu = []string{appID}
	return j.client.Message().To(jums.ToWechatWork("app", adu)).Content(jums.MsgSms(signID, tempID, map[string]interface{}{"code": code})).Send()
}
