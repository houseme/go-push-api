package builder

import (
	"encoding/json"
	"fmt"
)

// @Project: go-push-api
// @Author: houseme
// @Description:
// @File: define
// @Version: 1.0.0
// @Date: 2021/6/28 01:12
// @Package builder

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

// NamedValue .
type NamedValue map[string]interface{}

// TypePlatform .
type TypePlatform string

const (
	// PlatformAndroid .
	PlatformAndroid TypePlatform = "android"
	// PlatformIOS .
	PlatformIOS TypePlatform = "ios"
	// PlatformWinPhone .
	PlatformWinPhone TypePlatform = "winphone"
	// ALL .
	ALL string = "all"
	// ALIAS .
	ALIAS string = "alias"
	// ALERT .
	ALERT string = "alert"
	// AUDIENCE .
	AUDIENCE string = "audience"
	// ApnsProduction is the apns production.
	ApnsProduction string = "apns_production"
	// BuilderID .
	BuilderID string = "builder_id"
	// ContentType .
	ContentType string = "content_type"
	// EXTRAS .
	EXTRAS string = "extras"
	// IosBadge .
	IosBadge string = "badge"
	// IosSound .
	IosSound string = "sound"
	// IosContentAvailable .
	IosContentAvailable string = "content-available"
	// MaxIosLength .
	MaxIosLength int = 220
	// MaxContentLength of the message.
	MaxContentLength int = 1200
	// MESSAGE .
	MESSAGE string = "message"
	// MsgContent .
	MsgContent string = "msg_content"
	// NOTIFICATION .
	NOTIFICATION  string = "notification"
	OverrideMsgID string = "override_msg_id"
	OPTIONS       string = "options"
	PLATFORM      string = "platform"
	SENDNO        string = "sendno"
	TAG           string = "tag"
	TagAnd        string = "tag_and"
	TITLE         string = "title"
	TTL           string = "time_to_live"
	REGISTRATION  string = "registration_id"
	WpOpenPage    string = "_open_page"
)

// Platform .
type Platform []TypePlatform

// NewPlatform .
func NewPlatform() *Platform {
	t := make(Platform, 0)
	return &t
}

// AddPlatform .
func (vp *Platform) AddPlatform(platform ...TypePlatform) *Platform {
	*vp = append(*vp, platform...)
	return vp
}

// Value .
func (vp *Platform) Value() interface{} {
	if len(*vp) == 0 {
		return ALL
	}

	return vp
}

// Audience .
type Audience map[string][]string

// NewAudience .
func NewAudience() *Audience {
	t := make(Audience)
	return &t
}

// AddCond .
func (va *Audience) AddCond(key, value string) *Audience {
	if _, ok := (*va)[key]; !ok {
		(*va)[key] = make([]string, 0)
	}

	(*va)[key] = append((*va)[key], value)
	return va
}

// AddTag .
func (va *Audience) AddTag(tag string) *Audience {
	return va.AddCond(TAG, tag)
}

// AddTadAnd .
func (va *Audience) AddTadAnd(tagAnd string) *Audience {
	return va.AddCond(TagAnd, tagAnd)
}

// AddAlias .
func (va *Audience) AddAlias(alias string) *Audience {
	return va.AddCond(ALIAS, alias)
}

// AddRegistrationID .
func (va *Audience) AddRegistrationID(id string) *Audience {
	return va.AddCond(REGISTRATION, id)
}

// Value .
func (va *Audience) Value() interface{} {
	if len(*va) == 0 {
		return ALL
	}

	return va
}

// Extras .
type Extras NamedValue

// NewExtras .
func NewExtras() *Extras {
	e := make(Extras)
	return &e
}

// Add .
func (ve *Extras) Add(name string, value interface{}) *Extras {
	(*ve)[name] = value
	return ve
}

// Android .
type Android NamedValue

// NewAndroid .
func NewAndroid(alert string) *Android {
	a := make(Android)
	a[ALERT] = alert
	return &a
}

// AddTitle .
func (va *Android) AddTitle(title string) *Android {
	(*va)[TITLE] = title
	return va
}

// AddBuilderID .
func (va *Android) AddBuilderID(bid int) *Android {
	(*va)[BuilderID] = bid
	return va
}

// AddExtras .
func (va *Android) AddExtras(e *Extras) *Android {
	(*va)[EXTRAS] = e
	return va
}

// iOS .
type iOS NamedValue

// NewIOS .
func NewIOS(alert string) *iOS {
	i := make(iOS)
	i[ALERT] = alert
	i[IosBadge] = 1
	i[IosSound] = ""
	return &i
}

// AddSound .
func (vi *iOS) AddSound(sd string) *iOS {
	(*vi)[IosSound] = sd
	return vi
}

// AddBadge .
func (vi *iOS) AddBadge(bd int) *iOS {
	(*vi)[IosBadge] = bd
	return vi
}

// AddContentAvailable .
func (vi *iOS) AddContentAvailable(b bool) *iOS {
	(*vi)[IosContentAvailable] = b
	return vi
}

// AddExtras .
func (vi *iOS) AddExtras(e *Extras) *iOS {
	(*vi)[EXTRAS] = e
	return vi
}

// WinPhone .
type WinPhone NamedValue

// NewWinPhone .
func NewWinPhone(alert string) *WinPhone {
	w := make(WinPhone)
	w[ALERT] = alert
	return &w
}

// AddTitle .
func (vw *WinPhone) AddTitle(title string) *WinPhone {
	(*vw)[TITLE] = title
	return vw
}

// AddOpenPage .
func (vw *WinPhone) AddOpenPage(op int) *WinPhone {
	(*vw)[WpOpenPage] = op
	return vw
}

// AddExtras .
func (vw *WinPhone) AddExtras(e *Extras) *WinPhone {
	(*vw)[EXTRAS] = e
	return vw
}

// Notification .
type Notification NamedValue

// NewNotification .
func NewNotification(defaultAlert string) *Notification {
	n := make(Notification)
	n[ALERT] = defaultAlert
	return &n
}

// AddAndroid .
func (vn *Notification) AddAndroid(android *Android) *Notification {
	(*vn)[string(PlatformAndroid)] = android
	return vn
}

// AddIOS .
func (vn *Notification) AddIOS(ios *iOS) *Notification {
	(*vn)[string(PlatformIOS)] = ios
	return vn
}

// AddWinPhone .
func (vn *Notification) AddWinPhone(wp *WinPhone) *Notification {
	(*vn)[string(PlatformWinPhone)] = wp
	return vn
}

// Message .
type Message NamedValue

// NewMessage .
func NewMessage(msgContent string) *Message {
	m := make(Message)
	m[MsgContent] = msgContent
	return &m
}

// AddTitle .
func (om *Message) AddTitle(t string) *Message {
	(*om)[TITLE] = t
	return om
}

// AddContentType .
func (om *Message) AddContentType(ct string) *Message {
	(*om)[ContentType] = ct
	return om
}

// AddExtras .
func (om *Message) AddExtras(e *Extras) *Message {
	(*om)[EXTRAS] = e
	return om
}

// Options .
type Options NamedValue

// NewOptions .
func NewOptions() *Options {
	o := make(Options)
	return &o
}

// AddSendNo .
func (vo *Options) AddSendNo(sn int) *Options {
	(*vo)[SENDNO] = sn
	return vo
}

// AddTimeToLive .
func (vo *Options) AddTimeToLive(ttl int) *Options {
	(*vo)[TTL] = ttl
	return vo
}

// AddOverrideMsgID .
func (vo *Options) AddOverrideMsgID(omi int) *Options {
	(*vo)[OverrideMsgID] = omi
	return vo
}

// AddApnsProduction .
func (vo *Options) AddApnsProduction(ap bool) *Options {
	(*vo)[ApnsProduction] = ap
	return vo
}

// PushPayload .
type PushPayload NamedValue

// NewPushPayload .
func NewPushPayload(p *Platform, a *Audience) *PushPayload {
	msg := make(PushPayload)
	msg[PLATFORM] = p.Value()
	msg[AUDIENCE] = a.Value()
	return &msg
}

// Build .
func (vm *PushPayload) Build(key string, comp interface{}) *PushPayload {
	(*vm)[key] = comp
	return vm
}

// AddNotification .
func (vm *PushPayload) AddNotification(n *Notification) *PushPayload {
	return vm.Build(NOTIFICATION, n)
}

// AddMessage .
func (vm *PushPayload) AddMessage(m *Message) *PushPayload {
	return vm.Build(MESSAGE, m)
}

// AddOptions .
func (vm *PushPayload) AddOptions(o *Options) *PushPayload {
	return vm.Build(OPTIONS, o)
}

// ToJSONString .
func (vm *PushPayload) ToJSONString() (string, error) {
	for k, v := range *vm {
		if v == nil {
			return "", fmt.Errorf("%s without a value", k)
		}
	}

	if (*vm)[NOTIFICATION] == nil && (*vm)[MESSAGE] == nil {
		return "", fmt.Errorf("without Notification/Message")
	}

	var length int
	if ntf := (*vm)[NOTIFICATION]; ntf != nil {
		if n, _ := ntf.(*Notification); n != nil {
			if ios := (*n)[string(PlatformIOS)]; ios != nil {
				if b, e := json.Marshal(ios); e != nil || len(b) > MaxIosLength {
					return "", fmt.Errorf("invalidate ios notification")
				}
			}
		}

		b, e := json.Marshal(ntf)
		if e != nil {
			return "", e
		}
		length += len(b)
	}

	if msg := (*vm)[MESSAGE]; msg != nil {
		b, e := json.Marshal(msg)
		if e != nil {
			return "", e
		}
		length += len(b)
	}

	if length > MaxContentLength {
		return "", fmt.Errorf("Notification/Message too large")
	}

	b, e := json.Marshal(vm)
	if e != nil {
		return "", e
	}

	return string(b), e
}
