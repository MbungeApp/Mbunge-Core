package dao

import (
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/models/response"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func TestNewMPDaoInterface_GetAllMps(t *testing.T) {
	type fields struct {
		Client *mongo.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    []db.MP
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMPDaoInterface{
				Client: tt.fields.Client,
			}
			got, err := m.GetAllMps()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllMps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllMps() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMPDaoInterface_GetMpOfTheWeek(t *testing.T) {
	type fields struct {
		Client *mongo.Client
	}
	tests := []struct {
		name   string
		fields fields
		want   response.MpOftheWeek
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMPDaoInterface{
				Client: tt.fields.Client,
			}
			if got := m.GetMpOfTheWeek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMpOfTheWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mpCollection(t *testing.T) {
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
			if got := mpCollection(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mpCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}
