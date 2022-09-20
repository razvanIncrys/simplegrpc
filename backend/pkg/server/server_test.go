package server

import (
	"context"
	"github.com/razvanIncrys/simplegrpc/pb"
	"reflect"
	"testing"
)

func TestServer_Get(t *testing.T) {
	type fields struct {
		VariableServiceServer pb.VariableServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.GetRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.GetResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				VariableServiceServer: tt.fields.VariableServiceServer,
			}
			got, err := s.Get(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
