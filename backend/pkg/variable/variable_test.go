package variable

import (
	"reflect"
	"testing"
)

func TestSetMyVariableToEnvironment(t *testing.T) {
	type args struct {
		variableName  string
		variableValue string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"set my variable to environment", args{"myVariable", "myValue"}, "myValue", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SetMyVariableToEnvironment(tt.args.variableName, tt.args.variableValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetMyVariableToEnvironment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SetMyVariableToEnvironment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMyVariableFromEnvironment(t *testing.T) {
	type args struct {
		variableName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"get my variable to environment", args{"myVariable"}, "myValue", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMyVariableFromEnvironment(tt.args.variableName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMyVariableFromEnvironment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetMyVariableFromEnvironment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllMyVariablesFromEnvironment(t *testing.T) {
	tests := []struct {
		name string
		want []Variable
	}{
		{"get all my variables from environment", []Variable{{"myVariable", "myValue"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllMyVariablesFromEnvironment(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllMyVariablesFromEnvironment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_composeVariableName(t *testing.T) {
	type args struct {
		variableName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"compose variable name", args{"myVariable"}, "SIMPLEGRPC_myVariable"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := composeVariableName(tt.args.variableName); got != tt.want {
				t.Errorf("composeVariableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteAllMyVariablesFromEnvironment(t *testing.T) {
	tests := []struct {
		name    string
		want    bool
		wantErr bool
	}{
		{"delete all my variables from environment", true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteAllMyVariablesFromEnvironment()
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteAllMyVariablesFromEnvironment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeleteAllMyVariablesFromEnvironment() got = %v, want %v", got, tt.want)
			}
		})
	}
}
