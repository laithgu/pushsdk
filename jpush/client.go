package jpush

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	cfg    *Config
	client *http.Client
}

func NewClient(cfg *Config) *Client {
	return &Client{
		cfg:    cfg,
		client: &http.Client{},
	}
}

func (c *Client) post(path string, payload interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.cfg.BaseURL, path)
	data, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	auth := base64.StdEncoding.EncodeToString([]byte(c.cfg.AppKey + ":" + c.cfg.MasterSecret))
	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("推送失败: %s", string(bodyBytes))
	}

	return bodyBytes, nil
}
