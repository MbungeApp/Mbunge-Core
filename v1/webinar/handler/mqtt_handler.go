package handler

import (
	"encoding/json"
	"fmt"
	"github.com/MbungeApp/mbunge-core/models/response"
	"github.com/MbungeApp/mbunge-core/utils"
	"github.com/MbungeApp/mbunge-core/v1/webinar/service"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type WebinarMqtt struct {
	Client               *mqtt.Client
	participationService service.ParticipationService
}

func NewMqttWebinarHandler(mqttClient *mqtt.Client, service service.ParticipationService) WebinarMqtt {
	return WebinarMqtt{
		Client:               mqttClient,
		participationService: service,
	}
}

//  -
func (reg *WebinarMqtt) SystemActions(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))

	// convert to map
	m := make(map[string]string)
	err := json.Unmarshal(msg.Payload(), &m)
	if err != nil {
		fmt.Println(err)
	}
	switch m["comm_event"] {
	case "1":
		channelInfo, err := utils.IsChannelAvailable(m["id"])
		if err != nil {
			fmt.Println(err)
		}
		if channelInfo.Success {
			payload := response.AgoraChannelStatus{
				Success:      channelInfo.Success,
				ChannelExist: channelInfo.Data.ChannelExist,
				Audience:     len(channelInfo.Data.Audience),
				Broadcasters: len(channelInfo.Data.Broadcasters),
			}
			data, _ := json.Marshal(payload)
			client.Publish(fmt.Sprintf("topics/users/%s", m["user_id"]), 0, false, data)
		} else {
			payload := response.AgoraChannelStatus{
				Success:      false,
				ChannelExist: false,
				Audience:     0,
				Broadcasters: 0,
			}
			data, _ := json.Marshal(payload)
			client.Publish(fmt.Sprintf("topics/users/%s", m["user_id"]), 0, false, data)
		}
		break
	case "2":
		// TODO
		break
	}

}
