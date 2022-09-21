package main

import (
	"github.com/razvanIncrys/simplegrpc/frontend/pkg/gClient"
	"github.com/razvanIncrys/simplegrpc/frontend/pkg/webserver"
)

const (
	webPort  = ":8080"
	GRPCPort = ":9000"
)

func main() {

	gClient := gClient.NewClient(GRPCPort)
	webServer := webserver.NewWebServer(gClient)
	webServer.Start(webPort)

}
