package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/razvanIncrys/simplegrpc/frontend/pkg/gClient"
	"github.com/razvanIncrys/simplegrpc/frontend/pkg/webserver/handlers"
)

type WebServer struct {
	*gin.Engine
	client *gClient.Client
}

func NewWebServer(gClient *gClient.Client) *WebServer {
	return &WebServer{gin.Default(), gClient}
}

func (w *WebServer) Start(port string) error {

	v1 := w.Group("/v1")
	{
		v1.GET("/", handlers.HomePageHandler)
		v1.POST("/delete", handlers.DeleteAllMyVariablesHandler(w.client))
		v1.POST("/list", handlers.ListAllMyVariablesHandler(w.client))
		v1.POST("/variable/:name", handlers.GetMyVariableHandler(w.client))
	}

	//load html file
	w.LoadHTMLFiles("./../templates/index.tmpl")

	//static path
	w.Static("assets", "./../assets")

	return w.Run(port)
}

func (w *WebServer) Stop() error {
	return w.Stop()
}
