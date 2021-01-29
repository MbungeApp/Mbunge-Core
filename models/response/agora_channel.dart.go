package response

type AgoraChannelStatus struct {
	Success      bool `json:"success"`
	ChannelExist bool `json:"channel_exist"`
	Audience     int  `json:"audience"`
	Broadcasters int  `json:"broadcasters"`
}
