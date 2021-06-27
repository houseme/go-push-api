package builder

// @Project: go-push-api
// @Author: houseme
// @Description:
// @File: define
// @Version: 1.0.0
// @Date: 2021/6/28 01:12
// @Package builder

// PushPayload .
type PushPayload struct {
	Platform     Platform     `json:"platform"`
	Audience     Audience     `json:"audience"`
	Notification Notification `json:"notification"`
	Message      Message      `json:"message"`
	Options      Options      `json:"options"`
	SMS          SMS          `json:"sms"`
	Cid          string       `json:"cid"`
	Target       string       `json:"target"`
	Custom       map[string]interface{}
}

// Platform .
type Platform struct {
}

// Audience .
type Audience struct {
	All     bool             `json:"all"`
	File    bool             `json:"file"`
	Targets []AudienceTarget `json:"targets"`
}

// Notification .
type Notification struct {
}

// Message .
type Message struct {
	Title         string                 `json:"title"`
	MsgContent    string                 `json:"msgContent"`
	ContentType   string                 `json:"contentType"`
	Extras        map[string]string      `json:"extras"`
	NumberExtras  map[string]int64       `json:"numberExtras"`
	BooleanExtras map[string]bool        `json:"booleanExtras"`
	JSONExtras    map[string]string      `json:"jsonExtras"`
	CustomData    map[string]interface{} `json:"customData"`
}

// Options .
type Options struct {
}

// SMS .
type SMS struct {
}

// AudienceTarget .
type AudienceTarget struct {
	AudienceType string   `json:"audienceType"`
	Values       []string `json:"values"`
}

// Builder .
type Builder struct {
	Title                string                 `json:"title"`
	MsgContent           string                 `json:"msgContent"`
	ContentType          string                 `json:"contentType"`
	Extras               map[string]string      `json:"extrasBuilder"`
	NumberExtrasBuilder  map[string]int64       `json:"numberExtrasBuilder"`
	BooleanExtrasBuilder map[string]bool        `json:"booleanExtrasBuilder"`
	JSONExtrasBuilder    map[string]string      `json:"jsonExtrasBuilder"`
	CustomData           map[string]interface{} `json:"customData"`
}
