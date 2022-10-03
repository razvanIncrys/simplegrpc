package httpClient

import (
	"github.com/razvanIncrys/simplegrpc/frontend/models"
	"reflect"
	"testing"
)

func TestEndpoints_SetMyVariable(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		variableName  string
		variableValue string
		target        []string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantVariables models.Variables
		wantErr       bool
	}{
		{"Test SetMyVariable", fields{NewClient("http://localhost:8080")}, args{"test", "test", []string{"test_value"}}, models.Variables{[]models.Variable{{Name: "test", Value: "test_value"}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Endpoints{
				client: tt.fields.client,
			}
			gotVariables, err := c.SetMyVariable(tt.args.variableName, tt.args.variableValue, tt.args.target...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetMyVariable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVariables, tt.wantVariables) {
				t.Errorf("SetMyVariable() gotVariables = %v, want %v", gotVariables, tt.wantVariables)
			}
		})
	}
}

func TestEndpoints_GetMyVariable(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		variableName string
		target       []string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantVariables models.Variables
		wantErr       bool
	}{
		{name: "Test GetMyVariable", fields: fields{NewClient("http://localhost:8080")}, args: args{"test", []string{"test_value"}}, wantVariables: models.Variables{Variables: []models.Variable{{Name: "test", Value: "test_value"}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Endpoints{
				client: tt.fields.client,
			}
			gotVariables, err := c.GetMyVariable(tt.args.variableName, tt.args.target...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMyVariable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVariables, tt.wantVariables) {
				t.Errorf("GetMyVariable() gotVariables = %v, want %v", gotVariables, tt.wantVariables)
			}
		})
	}
}

func TestEndpoints_ListAllMyVariables(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		target []string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantVariables models.Variables
		wantErr       bool
	}{
		{name: "Test ListAllMyVariables", fields: fields{NewClient("http://localhost:8080")}, args: args{}, wantVariables: models.Variables{Variables: []models.Variable{{Name: "test", Value: "test_value"}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Endpoints{
				client: tt.fields.client,
			}
			gotVariables, err := c.ListAllMyVariables(tt.args.target...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAllMyVariables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVariables, tt.wantVariables) {
				t.Errorf("ListAllMyVariables() gotVariables = %v, want %v", gotVariables, tt.wantVariables)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		baseURL string
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{"Test NewClient", args{"http://localhost:8080"}, &Client{BaseURL: "http://localhost:8080"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.baseURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEndpoints(t *testing.T) {
	type args struct {
		baseURL string
	}
	tests := []struct {
		name string
		args args
		want *Endpoints
	}{
		{"Test NewEndpoints", args{"http://localhost:8080"}, &Endpoints{client: NewClient("http://localhost:8080")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEndpoints(tt.args.baseURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEndpoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndpoints_DeleteAllMyVariables(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		target []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{name: "Test DeleteAllMyVariables", fields: fields{NewClient("http://localhost:8080")}, args: args{}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Endpoints{
				client: tt.fields.client,
			}
			got, err := c.DeleteAllMyVariables(tt.args.target...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteAllMyVariables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeleteAllMyVariables() got = %v, want %v", got, tt.want)
			}
		})
	}
}
