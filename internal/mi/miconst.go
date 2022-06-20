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
	MessageRegIDURL = "/v4/message/regid"
	// MessageAliasURL .
	MessageAliasURL = "/v3/message/alias"
	// MessageAllURL .
	MessageAllURL = "/v3/message/all"
	// MessageAccountURL .
	MessageAccountURL = "/v2/message/user_account"
	// MessageTopicURL .
	MessageTopicURL = "/v3/message/topic"
	// MessageTopicOpURL .
	MessageTopicOpURL = "/v3/message/multi_topic"

	// MultiRegIDURL .
	MultiRegIDURL = "/v2/multi_messages/regids"
	// MultiAliasURL .
	MultiAliasURL = "/v2/multi_messages/aliases"
	// MultiAccountURL .
	MultiAccountURL = "/v2/multi_messages/user_accounts"
	// StatsURL .
	StatsURL = "/v1/stats/message/counters"
	// MessageStatusURL .
	MessageStatusURL = "/v1/trace/message/status"
	// MessagesStatusURL .
	MessagesStatusURL = "/v1/trace/messages/status"
	// SubscribeURL .
	SubscribeURL = "/v2/topic/subscribe"
	// UnsubscribeURL .
	UnsubscribeURL = "/v2/topic/unsubscribe"
	// SubscribeAliasURL .
	SubscribeAliasURL = "/v2/topic/subscribe/alias"
	// UnsubscribeAliasURL .
	UnsubscribeAliasURL = "/v2/topic/unsubscribe/alias"
	// InvalidRegIdsURL .
	InvalidRegIdsURL = "https://feedback.xmpush.xiaomi.com/v1/feedback/fetch_invalid_regids"

	// RegIDAliasURL .
	RegIDAliasURL = "/v1/alias/all"
	// RegIDTopicURL .
	RegIDTopicURL = "/v1/topic/all"
	// ScheduleJobExistURL .
	ScheduleJobExistURL = "/v2/schedule_job/exist"
	// ScheduleJobDeleteURL .
	ScheduleJobDeleteURL = "/v2/schedule_job/delete"
	// GetRegIdsByUserAccountURL .
	GetRegIdsByUserAccountURL = "/v1/useraccount/get_regids_by_useraccount"
	// MediaUploadImageURL 上传大图API
	MediaUploadImageURL = "/media/upload/image"

	// ChannelAddURL 添加channel
	ChannelAddURL = "/v1/channel/add"

	// ChannelListURL channel list
	ChannelListURL = "/v1/channel/list"

	// ChannelDiscardURL 删除一个channel
	ChannelDiscardURL = "/v1/channel/discard"

	// GetRegionByRegIDURL 批量查询RegId的注册地区
	GetRegionByRegIDURL = "/v1/feedback/get_region_by_regid"

	// StopByIDURL 消息ID停止消息
	StopByIDURL = "/v1/message/switch/stop_by_id"

	// StopByJobKeyURL 消息jobkey停止消息
	StopByJobKeyURL = "/v1/message/switch/stop_by_jobkey"

	// MessageAliasRevokeURL 设备alias撤回消息
	MessageAliasRevokeURL = "/v1/message/alias/revoke"

	// MessageRegIDRevokeURL 设备regId撤回消息
	MessageRegIDRevokeURL = "/v1/message/regid/revoke"

	// MessageUserAccountRevokeURL 设备userAccount撤回消息
	MessageUserAccountRevokeURL = "/v1/message/user_account/revoke"

	// MessageTopicRevokeURL 撤回topic消息
	MessageTopicRevokeURL = "/v1/message/topic/revoke"

	// MessageMultiTopicRevokeURL 撤回多topic消息
	MessageMultiTopicRevokeURL = "/v1/message/multi_topic/revoke"

	// MediaUploadSmallIconQueryURL 查询小图上传状态
	MediaUploadSmallIconQueryURL = "/media/upload/smallIcon/query"
	// MediaUploadSmallIconURL 小图上传
	MediaUploadSmallIconURL = "/media/upload/smallIcon"

	// RegionChina .中国区
	RegionChina Region = "China"
	// RegionEurope .欧洲区
	RegionEurope Region = "Europe"
	// RegionRussia .俄罗斯区
	RegionRussia Region = "Russia"
	// RegionIndia .印度区
	RegionIndia Region = "India"
	// RegionOther .其他区
	RegionOther Region = "Other"

	// SandBoxRequestEnv 请求环境 生成还是测试 沙箱 or 正式
	SandBoxRequestEnv RequestEnv = "sandbox"
	// OfficialRequestEnv 请求环境 生成还是测试 沙箱 or 正式
	OfficialRequestEnv RequestEnv = "official"

	// DefaultTimeOut 请求超时时间
	DefaultTimeOut time.Duration = 10000000000

	// SuccessCode error code 成功
	SuccessCode ErrorCode = 0
	// FailCode error code 失败
	FailCode ErrorCode = 0

	// RegIDAccountType account type 发消息的账户类型
	RegIDAccountType AccountType = "regId"
	// AliasAccountType account type 发消息的账户类型
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
