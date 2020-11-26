package utils

import (
	"reflect"
	"testing"
)

func TestComparePasswordAndHash(t *testing.T) {
	type args struct {
		password    string
		encodedHash string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ComparePasswordAndHash(tt.args.password, tt.args.encodedHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComparePasswordAndHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ComparePasswordAndHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateHash(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateHash(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeHash(t *testing.T) {
	type args struct {
		encodedHash string
	}
	tests := []struct {
		name     string
		args     args
		wantP    *params
		wantSalt []byte
		wantHash []byte
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP, gotSalt, gotHash, err := decodeHash(tt.args.encodedHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("decodeHash() gotP = %v, want %v", gotP, tt.wantP)
			}
			if !reflect.DeepEqual(gotSalt, tt.wantSalt) {
				t.Errorf("decodeHash() gotSalt = %v, want %v", gotSalt, tt.wantSalt)
			}
			if !reflect.DeepEqual(gotHash, tt.wantHash) {
				t.Errorf("decodeHash() gotHash = %v, want %v", gotHash, tt.wantHash)
			}
		})
	}
}

func Test_generateRandomSalt(t *testing.T) {
	type args struct {
		n uint32
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateRandomSalt(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateRandomSalt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateRandomSalt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
