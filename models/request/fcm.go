package request

import "encoding/json"

func UnmarshalNotification(data []byte) (Notification, error) {
	var r Notification
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Notification) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Notification struct {
	Body  string `json:"body"`
	Title string `json:"title"`
}
