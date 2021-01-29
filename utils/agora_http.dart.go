package utils

import (
	"fmt"
	"github.com/MbungeApp/mbunge-core/models/request"
	"io/ioutil"
	"net/http"
)

var (
	token     string
	serverUrl = "https://api.agora.io/dev/v1"
	client    = &http.Client{}
)

func IsChannelAvailable(channelId string) (request.ChannelInfo, error) {
	var username string = "467917eb7bc549be8228ba5132f48b99"
	var passwd string = "11ecbc0b2a6f41c69bab77f5779aa743"
	subEndpoint := fmt.Sprintf("channel/user/c8449ba570c04c52b6a20e01c5e7f6ea/%s", channelId)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", serverUrl, subEndpoint), nil)
	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	bodyText, _ := ioutil.ReadAll(resp.Body)
	channelInfo, err := request.UnmarshalChannelInfo(bodyText)
	if err != nil {
		return request.ChannelInfo{}, err
	}
	return channelInfo, nil
}
