package mi

import (
	"time"
)

// Region .
type Region string

// RequestEnv .
type RequestEnv string

// ErrorCode .
type ErrorCode int64

// AccountType .
type AccountType string

// Kind .
type Kind string

const (
	// SandboxHost .
	SandboxHost = "https://sandbox.xmpush.xiaomi.com"
	// OfficialHost .
	OfficialHost = "https://api.xmpush.xiaomi.com"
	// GlobalProductionHost .
	GlobalProductionHost = "https://api.xmpush.global.xiaomi.com"

	// MessageRegIDURL .
	MessageRegIDURL   = "/v4/message/regid"
	MessageAliasURL   = "/v3/message/alias"
	MessageAllURL     = "/v3/message/all"
	MessageAccountURL = "/v2/message/user_account"
	MessageTopicURL   = "/v3/message/topic"
	MessageTopicOpURL = "/v3/message/multi_topic"

	MultiRegIdURL   = "/v2/multi_messages/regids"
	MultiAliasURL   = "/v2/multi_messages/aliases"
	MultiAccountURL = "/v2/multi_messages/user_accounts"

	StatsURL          = "/v1/stats/message/counters"
	MessageStatusURL  = "/v1/trace/message/status"
	MessagesStatusURL = "/v1/trace/messages/status"

	SubscribeURL   = "/v2/topic/subscribe"
	UnsubscribeURL = "/v2/topic/unsubscribe"

	SubscribeAliasURL   = "/v2/topic/subscribe/alias"
	UnsubscribeAliasURL = "/v2/topic/unsubscribe/alias"

	InvalidRegIdsURL = "https://feedback.xmpush.xiaomi.com/v1/feedback/fetch_invalid_regids"

	RegIdAliasURL = "/v1/alias/all"
	RegIdTopicURL = "/v1/topic/all"

	ScheduleJobExistURL  = "/v2/schedule_job/exist"
	ScheduleJobDeleteURL = "/v2/schedule_job/delete"

	GetRegIdsByUserAccountURL = "/v1/useraccount/get_regids_by_useraccount"
	// 上传大图API
	MediaUploadImageURL = "/media/upload/image"

	// 添加channel
	ChannelAddURL = "/v1/channel/add"

	// channel list
	ChannelListURL = "/v1/channel/list"

	// 删除一个channel
	ChannelDiscardURL = "/v1/channel/discard"

	// 批量查询RegId的注册地区
	GetRegionByRegIdURL = "/v1/feedback/get_region_by_regid"

	// 消息ID停止消息
	StopByIdURL = "/v1/message/switch/stop_by_id"

	// 消息jobkey停止消息
	StopByJobKeyURL = "/v1/message/switch/stop_by_jobkey"

	// 设备alias撤回消息
	MessageAliasRevokeURL = "/v1/message/alias/revoke"

	// 设备regId撤回消息
	MessageRegIdRevokeURL = "/v1/message/regid/revoke"

	// 设备userAccount撤回消息
	MessageUserAccountRevokeURL = "/v1/message/user_account/revoke"

	// 撤回topic消息
	MessageTopicRevokeURL = "/v1/message/topic/revoke"

	// 撤回多topic消息
	MessageMultiTopicRevokeURL = "/v1/message/multi_topic/revoke"

	MediaUploadSmallIconQueryURL = "/media/upload/smallIcon/query"
	MediaUploadSmallIconUrl      = "/media/upload/smallIcon"

	RegionChina  Region = "China"
	RegionEurope Region = "Europe"
	RegionRussia Region = "Russia"
	RegionIndia  Region = "India"
	RegionOther  Region = "Other"

	// 请求环境 生成还是测试 沙箱 or 正式
	SandBoxRequestEnv  RequestEnv = "sandbox"
	OfficialRequestEnv RequestEnv = "official"

	DefaultTimeOut time.Duration = 10000000000

	// error code

	SuccessCode ErrorCode = 0
	FailCode    ErrorCode = 0

	// account type 发消息的账户类型

	RegIdAccountType      AccountType = "regId"
	AliasAccountType      AccountType = "alias"
	TopicAccountType      AccountType = "topic"
	MultiTopicAccountType AccountType = "multiTopic"
	AllAccountType        AccountType = "all"
	UserAccountType       AccountType = "userAccount"

	// ChannelKind 种类 分为 channel message topic tracer 渠道 消息 主题 跟踪
	ChannelKind Kind = "channel"
	// MessageKind 种类 分为 message topic 消息 主题
	MessageKind Kind = "message"
	// TopicKind 种类 分为 topic multiTopic 主题 多主题
	TopicKind Kind = "topic"
	// Tracer 种类 分为 tracer 渠道
	Tracer Kind = "tracer"
)
