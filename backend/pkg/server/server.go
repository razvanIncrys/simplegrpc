package server

import (
	"context"
	"github.com/razvanIncrys/simplegrpc/backend/pb"
	"github.com/razvanIncrys/simplegrpc/backend/pkg/variable"
	"log"
)

type Server struct {
	pb.UnimplementedVariableServiceServer
}

func (Server) SetMyVariable(ctx context.Context, req *pb.SetMyVariableRequest) (*pb.SetMyVariableResponse, error) {
	var err error
	variableName := req.GetVariableName()
	variableValue := req.GetVariableValue()
	variableValue, err = variable.SetMyVariableToEnvironment(variableName, variableValue)
	if err != nil {
		log.Printf("Error on setting  system variable :%s ", err.Error())
		return nil, err
	}
	log.Printf("SetMyVariable seting variable:%s with value:%s  ", variableName, variableValue)
	return &pb.SetMyVariableResponse{Variable: &pb.Variable{Name: variableName, Value: variableValue}}, err
}

func (Server) GetMyVariable(ctx context.Context, req *pb.GetMyVariableRequest) (*pb.GetMyVariableResponse, error) {
	variableName := req.GetVariableName()
	variableValue, err := variable.GetMyVariableFromEnvironment(variableName)
	if err != nil {
		return nil, err
	}
	log.Printf("GetMyVariable received must return variable:%s with value:%s  ", variableName, variableValue)
	return &pb.GetMyVariableResponse{Variable: &pb.Variable{Name: variableName, Value: variableValue}}, err
}

func (Server) ListAllMyVariables(ctx context.Context, req *pb.ListMyAllVariableRequest) (*pb.ListAllMyVariableResponse, error) {
	sysVariables := variable.GetAllMyVariablesFromEnvironment()
	log.Printf("listAllRequest received must return all variables: %v", sysVariables)
	respVariables := make([]*pb.Variable, len(sysVariables))
	for i, sysVariable := range sysVariables {
		respVariables[i] = &pb.Variable{Name: sysVariable.Name, Value: sysVariable.Value}
	}
	return &pb.ListAllMyVariableResponse{Variables: respVariables}, nil
}

func (Server) DeleteAllMyVariables(ctx context.Context, req *pb.DeleteAllMyVariableRequest) (*pb.DeleteAllMyVariableResponse, error) {
	deleted, err := variable.DeleteAllMyVariablesFromEnvironment()
	return &pb.DeleteAllMyVariableResponse{Deleted: deleted}, err
}
