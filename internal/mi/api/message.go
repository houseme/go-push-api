package api

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/houseme/go-push-api/internal/mi"
	"github.com/houseme/go-push-api/internal/mi/builder"
)

var Message = apiMessage{}

type apiMessage struct {
}

// SendMessage .
func (a *apiMessage) SendMessage(builder *builder.Builder, params mi.Params) (*mi.Result, error) {
	var buffer bytes.Buffer
	if params.MiEnv == mi.SandBoxRequestEnv {
		buffer.WriteString(mi.SandboxHost)
	} else {
		buffer.WriteString(mi.OfficialHost)
	}

	form := &url.Values{}
	form.Add("restricted_package_name", builder.RestrictedPackageName)
	form.Add("payload", builder.Payload)
	form.Add("title", builder.Title)
	form.Add("description", builder.Description)
	form.Add("notify_type", fmt.Sprintf("%d", builder.NotifyType))
	form.Add("pass_through", fmt.Sprintf("%d", builder.PassThrough))

	if builder.NotifyID > 0 {
		form.Add("notify_id", fmt.Sprintf("%d", builder.NotifyID))
	}

	if builder.TimeToLive > 0 {
		form.Add("time_to_live", fmt.Sprintf("%d", builder.TimeToLive))
	}

	if builder.TimeToSend > 0 {
		form.Add("time_to_send", fmt.Sprintf("%d", builder.TimeToSend))
	}

	if builder.Extra != nil {
		for k, v := range builder.Extra {
			form.Add(fmt.Sprintf("extra.%s", k), v)
		}
	}
	header := make(http.Header)
	header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	header.Add("Authorization", "key=%s"+params.AppSecret)

	switch params.AccountType {
	case mi.RegIdAccountType:
		buffer.WriteString(mi.MessageRegIdURL)
		if len(builder.RegistrationId) == 0 {
			return nil, errors.New("registration_id is required")
		}
		header.Add("registration_id", strings.Join(builder.RegistrationId, ","))
	case mi.AliasAccountType:
		buffer.WriteString(mi.MessageAliasURL)
		if len(builder.Alias) == 0 {
			return nil, errors.New("registration_id is required")
		}
		header.Add("registration_id", strings.Join(builder.Alias, ","))
	case mi.UserAccountType:
		buffer.WriteString(mi.MessageAccountURL)
		if len(builder.Alias) == 0 {
			return nil, errors.New("registration_id is required")
		}
		header.Add("user_account", strings.Join(builder.UserAccount, ","))
	case mi.TopicAccountType:
		buffer.WriteString(mi.MessageTopicURL)
		if len(builder.Alias) == 0 {
			return nil, errors.New("topic is required")
		}
		header.Add("topic", builder.Topic)
	default:
		buffer.WriteString(mi.MessageAllURL)
	}

	req, err := http.NewRequest("POST", buffer.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	var client = &http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Close()

	return mi.ToJSON(body)

}
