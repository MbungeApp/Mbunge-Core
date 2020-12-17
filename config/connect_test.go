/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package config

import (
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func TestConnectDB(t *testing.T) {
	tests := []struct {
		name string
		want *mongo.Client
	}{
		{
			name: "database connections",
			want: ConnectDB(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConnectDB(); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
				t.Errorf("ConnectDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorReporter(t *testing.T) {
	type args struct {
		report string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test sentry reporter",
			args: args{
				report: "test report",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestGetKey(t *testing.T) {
	type args struct {
		section string
		key     string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test db config",
			args: args{
				section: "mongodb-dev",
				key:     "url",
			},
			want: "mongodb://localhost:27017",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetKey(tt.args.section, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}
