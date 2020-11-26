package utils

import (
	"reflect"
	"testing"
)

func TestSendMail(t *testing.T) {
	type args struct {
		email   string
		otpCode int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_dynamicTemplateEmail(t *testing.T) {
	type args struct {
		email   string
		otpCode int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dynamicTemplateEmail(tt.args.email, tt.args.otpCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dynamicTemplateEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
