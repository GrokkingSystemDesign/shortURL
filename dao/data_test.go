package dao

import (
	"context"
	"github.com/GrokkingSystemDesign/shortURL/config"
	"testing"
)

func TestInsertData(t *testing.T) {
	config.LoadConf()
	config.InitDB()
	type args struct {
		ctx  context.Context
		data *URLData
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{args: args{ctx: context.TODO(), data: &URLData{Did: "123", Value: "1234"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InsertData(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InsertData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetData(t *testing.T) {
	config.LoadConf()
	config.InitDB()
	type args struct {
		ctx context.Context
		did string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{args: args{ctx: context.TODO(), did: "123"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetData(tt.args.ctx, tt.args.did)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
