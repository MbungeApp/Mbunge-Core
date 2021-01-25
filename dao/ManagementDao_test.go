package dao

import (
	"github.com/MbungeApp/mbunge-core/config"
	"github.com/MbungeApp/mbunge-core/models/db"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func TestNewManagementDaoInterface_DeleteManager(t *testing.T) {
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
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewManagementDaoInterface{
				Client: tt.fields.Client,
			}
			if err := m.DeleteManager(tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteManager() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewManagementDaoInterface_FindManagerByEmail(t *testing.T) {
	type fields struct {
		Client *mongo.Client
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    db.Management
		wantErr bool
	}{
		{
			name:   "find by email",
			fields: fields{Client: config.ConnectDB()},
			args: args{
				email: "858wpwaweru@gmail.com",
			},
			want: db.Management{
				Name:         "test001",
				NationalID:   "36433941",
				EmailAddress: "858wpwaweru@gmail.com",
				Password:     "$argon2id$v=19$m=65536, t=3, p=2$wcKWaD0gtmoP8ncW0zYSZg$kVVu7ReoVj8j51iLQkYsFdKjuodmBSvcmWQsCVTc7XA",
				Role:         0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewManagementDaoInterface{
				Client: tt.fields.Client,
			}
			got, err := m.FindManagerByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindManagerByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindManagerByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewManagementDaoInterface_InsertManagers(t *testing.T) {
	type fields struct {
		Client *mongo.Client
	}
	type args struct {
		user db.Management
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
			m := NewManagementDaoInterface{
				Client: tt.fields.Client,
			}
			if err := m.InsertManagers(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("InsertManagers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewManagementDaoInterface_ReadManagers(t *testing.T) {
	type fields struct {
		Client *mongo.Client
	}
	tests := []struct {
		name   string
		fields fields
		want   []db.Management
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewManagementDaoInterface{
				Client: tt.fields.Client,
			}
			if got := m.ReadManagers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadManagers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewManagementDaoInterface_UpdateManager(t *testing.T) {
	type fields struct {
		Client *mongo.Client
	}
	type args struct {
		userId string
		key    string
		value  string
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
			m := NewManagementDaoInterface{
				Client: tt.fields.Client,
			}
			if err := m.UpdateManager(tt.args.userId, tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("UpdateManager() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_managersCollection(t *testing.T) {
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
			if got := managersCollection(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("managersCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}
