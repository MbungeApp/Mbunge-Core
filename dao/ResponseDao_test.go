package dao

import (
	"github.com/MbungeApp/mbunge-core/models/db"
	"reflect"
	"testing"
)

func TestNewResponseDao_AddResponse(t *testing.T) {
	type fields struct {
		Client *mongo.Client
	}
	type args struct {
		response db.Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewResponseDao{
				Client: tt.fields.Client,
			}
			if err := r.AddResponse(tt.args.response); (err != nil) != tt.wantErr {
				t.Errorf("AddResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewResponseDao_DeleteResponse(t *testing.T) {
	type fields struct {
		Client *mongo.Client
	}
	type args struct {
		responseId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewResponseDao{
				Client: tt.fields.Client,
			}
			if err := r.DeleteResponse(tt.args.responseId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewResponseDao_GetAllResponseByParti(t *testing.T) {
	type fields struct {
		Client *mongo.Client
	}
	type args struct {
		participationID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []db.Participation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewResponseDao{
				Client: tt.fields.Client,
			}
			if got := r.GetAllResponseByParti(tt.args.participationID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllResponseByParti() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_responseCollection(t *testing.T) {
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
			if got := responseCollection(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("responseCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}
