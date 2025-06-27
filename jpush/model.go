package jpush

type Notification struct {
	Alert string `json:"alert"`
}

type Message struct {
	MsgContent string `json:"msg_content"`
	Title      string `json:"title"`
}

type PushPayload struct {
	Platform     interface{}   `json:"platform"` // "all" or []string{"android", "ios"}
	Audience     interface{}   `json:"audience"` // "all", {"alias": ["user1", "user2"]}
	Notification *Notification `json:"notification,omitempty"`
	Message      *Message      `json:"message,omitempty"`
	Options      *Options      `json:"options,omitempty"`
}

type Options struct {
	ApnsProduction bool `json:"apns_production"`
	TimeToLive     int  `json:"time_to_live"`
}
