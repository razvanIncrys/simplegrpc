package main

import (
	"github.com/razvanIncrys/simplegrpc/frontend/pkg/httpClient"
	"github.com/razvanIncrys/simplegrpc/frontend/pkg/webserver"
)

const (
	webPort            = ":8080"
	GatewayGRPCPort    = ":9001"
	GatewayGRPCAddress = "http://backend"
)

func main() {

	//gClient := gClient.NewClient(GRPCPort)
	endpoints := httpClient.NewEndpoints(GatewayGRPCAddress + GatewayGRPCPort + "/v1")
	webServer := webserver.NewWebServer(endpoints)
	webServer.Start(webPort)

}
