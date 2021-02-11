package utils

import (
	"fmt"
	"github.com/NaySoftware/go-fcm"
)

const serverKey = "AAAAUZOkvLk:APA91bGnP9C_YFnXS-D2gzIYl6HC3uRsIMo4zoHJYiAIQPIUTGbD8G_fho4oZT2KM4TOiIRY3nG0fY4YUj9MJQV_k9mIK7JmJaDXBeI7Dcfon0UqGMfi8CQp790ryMT600X3SV-xH3QQ"

func SendNotification(topic string, data interface{}) error {
	c := fcm.NewFcmClient(serverKey)
	c.NewFcmMsgTo(topic, data)
	status, err := c.Send()
	if err != nil {
		return err
	}
	fmt.Println(status.StatusCode)
	return nil
}
