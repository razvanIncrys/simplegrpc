package variable

import (
	"reflect"
	"testing"
)

func TestGetAllVariablesFromEnvironment(t *testing.T) {
	tests := []struct {
		name string
		want []Variable
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllVariablesFromEnvironment(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllVariablesFromEnvironment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetVariableFromEnvironment(t *testing.T) {
	type args struct {
		variableName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetVariableFromEnvironment(tt.args.variableName); got != tt.want {
				t.Errorf("GetVariableFromEnvironment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateRandomString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateRandomString(tt.args.n); got != tt.want {
				t.Errorf("generateRandomString() = %v, want %v", got, tt.want)
			}
		})
	}
}
