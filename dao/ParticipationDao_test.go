package dao

import (
	"github.com/MbungeApp/mbunge-core/models/db"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func TestNewParticipationDaoInterface_GetAllParticipation(t *testing.T) {
	type fields struct {
		Client *mongo.Client
	}
	tests := []struct {
		name   string
		fields fields
		want   []db.Participation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParticipationDaoInterface{
				Client: tt.fields.Client,
			}
			if got := p.GetAllParticipation(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllParticipation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewParticipationDaoInterface_ParticipationChanges(t *testing.T) {
	type fields struct {
		Client *mongo.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    *mongo.ChangeStream
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParticipationDaoInterface{
				Client: tt.fields.Client,
			}
			got, err := p.ParticipationChanges()
			if (err != nil) != tt.wantErr {
				t.Errorf("ParticipationChanges() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParticipationChanges() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_participationCollection(t *testing.T) {
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
			if got := participationCollection(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("participationCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}
