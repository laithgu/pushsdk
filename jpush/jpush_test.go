package jpush

import (
	"log"
	"testing"
)

func TestJPush(t *testing.T) {
	cfg := &Config{
		AppKey:       "your-app-key",
		MasterSecret: "your-master-secret",
		BaseURL:      "https://api.jpush.cn",
	}

	jp := NewJPush(cfg)

	err := jp.PushToAlias("user123", "新订单提醒", "您有一个新订单，请及时处理！")
	if err != nil {
		log.Fatalf("推送失败: %v", err)
	}
}
