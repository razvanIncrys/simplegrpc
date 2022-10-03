package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/razvanIncrys/simplegrpc/frontend/pkg/httpClient"
	"github.com/razvanIncrys/simplegrpc/frontend/pkg/webserver/handlers"
)

type WebServer struct {
	*gin.Engine
	//client *gClient.Client
	client *httpClient.Endpoints
}

func NewWebServer(hClient *httpClient.Endpoints) *WebServer {
	return &WebServer{gin.Default(), hClient}
}

func (w *WebServer) Start(port string) error {

	v1 := w.Group("/v1")
	{
		v1.GET("/", handlers.HomePageHandler)
		v1.GET("/delete/:target", handlers.DeleteAllMyVariablesHandler(w.client))
		v1.GET("/list/:target", handlers.ListAllMyVariablesHandler(w.client))
		v1.GET("/variable/:name/:target", handlers.GetMyVariableHandler(w.client))
		v1.GET("/variable/set/:name/:value/:target", handlers.SetMyVariableHandler(w.client))
	}

	//load html file
	w.LoadHTMLFiles("./templates/index.tmpl")

	//static path
	w.Static("assets", "./../assets")

	return w.Run(port)
}

func (w *WebServer) Stop() error {
	return w.Stop()
}
