package jums

import (
	"context"
	"errors"
	"time"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Jums 极光统一推送
type Jums struct {
	config Config
}

// Config 配置
type Config struct {
	Key           string
	Secret        string
	AccountKey    string
	AccountSecret string
}

// New 极光统一推送
func New(config Config) *Jums {
	return &Jums{
		config: config,
	}
}

// Message 普通消息推送
type Message struct {
	Key    string
	Secret string
	Data   map[string]interface{}
}

// Message 消息模式
func (u *Jums) Message() *Message {
	return &Message{
		Key:    u.config.Key,
		Secret: u.config.Secret,
		Data:   map[string]interface{}{},
	}
}

// Users 用户模式
type Users struct {
	Key           string
	Secret        string
	AccountKey    string
	AccountSecret string
	Data          map[string]interface{}
	UserID        uint
}

// User 用户模式
func (u *Jums) User() *Users {
	return &Users{
		Key:           u.config.Key,
		Secret:        u.config.Secret,
		AccountKey:    u.config.AccountKey,
		AccountSecret: u.config.AccountSecret,
		Data:          map[string]interface{}{},
	}
}

// UserDel 批量删除用户
func (u *Jums) UserDel(userid ...uint) error {
	if len(userid) == 0 {
		return errors.New("删除用户为空")
	}
	return Request("v1/user/delete", u.config.AccountKey, u.config.AccountSecret, userid)
}

// Request 请求数据
func Request(url string, key string, secret string, data interface{}) error {
	var (
		jsonByte, err = sonic.Marshal(data)
		c             *client.Client
	)
	if err != nil {
		return err
	}

	if c, err = client.NewClient(
		client.WithDialTimeout(5 * time.Second),
	); err != nil {
		return err
	}
	req := &protocol.Request{}
	res := &protocol.Response{}
	req.SetMethod(consts.MethodPost)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	req.URI().SetUsername(key)
	req.URI().SetPassword(secret)
	req.SetRequestURI("https://api.ums.jiguang.cn/" + url)
	req.SetBody(jsonByte)
	if err = c.Do(context.Background(), req, res); err != nil {
		return err
	}

	var api struct {
		Code    int
		Message string
	}
	if err = sonic.Unmarshal(res.Body(), &api); err != nil {
		return errors.New("jums request failed")
	}

	if api.Code != 0 {
		return errors.New(api.Message)
	}
	return nil
}
