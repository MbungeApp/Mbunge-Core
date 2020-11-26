package dao

import (
	"github.com/MbungeApp/mbunge-core/models/db"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func TestNewEventDaoInterface_ReadNews(t *testing.T) {
	type fields struct {
		Client *mongo.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    []db.EventNew
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewEventDaoInterface{
				Client: tt.fields.Client,
			}
			got, err := u.ReadNews()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadNews() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadNews() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventCollection(t *testing.T) {
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
			if got := eventCollection(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}
