package main

import "github.com/razvanIncrys/simplegrpc/frontend/pkg/webserver"

const webPort = ":8080"

func main() {

	webServer := webserver.NewWebServer()
	webServer.Start(webPort)

}

/*
const port = ":9000"

func main() {

	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	gClient.NewClient(conn)
}
*/
