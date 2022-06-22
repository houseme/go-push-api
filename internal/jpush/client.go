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
func (c *Client) Push(po *builder.PushPayload) (string, error) {
	if len(c.appKey) != KeyLength || len(c.masterSecret) != KeyLength {
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
	req.SetBasicAuth(c.appKey, c.masterSecret)

	client := http.Client{
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

	resp, e := client.Do(req)
	if e != nil {
		return resp.Status, e
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	buf := make([]byte, 4096)
	nr, _ := resp.Body.Read(buf)

	if resp.StatusCode == http.StatusOK {
		return string(buf[:nr]), nil
	}

	return resp.Status, fmt.Errorf(string(buf[:nr]))
}
