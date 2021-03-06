package builder

// Channel .
type Channel struct {
	ChannelID          string `json:"channelId"`
	ChannelName        string `json:"channelName"`
	NotifyType         int32  `json:"notifyType"`
	ChannelDescription string `json:"channelDescription,omitempty"`
	SoundURI           string `json:"soundUri,omitempty"`
}
