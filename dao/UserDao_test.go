package dao

import (
	"github.com/MbungeApp/mbunge-core/models/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func TestNewUserDaoInterface_AddUser(t *testing.T) {
	type fields struct {
		Client *mongo.Client
	}
	type args struct {
		user db.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    db.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUserDaoInterface{
				Client: tt.fields.Client,
			}
			got, err := u.AddUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserDaoInterface_DoesUserExist(t *testing.T) {
	type fields struct {
		Client *mongo.Client
	}
	type args struct {
		phone string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUserDaoInterface{
				Client: tt.fields.Client,
			}
			if got := u.DoesUserExist(tt.args.phone); got != tt.want {
				t.Errorf("DoesUserExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserDaoInterface_GetUserByPhone(t *testing.T) {
	type fields struct {
		Client *mongo.Client
	}
	type args struct {
		phone string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    db.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUserDaoInterface{
				Client: tt.fields.Client,
			}
			got, err := u.GetUserByPhone(tt.args.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByPhone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByPhone() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserDaoInterface_UpdateUser(t *testing.T) {
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
		want    db.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUserDaoInterface{
				Client: tt.fields.Client,
			}
			got, err := u.UpdateUser(tt.args.id, tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserDaoInterface_VerifyUser(t *testing.T) {
	type fields struct {
		Client *mongo.Client
	}
	type args struct {
		userid string
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
			u := NewUserDaoInterface{
				Client: tt.fields.Client,
			}
			if err := u.VerifyUser(tt.args.userid); (err != nil) != tt.wantErr {
				t.Errorf("VerifyUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_findUserById(t *testing.T) {
	type args struct {
		id     primitive.ObjectID
		client *mongo.Client
	}
	tests := []struct {
		name string
		args args
		want db.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findUserById(tt.args.id, tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}
