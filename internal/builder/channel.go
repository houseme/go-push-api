/**
 * @Project: go-mi-push
 * @Author: houseme
 * @Description:
 * @File: channel
 * @Version: 1.0.0
 * @Date: 2020/2/1 16:14
 *
 */
package builder

// Channel .
type Channel struct {
    ChannelId          string `json:"channelId"`
    ChannelName        string `json:"channelName"`
    NotifyType         int32  `json:"notifyType"`
    ChannelDescription string `json:"channelDescription,omitempty"`
    SoundUri           string `json:"soundUri,omitempty"`
}
