package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/razvanIncrys/simplegrpc/frontend/pkg/webserver/handlers"
)

type WebServer struct {
	*gin.Engine
}

func NewWebServer() *WebServer {
	return &WebServer{gin.Default()}
}

func (w *WebServer) Start(port string) error {

	v1 := w.Group("/v1")
	{
		v1.GET("/", handlers.HomePageHandler)
		v1.POST("/delete", handlers.DeleteAllMyVariablesHandler)
		v1.POST("/list", handlers.ListAllMyVariablesHandler)
		v1.POST("/variable/:name", handlers.GetMyVariableHandler)
	}

	//load html file
	w.LoadHTMLFiles("./../templates/index.tmpl")

	//static path
	w.Static("assets", "./../assets")

	return w.Run(port)
}
