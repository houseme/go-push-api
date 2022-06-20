package jums

import (
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
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
	Data   map[string]any
}

// Message 消息模式
func (u *Jums) Message() *Message {
	return &Message{
		Key:    u.config.Key,
		Secret: u.config.Secret,
		Data:   map[string]any{},
	}
}

// Users 用户模式
type Users struct {
	Key           string
	Secret        string
	AccountKey    string
	AccountSecret string
	Data          map[string]any
	UserID        uint
}

// User 用户模式
func (u *Jums) User() *Users {
	return &Users{
		Key:           u.config.Key,
		Secret:        u.config.Secret,
		AccountKey:    u.config.AccountKey,
		AccountSecret: u.config.AccountSecret,
		Data:          map[string]any{},
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
func Request(url string, key string, secret string, data any) error {
	var err error
	_, body, errs := fiber.Post("https://api.ums.jiguang.cn/"+url).Debug().BasicAuth(key, secret).JSON(data).Bytes()
	if len(errs) > 0 {
		return errs[0]
	}
	var api struct {
		Code    int
		Message string
	}
	err = json.Unmarshal(body, &api)
	if err != nil {
		return errors.New("jums request failed")
	}

	if api.Code != 0 {
		return errors.New(api.Message)
	}
	return nil
}
