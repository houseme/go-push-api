package jums

import "github.com/gofrs/uuid"

// ToData 发送消息对象
type ToData map[string]any

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
type MsgData map[string]any

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
		"aud_app": []map[string]any{
			{
				"instance": instance,
				"data":     aud,
			},
		},
	}
}

// ToWechatoa 微信公众号
func ToWechatoa(instance string, aud []string) *ToData {
	return &ToData{
		"aud_wechatoa": []map[string]any{
			{
				"instance": instance,
				"data":     aud,
			},
		},
	}
}

// ToWechatmp 微信小程序
func ToWechatmp(instance string, aud []string) *ToData {
	return &ToData{
		"aud_wechatmp": []map[string]any{
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
		"aud_sms": []map[string]any{
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
		"aud_email": []map[string]any{
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
		"aud_alipay_life": []map[string]any{
			{
				"instance": instance,
				"data":     aud,
			},
		},
	}
}

// ToDingtalkCC 钉钉工作
func ToDingtalkCC(instance string, aud []string) *ToData {
	return &ToData{
		"aud_dingtalk_cc": []map[string]any{
			{
				"instance": instance,
				"data":     aud,
			},
		},
	}
}

// ToWechatwk 企业微信
func ToWechatwk(instance string, aud []string) *ToData {
	return &ToData{
		"aud_wechatwk": []map[string]any{
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
		"msg_app": []map[string]any{
			{
				"cid":      id.String(),
				"platform": "all",
				"notification": map[string]any{
					"android": map[string]any{
						"alert": msg,
						"extras": map[string]any{
							"url": url,
						},
					},
					"ios": map[string]any{
						"alert": msg,
						"extras": map[string]any{
							"url": url,
						},
					},
					"quick": map[string]any{
						"alert": msg,
						"extras": map[string]any{
							"url": url,
						},
					},
				},
			},
		},
	}
}

type WechatoaData struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type WechatMiniprogram struct {
	AppId    string `json:"app_id"`
	PagePath string `json:"page_path"`
}

// MsgWechatoa 公众号消息
func MsgWechatoa(types uint, tplId string, data map[string]WechatoaData, url string, minis ...WechatMiniprogram) *MsgData {
	var mini WechatMiniprogram
	if len(minis) > 0 {
		mini = minis[0]
	}
	return &MsgData{
		"msg_wechatoa": []map[string]any{
			{
				"type":        types,
				"template_id": tplId,
				"url":         url,
				"miniprogram": mini,
				"data":        data,
			},
		},
	}
}

type WechatmpData struct {
	Value string `json:"value"`
}

// MsgWechatmp 小程序消息
func MsgWechatmp(tplId string, data map[string]WechatmpData, pages ...string) *MsgData {
	var page string
	if len(pages) > 0 {
		page = pages[0]
	}
	return &MsgData{
		"msg_wechatmp": []map[string]any{
			{
				"template_id":       tplId,
				"page":              page,
				"miniprogram_state": "formal",
				"data":              data,
			},
		},
	}
}

// MsgSms 短信消息
func MsgSms(signId string, tempId int, params map[string]any) *MsgData {
	return &MsgData{
		"msg_sms": []map[string]any{
			{
				"sign_id":   signId,
				"temp_id":   tempId,
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
		"msg_email": []map[string]any{
			{
				"subject": title,
				"text":    content,
				"files":   file,
			},
		},
	}
}

type AlipayData struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

// MsgAlipayLife 阿里生活号
func MsgAlipayLife(tempId int, headColor string, url string, desc string, data map[string]AlipayData) *MsgData {
	result := MsgData{
		"template_id": tempId,
		"context": map[string]any{
			"head_color":  headColor,
			"url":         url,
			"action_name": desc,
		},
	}
	for key, value := range data {
		result[key] = value
	}

	return &MsgData{
		"msg_wechatwk": []map[string]any{
			result,
		},
	}
}

// MsgDingtalkCC 钉钉通知
func MsgDingtalkCC(content string) *MsgData {
	return &MsgData{
		"msg_dingtalk_cc": []map[string]any{
			{
				"msg": map[string]any{
					"msgtype": "text",
					"text":    content,
				},
			},
		},
	}
}

type wechatWkType string

const (
	WechatWkText  wechatWkType = "text"
	WechatWkImage wechatWkType = "image"
	WechatWkFile  wechatWkType = "file"
	WechatWkNews  wechatWkType = "news"
	WechatWkMpews wechatWkType = "mpnews"
)

type WechatWkConfigText struct {
	Content string `json:"content"`
}

type WechatWkConfigImage struct {
	MediaId string `json:"media_id"`
}

type WechatWkConfigFile struct {
	MediaId string `json:"media_id"`
}

type WechatWkConfigNews struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Picurl      string `json:"picurl"`
}

type WechatWkConfigMpnews struct {
	thumbMediaId     string `json:"thumb_media_id"`
	Title            string `json:"title"`
	Content          string `json:"content"`
	Digest           string `json:"digest"`
	Author           string `json:"author"`
	ContentSourceUrl string `json:"content_source_url"`
}

// MsgWechatwk 企业微信
func MsgWechatwk(config any) *MsgData {
	result := map[string]any{
		"enable_duplicate_check":   1,
		"duplicate_check_interval": 1800,
	}
	switch (interface{})(config).(type) {
	case WechatWkConfigText:
		data := config.(WechatWkConfigText)
		result["msgtype"] = "text"
		result["text"] = map[string]any{
			"content": data.Content,
		}
	case WechatWkConfigImage:
		data := config.(WechatWkConfigImage)
		result["msgtype"] = "image"
		result["image"] = map[string]any{
			"media_id": data.MediaId,
		}
	case WechatWkConfigFile:
		data := config.(WechatWkConfigFile)
		result["msgtype"] = "file"
		result["file"] = map[string]any{
			"media_id": data.MediaId,
		}
	case []WechatWkConfigNews:
		data := config.([]WechatWkConfigNews)
		result["msgtype"] = "news"
		result["news"] = map[string]any{
			"articles": data,
		}
	case []WechatWkConfigMpnews:
		data := config.([]WechatWkConfigNews)
		result["msgtype"] = "mpnews"
		result["mpnews"] = map[string]any{
			"articles": data,
		}
	}
	return &MsgData{
		"msg_wechatwk": []map[string]any{
			result,
		},
	}
}
