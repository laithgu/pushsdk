package jpush

type JPush struct {
	client *Client
}

func NewJPush(cfg *Config) *JPush {
	return &JPush{
		client: NewClient(cfg),
	}
}

// 推送通知
func (j *JPush) PushToAlias(alias string, title, content string) error {
	payload := PushPayload{
		Platform: "all",
		Audience: map[string][]string{
			"alias": {alias},
		},
		Notification: &Notification{
			Alert: content,
		},
		Message: &Message{
			Title:      title,
			MsgContent: content,
		},
		Options: &Options{
			ApnsProduction: true,
			TimeToLive:     86400,
		},
	}

	_, err := j.client.post("/v3/push", payload)
	return err
}
