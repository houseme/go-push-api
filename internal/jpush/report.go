package jpush

// @Project: go-push-api
// @Author: houseme
// @Description:
// @File: report
// @Version: 1.0.0
// @Date: 2021/6/29 11:12
// @Package jpush

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

// ReportClient .
type ReportClient struct {
	appKey       string
	masterSecret string
}

// ReportObject .
type ReportObject struct {
	AndroidReceived int
	IosApnsSent     int
	MsgID           int
}

// NewReportClient .
func NewReportClient(appKey, masterSecret string) *ReportClient {
	return &ReportClient{appKey, masterSecret}
}

// GetReport .
func (jpc *ReportClient) GetReport(msgID ...int) (string, error) {
	if len(jpc.appKey) != KeyLength || len(jpc.masterSecret) != KeyLength {
		return "", fmt.Errorf("invalidate appkey/masterSecret")
	}

	b, e := json.Marshal(msgID)
	if e != nil {
		return "", e
	}

	reqEndpoint := "https://report.jpush.cn"
	reqURL := fmt.Sprintf("%s/v2/received?msg_ids=%s", reqEndpoint, string(b[1:len(b)-1]))
	req, e := http.NewRequest("GET", reqURL, nil)
	if e != nil {
		return "NewRequestReport", e
	}

	req.Header.Set("User-Agent", "JPush-API-GO-Client")
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("Accept-Charset", "UTF-8")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(jpc.appKey, jpc.masterSecret)

	c := http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(network, addr, time.Second*ConnTimeout)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(RwTimeout * time.Second))
				return c, nil
			},
		},
	}

	resp, e := c.Do(req)
	if e != nil {
		return resp.Status, e
	}
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	nr, _ := resp.Body.Read(buf)

	if resp.StatusCode == http.StatusOK {
		return string(buf[:nr]), nil
	}

	return resp.Status, fmt.Errorf(string(buf[:nr]))
}

// GetReportObject .
func (jpc *ReportClient) GetReportObject(msgIds ...int) ([]ReportObject, error) {
	report, err := jpc.GetReport(msgIds...)
	if err != nil {
		return nil, err
	}

	var res []ReportObject
	if e := json.Unmarshal([]byte(report), &res); e != nil {
		return nil, e
	}

	return res, nil
}
