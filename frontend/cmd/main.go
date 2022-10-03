package main

import (
	"github.com/razvanIncrys/simplegrpc/frontend/pkg/httpClient"
	"github.com/razvanIncrys/simplegrpc/frontend/pkg/webserver"
)

const (
	webPort            = ":8080"
	GatewayGRPCPort    = ":80"
	GatewayGRPCAddress = "http://10.7.226.225"
)

func main() {

	//gClient := gClient.NewClient(GRPCPort)
	endpoints := httpClient.NewEndpoints(GatewayGRPCAddress + GatewayGRPCPort + "/v1")
	webServer := webserver.NewWebServer(endpoints)
	webServer.Start(webPort)

}
