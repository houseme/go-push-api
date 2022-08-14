package builder

// Builder .
type Builder struct {
	RegistrationID        []string          `json:"registration_id"`
	Alias                 []string          `json:"alias"`
	UserAccount           []string          `json:"user_account"`
	Topic                 string            `json:"topic"`
	Topics                []string          `json:"topics"`
	TopicOp               string            `json:"topic_op"`
	Payload               string            `json:"payload,omitempty"`
	RestrictedPackageName string            `json:"restricted_package_name,omitempty"`
	Title                 string            `json:"title"`
	Description           string            `json:"description"`
	PassThrough           int64             `json:"pass_through"`          // 0 通知栏消息, 1 透传消息
	NotifyType            int64             `json:"notify_type,omitempty"` // -1: DEFAULT_ALL 1: 使用默认提示音提示, 2: 使用默认震动提示, 4: 使用默认led灯光提示
	TimeToLive            int64             `json:"time_to_live,omitempty"`
	TimeToSend            int64             `json:"time_to_send,omitempty"`
	NotifyID              int64             `json:"notify_id,omitempty"`
	Extra                 map[string]string `json:"extra,omitempty"`
	SoundURI              string            `json:"sound_uri"`
	Ticker                string            `json:"ticker"`
	NotifyForeground      string            `json:"notify_foreground"`
	NotifyEffect          string            `json:"notify_effect"`
	IntentURI             string            `json:"intent_uri"`
	FlowControl           string            `json:"flow_control"`
	AppVersionNotIn       string            `json:"app_version_not_in"`
	Badge                 string            `json:"badge"`    // ios使用
	Category              string            `json:"category"` // ios使用
}
