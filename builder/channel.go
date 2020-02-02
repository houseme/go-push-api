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

type Channel struct {
    ChannelId          string `json:"channel_id"`
    ChannelName        string `json:"channel_name"`
    NotifyType         int32  `json:"notify_type"`
    ChannelDescription string `json:"channel_description,omitempty"`
    SoundUri           string `json:"sound_uri,omitempty"`
}
