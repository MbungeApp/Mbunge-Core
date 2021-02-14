package utils

import (
	"fmt"
	"github.com/NaySoftware/go-fcm"
)

const serverKey = "AAAAUZOkvLk:APA91bHBg_eOYcxTuuoW7LvyezgkcdT1VSgk1pvBKV6SqID36k_iMPkOfuaotTvLhOEBPmkjREDmarDX9PaGd3kkyCET-ieteGaTOlvhofG0oPZeXB0TNlPAV-HvbCAQxhDDT3XX8vRH"

func SendNotification(topic string, data map[string]string) error {
	c := fcm.NewFcmClient(serverKey)
	fmt.Println(data)
	c.NewFcmMsgTo(topic, data)

	status, err := c.Send()
	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}
	return nil
}
