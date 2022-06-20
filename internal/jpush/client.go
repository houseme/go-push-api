package jpush

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/houseme/go-push-api/internal/jpush/builder"
)

// @Project: go-push-api
// @Author: houseme
// @Description:
// @File: client
// @Version: 1.0.0
// @Date: 2021/6/29 11:14
// @Package jpush

const (
	// KeyLength .
	KeyLength = 24
	// ConnTimeout .
	ConnTimeout = 5
	// RwTimeout .
	RwTimeout = 30
)

// Client .
type Client struct {
	appKey       string
	masterSecret string
}

// NewClient .
func NewClient(appKey, masterSecret string) *Client {
	return &Client{appKey, masterSecret}
}

// Push .
func (jc *Client) Push(po *builder.PushPayload) (string, error) {
	if len(jc.appKey) != KeyLength || len(jc.masterSecret) != KeyLength {
		return "", fmt.Errorf("invalidate appkey/masterSecret")
	}

	poster, e := po.ToJSONString()
	if e != nil {
		return poster, e
	}

	body := strings.NewReader(poster)
	req, e := http.NewRequest("POST", "https://api.jpush.cn/v3/push", body)
	if e != nil {
		return "NewRequest", e
	}

	req.Header.Set("User-Agent", "JPush-API-GO-Client")
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("Accept-Charset", "UTF-8")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(jc.appKey, jc.masterSecret)

	c := http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(network, addr, time.Second*ConnTimeout)
				if err != nil {
					return nil, err
				}
				if err = c.SetDeadline(time.Now().Add(RwTimeout * time.Second)); err != nil {
					return nil, err
				}
				return c, nil
			},
		},
	}

	resp, e := c.Do(req)
	defer resp.Body.Close()

	if e != nil {
		return resp.Status, e
	}

	buf := make([]byte, 4096)
	nr, _ := resp.Body.Read(buf)

	if resp.StatusCode == http.StatusOK {
		return string(buf[:nr]), nil
	}

	return resp.Status, fmt.Errorf(string(buf[:nr]))
}
