package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/razvanIncrys/simplegrpc/frontend/pkg/httpClient"
	"reflect"
	"testing"
)

func TestNewWebServer(t *testing.T) {
	type args struct {
		hClient *httpClient.Endpoints
	}
	tests := []struct {
		name string
		args args
		want *WebServer
	}{
		{name: "Test NewWebServer", args: args{httpClient.NewEndpoints("http://localhost:8080")}, want: &WebServer{Engine: gin.Default(), client: httpClient.NewEndpoints("http://localhost:8080")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWebServer(tt.args.hClient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWebServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWebServer_Start(t *testing.T) {
	type fields struct {
		Engine *gin.Engine
		client *httpClient.Endpoints
	}
	type args struct {
		port string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "Test Start", fields: fields{Engine: gin.Default(), client: httpClient.NewEndpoints("http://localhost:8080")}, args: args{port: "8080"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WebServer{
				Engine: tt.fields.Engine,
				client: tt.fields.client,
			}
			if err := w.Start(tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWebServer_Stop(t *testing.T) {
	type fields struct {
		Engine *gin.Engine
		client *httpClient.Endpoints
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "Test Stop", fields: fields{Engine: gin.Default(), client: httpClient.NewEndpoints("http://localhost:8080")}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WebServer{
				Engine: tt.fields.Engine,
				client: tt.fields.client,
			}
			if err := w.Stop(); (err != nil) != tt.wantErr {
				t.Errorf("Stop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
