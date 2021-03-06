package dao

import (
	"github.com/MbungeApp/mbunge-core/config"
	"github.com/MbungeApp/mbunge-core/models/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
	"time"
)

func TestNewDeviceDaoInterface_AddDevice(t *testing.T) {
	type fields struct {
		Client *mongo.Client
	}
	type args struct {
		device db.Device
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    db.Device
		wantErr bool
	}{
		{
			name:   "add device",
			fields: fields{Client: config.ConnectDB()},
			args: args{device: db.Device{
				UserId:    "user_id",
				Type:      "app",
				FCMToken:  "",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDeviceDaoInterface{
				Client: tt.fields.Client,
			}
			got, err := d.AddDevice(tt.args.device)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddDevice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddDevice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDeviceDaoInterface_GetDevice(t *testing.T) {
	type fields struct {
		Client *mongo.Client
	}
	type args struct {
		userId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    db.Device
		wantErr bool
	}{
		{
			name: "get device",
			fields: fields{
				Client: config.ConnectDB(),
			},
			args: args{
				userId: "user_id",
			},
			want: db.Device{
				UserId:    "user_id",
				Type:      "app",
				FCMToken:  "",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDeviceDaoInterface{
				Client: tt.fields.Client,
			}
			got, err := d.GetDevice(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDevice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDevice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDeviceDaoInterface_UpdateDevice(t *testing.T) {
	type fields struct {
		Client *mongo.Client
	}
	type args struct {
		id    string
		key   string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    db.Device
		wantErr bool
	}{

		{
			name: "Update device",
			fields: fields{
				Client: config.ConnectDB(),
			},
			args: args{
				id: "", //TODO
			},
			want: db.Device{
				UserId:    "user_id",
				Type:      "app",
				FCMToken:  "",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDeviceDaoInterface{
				Client: tt.fields.Client,
			}
			got, err := d.UpdateDevice(tt.args.id, tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateDevice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateDevice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deviceCollection(t *testing.T) {
	type args struct {
		client *mongo.Client
	}
	tests := []struct {
		name string
		args args
		want *mongo.Collection
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deviceCollection(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deviceCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDeviceById(t *testing.T) {
	type args struct {
		id     primitive.ObjectID
		client *mongo.Client
	}
	tests := []struct {
		name string
		args args
		want db.Device
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDeviceById(tt.args.id, tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDeviceById() = %v, want %v", got, tt.want)
			}
		})
	}
}
