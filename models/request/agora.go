package request

import "encoding/json"

func UnmarshalChannelInfo(data []byte) (ChannelInfo, error) {
	var r ChannelInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChannelInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ChannelInfo struct {
	Success bool `json:"success"`
	Data    Data `json:"data"`
}

type Data struct {
	ChannelExist  bool          `json:"channel_exist"`
	Mode          int64         `json:"mode"`
	Broadcasters  []int64       `json:"broadcasters"`
	Audience      []interface{} `json:"audience"`
	AudienceTotal int64         `json:"audience_total"`
}
