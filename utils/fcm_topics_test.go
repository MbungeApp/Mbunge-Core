package utils

import "testing"

func TestSendNotification(t *testing.T) {
	type args struct {
		topic string
		data  interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				topic: "/topics/all",
				data: map[string]string{
					"msg": "Hello World1",
					"sum": "Happy Day",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendNotification(tt.args.topic, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SendNotification() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
