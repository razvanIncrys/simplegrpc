package server

import (
	"context"
	"github.com/razvanIncrys/simplegrpc/backend/pkg/variable"
	"github.com/razvanIncrys/simplegrpc/pb"
	"log"
)

type Server struct {
	pb.UnimplementedVariableServiceServer
}

func (Server) GetMyVariable(ctx context.Context, req *pb.GetMyVariableRequest) (*pb.GetMyVariableResponse, error) {
	variableName := req.GetVariableName()
	variableValue, err := variable.GetMyVariableFromEnvironment(variableName)
	if err == nil {
		log.Printf("getRequest received must return variable:%s with value:%s  ", variableName, variableValue)
	}
	return &pb.GetMyVariableResponse{VariableValue: variableValue}, err
}

func (Server) ListAllMyVariables(ctx context.Context, req *pb.ListMyAllVariableRequest) (*pb.ListAllMyVariableResponse, error) {
	sysVariables := variable.GetAllMyVariablesFromEnvironment()
	log.Printf("listAllRequest received must return all variables: %v", sysVariables)
	respVariables := make([]*pb.Variables, len(sysVariables))
	for i, variable := range sysVariables {
		respVariables[i] = &pb.Variables{Name: variable.Name, Value: variable.Value}
	}
	return &pb.ListAllMyVariableResponse{Variables: respVariables}, nil
}

func (Server) DeleteAllMyVariables(ctx context.Context, req *pb.DeleteAllMyVariableRequest) (*pb.DeleteAllMyVariableResponse, error) {
	deleted, err := variable.DeleteAllMyVariablesFromEnvironment()
	return &pb.DeleteAllMyVariableResponse{Deleted: deleted}, err
}
