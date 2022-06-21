package jums

import "github.com/gofrs/uuid"

// ToData 发送消息对象
type ToData map[string]interface{}

// To 发送消息对象
func (u *Message) To(datas ...*ToData) *Message {
	for _, datum := range datas {
		for k, v := range *datum {
			u.Data[k] = v
		}
	}
	return u
}

// MsgData 消息内容
type MsgData map[string]interface{}

// Content 发送消息内容
func (u *Message) Content(datas ...*MsgData) *Message {
	for _, datum := range datas {
		for k, v := range *datum {
			u.Data[k] = v
		}
	}
	return u
}

// Send 发送消息内容
func (u *Message) Send() error {
	return Request("v1/sent", u.Key, u.Secret, u.Data)
}

// ToUser 用户id
func ToUser(aud []string) *ToData {
	return &ToData{
		"aud_tag": aud,
	}
}

// ToTag 用户标签
func ToTag(aud []string) *ToData {
	return &ToData{
		"aud_userid": aud,
	}
}

// ToSegment 用户群
func ToSegment(aud []string) *ToData {
	return &ToData{
		"aud_segment": aud,
	}
}

// ToApp App
func ToApp(instance string, aud []string) *ToData {
	return &ToData{
		"aud_app": []map[string]interface{}{
			{
				"instance": instance,
				"data":     aud,
			},
		},
	}
}

// ToWechatMp 微信公众号
func ToWechatMp(instance string, aud []string) *ToData {
	return &ToData{
		"aud_wechatoa": []map[string]interface{}{
			{
				"instance": instance,
				"data":     aud,
			},
		},
	}
}

// ToWechatLite 微信小程序
func ToWechatLite(instance string, aud []string) *ToData {
	return &ToData{
		"aud_wechatmp": []map[string]interface{}{
			{
				"instance": instance,
				"data":     aud,
			},
		},
	}
}

// ToSms 短信
func ToSms(instance string, aud []string) *ToData {
	return &ToData{
		"aud_sms": []map[string]interface{}{
			{
				"instance": instance,
				"data":     aud,
			},
		},
	}
}

// ToEmail 邮件
func ToEmail(instance string, aud []string) *ToData {
	return &ToData{
		"aud_email": []map[string]interface{}{
			{
				"instance": instance,
				"data":     aud,
			},
		},
	}
}

// ToAlipayLife 支付宝生活号
func ToAlipayLife(instance string, aud []string) *ToData {
	return &ToData{
		"aud_alipay_life": []map[string]interface{}{
			{
				"instance": instance,
				"data":     aud,
			},
		},
	}
}

// ToDingTalkCC 钉钉工作
func ToDingTalkCC(instance string, aud []string) *ToData {
	return &ToData{
		"aud_dingtalk_cc": []map[string]interface{}{
			{
				"instance": instance,
				"data":     aud,
			},
		},
	}
}

// ToWechatWork 企业微信
func ToWechatWork(instance string, aud []string) *ToData {
	return &ToData{
		"aud_wechatwk": []map[string]interface{}{
			{
				"instance": instance,
				"data":     aud,
			},
		},
	}
}

// MsgApp app消息
func MsgApp(msg string, url string) *MsgData {
	id, _ := uuid.NewV4()
	return &MsgData{
		"msg_app": []map[string]interface{}{
			{
				"cid":      id.String(),
				"platform": "all",
				"notification": map[string]interface{}{
					"android": map[string]interface{}{
						"alert": msg,
						"extras": map[string]interface{}{
							"url": url,
						},
					},
					"ios": map[string]interface{}{
						"alert": msg,
						"extras": map[string]interface{}{
							"url": url,
						},
					},
					"quick": map[string]interface{}{
						"alert": msg,
						"extras": map[string]interface{}{
							"url": url,
						},
					},
				},
			},
		},
	}
}

// WechatMpData 微信公众号消息
type WechatMpData struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

// WechatMiniProgram 小程序参数消息
type WechatMiniProgram struct {
	AppID    string `json:"app_id"`
	PagePath string `json:"page_path"`
}

// MsgWechatMp 公众号消息
func MsgWechatMp(types uint, tplID string, data map[string]WechatMpData, url string, minis ...WechatMiniProgram) *MsgData {
	var mini WechatMiniProgram
	if len(minis) > 0 {
		mini = minis[0]
	}
	return &MsgData{
		"msg_wechatoa": []map[string]interface{}{
			{
				"type":        types,
				"template_id": tplID,
				"url":         url,
				"miniprogram": mini,
				"data":        data,
			},
		},
	}
}

// WechatLiteData 小程序消息
type WechatLiteData struct {
	Value string `json:"value"`
}

// MsgWechatLite 小程序消息
func MsgWechatLite(tplID string, data map[string]WechatLiteData, pages ...string) *MsgData {
	var page string
	if len(pages) > 0 {
		page = pages[0]
	}
	return &MsgData{
		"msg_wechatmp": []map[string]interface{}{
			{
				"template_id":       tplID,
				"page":              page,
				"miniprogram_state": "formal",
				"data":              data,
			},
		},
	}
}

// MsgSms 短信消息
func MsgSms(signID string, tempID int, params map[string]interface{}) *MsgData {
	return &MsgData{
		"msg_sms": []map[string]interface{}{
			{
				"sign_id":   signID,
				"temp_id":   tempID,
				"temp_para": params,
			},
		},
	}
}

// MsgEmail 邮件消息
func MsgEmail(title string, content string, files ...[]string) *MsgData {
	var file []string
	if len(files) > 0 {
		file = files[0]
	}
	return &MsgData{
		"msg_email": []map[string]interface{}{
			{
				"subject": title,
				"text":    content,
				"files":   file,
			},
		},
	}
}

// AlipayData 支付宝生活号数据
type AlipayData struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

// MsgAlipayLife 阿里生活号
func MsgAlipayLife(tempID int, headColor string, url string, desc string, data map[string]AlipayData) *MsgData {
	result := MsgData{
		"template_id": tempID,
		"context": map[string]interface{}{
			"head_color":  headColor,
			"url":         url,
			"action_name": desc,
		},
	}
	for key, value := range data {
		result[key] = value
	}

	return &MsgData{
		"msg_wechatwk": []map[string]interface{}{
			result,
		},
	}
}

// MsgDingTalkCC 钉钉通知
func MsgDingTalkCC(content string) *MsgData {
	return &MsgData{
		"msg_dingtalk_cc": []map[string]interface{}{
			{
				"msg": map[string]interface{}{
					"msgtype": "text",
					"text":    content,
				},
			},
		},
	}
}

type wechatWkType string

const (
	// WechatWkText 文本消息
	WechatWkText wechatWkType = "text"
	// WechatWkImage 图片消息
	WechatWkImage wechatWkType = "image"
	// WechatWkFile 链接消息
	WechatWkFile wechatWkType = "file"
	// WechatWkNews 图文消息
	WechatWkNews wechatWkType = "news"
	// WechatWkMpNews 图文消息
	WechatWkMpNews wechatWkType = "mpnews"
)

// WechatWkConfigText 文本消息
type WechatWkConfigText struct {
	Content string `json:"content"`
}

// WechatWkConfigImage 图片
type WechatWkConfigImage struct {
	MediaID string `json:"media_id"`
}

// WechatWkConfigFile 图文消息
type WechatWkConfigFile struct {
	MediaID string `json:"media_id"`
}

// WechatWkConfigNews 图文消息
type WechatWkConfigNews struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"picurl"`
}

// WechatWkConfigMpNews 图文消息
type WechatWkConfigMpNews struct {
	ThumbMediaID     string `json:"thumb_media_id"`
	Title            string `json:"title"`
	Content          string `json:"content"`
	Digest           string `json:"digest"`
	Author           string `json:"author"`
	ContentSourceURL string `json:"content_source_url"`
}

// MsgWechatWork 企业微信
func MsgWechatWork(config interface{}) *MsgData {
	result := map[string]interface{}{
		"enable_duplicate_check":   1,
		"duplicate_check_interval": 1800,
	}
	switch config.(type) {
	case WechatWkConfigText:
		data := config.(WechatWkConfigText)
		result["msgtype"] = "text"
		result["text"] = map[string]interface{}{
			"content": data.Content,
		}
	case WechatWkConfigImage:
		data := config.(WechatWkConfigImage)
		result["msgtype"] = "image"
		result["image"] = map[string]interface{}{
			"media_id": data.MediaID,
		}
	case WechatWkConfigFile:
		data := config.(WechatWkConfigFile)
		result["msgtype"] = "file"
		result["file"] = map[string]interface{}{
			"media_id": data.MediaID,
		}
	case []WechatWkConfigNews:
		data := config.([]WechatWkConfigNews)
		result["msgtype"] = "news"
		result["news"] = map[string]interface{}{
			"articles": data,
		}
	case []WechatWkConfigMpNews:
		data := config.([]WechatWkConfigNews)
		result["msgtype"] = "mpnews"
		result["mpnews"] = map[string]interface{}{
			"articles": data,
		}
	}
	return &MsgData{
		"msg_wechatwk": []map[string]interface{}{
			result,
		},
	}
}
