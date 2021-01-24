/**
 * @Project: go-mi-push
 * @Author: houseme
 * @Description:
 * @File: topic
 * @Version: 1.0.0
 * @Date: 2020/2/1 16:13
 *
 */
package builder

// Topic .
type Topic struct {
    RegistrationId        string `json:"registration_id,omitempty"`
    Aliases               string `json:"aliases,omitempty"`
    UserAccount           string `json:"user_account,omitempty"`
    Topic                 string `json:"topic"`
    Topics                string `json:"topics"`
    Category              string `json:"category"` //ios使用
    RestrictedPackageName string `json:"restricted_package_name,omitempty"`
}
