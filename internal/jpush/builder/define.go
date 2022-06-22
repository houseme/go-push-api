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
	NOTIFICATION string = "notification"
	// OverrideMsgID fileter the message by message id.
	OverrideMsgID string = "override_msg_id"
	// OPTIONS .
	OPTIONS string = "options"
	// PLATFORM .
	PLATFORM string = "platform"
	// SENDNO .
	SENDNO string = "sendno"
	// TAG .
	TAG string = "tag"
	// TagAnd .
	TagAnd string = "tag_and"
	// TITLE .
	TITLE string = "title"
	// TTL .
	TTL string = "time_to_live"
	// REGISTRATION .
	REGISTRATION string = "registration_id"
	// WpOpenPage .
	WpOpenPage string = "_open_page"
)

// Platform .
type Platform []TypePlatform

// NewPlatform .
func NewPlatform() *Platform {
	t := make(Platform, 0)
	return &t
}

// AddPlatform .
func (p *Platform) AddPlatform(platform ...TypePlatform) *Platform {
	*p = append(*p, platform...)
	return p
}

// Value .
func (p *Platform) Value() interface{} {
	if len(*p) == 0 {
		return ALL
	}

	return p
}

// Audience .
type Audience map[string][]string

// NewAudience .
func NewAudience() *Audience {
	t := make(Audience)
	return &t
}

// AddCond .
func (a *Audience) AddCond(key, value string) *Audience {
	if _, ok := (*a)[key]; !ok {
		(*a)[key] = make([]string, 0)
	}

	(*a)[key] = append((*a)[key], value)
	return a
}

// AddTag .
func (a *Audience) AddTag(tag string) *Audience {
	return a.AddCond(TAG, tag)
}

// AddTadAnd .
func (a *Audience) AddTadAnd(tagAnd string) *Audience {
	return a.AddCond(TagAnd, tagAnd)
}

// AddAlias .
func (a *Audience) AddAlias(alias string) *Audience {
	return a.AddCond(ALIAS, alias)
}

// AddRegistrationID .
func (a *Audience) AddRegistrationID(id string) *Audience {
	return a.AddCond(REGISTRATION, id)
}

// Value .
func (a *Audience) Value() interface{} {
	if len(*a) == 0 {
		return ALL
	}

	return a
}

// Extras .
type Extras NamedValue

// NewExtras .
func NewExtras() *Extras {
	e := make(Extras)
	return &e
}

// Add .
func (e *Extras) Add(name string, value interface{}) *Extras {
	(*e)[name] = value
	return e
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
func (a *Android) AddTitle(title string) *Android {
	(*a)[TITLE] = title
	return a
}

// AddBuilderID .
func (a *Android) AddBuilderID(bid int) *Android {
	(*a)[BuilderID] = bid
	return a
}

// AddExtras .
func (a *Android) AddExtras(e *Extras) *Android {
	(*a)[EXTRAS] = e
	return a
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
func (i *iOS) AddSound(sd string) *iOS {
	(*i)[IosSound] = sd
	return i
}

// AddBadge .
func (i *iOS) AddBadge(bd int) *iOS {
	(*i)[IosBadge] = bd
	return i
}

// AddContentAvailable .
func (i *iOS) AddContentAvailable(b bool) *iOS {
	(*i)[IosContentAvailable] = b
	return i
}

// AddExtras .
func (i *iOS) AddExtras(e *Extras) *iOS {
	(*i)[EXTRAS] = e
	return i
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
func (w *WinPhone) AddTitle(title string) *WinPhone {
	(*w)[TITLE] = title
	return w
}

// AddOpenPage .
func (w *WinPhone) AddOpenPage(op int) *WinPhone {
	(*w)[WpOpenPage] = op
	return w
}

// AddExtras .
func (w *WinPhone) AddExtras(e *Extras) *WinPhone {
	(*w)[EXTRAS] = e
	return w
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
func (n *Notification) AddAndroid(android *Android) *Notification {
	(*n)[string(PlatformAndroid)] = android
	return n
}

// AddIOS .
func (n *Notification) AddIOS(ios *iOS) *Notification {
	(*n)[string(PlatformIOS)] = ios
	return n
}

// AddWinPhone .
func (n *Notification) AddWinPhone(wp *WinPhone) *Notification {
	(*n)[string(PlatformWinPhone)] = wp
	return n
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
func (m *Message) AddTitle(t string) *Message {
	(*m)[TITLE] = t
	return m
}

// AddContentType .
func (m *Message) AddContentType(ct string) *Message {
	(*m)[ContentType] = ct
	return m
}

// AddExtras .
func (m *Message) AddExtras(e *Extras) *Message {
	(*m)[EXTRAS] = e
	return m
}

// Options .
type Options NamedValue

// NewOptions .
func NewOptions() *Options {
	o := make(Options)
	return &o
}

// AddSendNo .
func (o *Options) AddSendNo(sn int) *Options {
	(*o)[SENDNO] = sn
	return o
}

// AddTimeToLive .
func (o *Options) AddTimeToLive(ttl int) *Options {
	(*o)[TTL] = ttl
	return o
}

// AddOverrideMsgID .
func (o *Options) AddOverrideMsgID(omi int) *Options {
	(*o)[OverrideMsgID] = omi
	return o
}

// AddApnsProduction .
func (o *Options) AddApnsProduction(ap bool) *Options {
	(*o)[ApnsProduction] = ap
	return o
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
func (p *PushPayload) Build(key string, comp interface{}) *PushPayload {
	(*p)[key] = comp
	return p
}

// AddNotification .
func (p *PushPayload) AddNotification(n *Notification) *PushPayload {
	return p.Build(NOTIFICATION, n)
}

// AddMessage .
func (p *PushPayload) AddMessage(m *Message) *PushPayload {
	return p.Build(MESSAGE, m)
}

// AddOptions .
func (p *PushPayload) AddOptions(o *Options) *PushPayload {
	return p.Build(OPTIONS, o)
}

// ToJSONString .
func (p *PushPayload) ToJSONString() (string, error) {
	for k, v := range *p {
		if v == nil {
			return "", fmt.Errorf("%s without a value", k)
		}
	}

	if (*p)[NOTIFICATION] == nil && (*p)[MESSAGE] == nil {
		return "", fmt.Errorf("without Notification/Message")
	}

	var length int
	if ntf := (*p)[NOTIFICATION]; ntf != nil {
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

	if msg := (*p)[MESSAGE]; msg != nil {
		b, e := json.Marshal(msg)
		if e != nil {
			return "", e
		}
		length += len(b)
	}

	if length > MaxContentLength {
		return "", fmt.Errorf("Notification/Message too large")
	}

	b, e := json.Marshal(p)
	if e != nil {
		return "", e
	}

	return string(b), e
}
