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

func TestSendMail1(t *testing.T) {
	type args struct {
		email   string
		otpCode int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test random email",
			args: args{
				email:   "858wpwaweru@gmail.com",
				otpCode: 3467,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestSendRandomEmail(t *testing.T) {
	type args struct {
		email   string
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test random email",
			args: args{
				email:   "858wpwaweru@gmail.com",
				message: "test message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_dynamicRandomEmail(t *testing.T) {
	type args struct {
		email   string
		message string
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
			if got := dynamicRandomEmail(tt.args.email, tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dynamicRandomEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dynamicTemplateEmail1(t *testing.T) {
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
