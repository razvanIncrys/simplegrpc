package server

import (
	"context"
	"github.com/razvanIncrys/simplegrpc/backend/pb"
	"reflect"
	"testing"
)

func TestServer_SetMyVariable(t *testing.T) {
	type fields struct {
		UnimplementedVariableServiceServer pb.UnimplementedVariableServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.SetMyVariableRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.SetMyVariableResponse
		wantErr bool
	}{
		{
			"set a variable with a valid key and value",
			fields{UnimplementedVariableServiceServer: pb.UnimplementedVariableServiceServer{}},
			args{context.Background(), &pb.SetMyVariableRequest{VariableName: "test", VariableValue: "test_value"}},
			&pb.SetMyVariableResponse{Variable: &pb.Variable{Name: "test", Value: "test_value"}}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := Server{
				UnimplementedVariableServiceServer: tt.fields.UnimplementedVariableServiceServer,
			}
			got, err := se.SetMyVariable(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetMyVariable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMyVariable() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetMyVariable(t *testing.T) {
	type fields struct {
		UnimplementedVariableServiceServer pb.UnimplementedVariableServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.GetMyVariableRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.GetMyVariableResponse
		wantErr bool
	}{
		{
			"Get my variable using key",
			fields{UnimplementedVariableServiceServer: pb.UnimplementedVariableServiceServer{}},
			args{context.Background(), &pb.GetMyVariableRequest{VariableName: "test"}},
			&pb.GetMyVariableResponse{Variable: &pb.Variable{Name: "test", Value: "test_value"}}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := Server{
				UnimplementedVariableServiceServer: tt.fields.UnimplementedVariableServiceServer,
			}
			got, err := se.GetMyVariable(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMyVariable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMyVariable() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_ListAllMyVariables(t *testing.T) {
	type fields struct {
		UnimplementedVariableServiceServer pb.UnimplementedVariableServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.ListMyAllVariableRequest
	}
	var variables []*pb.Variable
	variables = append(variables, &pb.Variable{Name: "test", Value: "test_value"})
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ListAllMyVariableResponse
		wantErr bool
	}{
		{
			"list all variables",
			fields{UnimplementedVariableServiceServer: pb.UnimplementedVariableServiceServer{}},
			args{context.Background(), &pb.ListMyAllVariableRequest{}},
			&pb.ListAllMyVariableResponse{Variables: variables}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := Server{
				UnimplementedVariableServiceServer: tt.fields.UnimplementedVariableServiceServer,
			}
			got, err := se.ListAllMyVariables(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAllMyVariables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllMyVariables() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_DeleteAllMyVariables(t *testing.T) {
	type fields struct {
		UnimplementedVariableServiceServer pb.UnimplementedVariableServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.DeleteAllMyVariableRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.DeleteAllMyVariableResponse
		wantErr bool
	}{
		{
			"delete all variables",
			fields{UnimplementedVariableServiceServer: pb.UnimplementedVariableServiceServer{}},
			args{context.Background(), &pb.DeleteAllMyVariableRequest{}},
			&pb.DeleteAllMyVariableResponse{Deleted: true}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := Server{
				UnimplementedVariableServiceServer: tt.fields.UnimplementedVariableServiceServer,
			}
			got, err := se.DeleteAllMyVariables(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteAllMyVariables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteAllMyVariables() got = %v, want %v", got, tt.want)
			}
		})
	}
}
