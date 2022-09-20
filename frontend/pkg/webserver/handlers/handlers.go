package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomePageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "SimpleApp",
	})
}

func GetMyVariableHandler(c *gin.Context) {
	c.JSON(200, gin.H{"hello": "world"})
}

func ListAllMyVariablesHandler(c *gin.Context) {
	c.JSON(200, gin.H{"hello": "world"})
}

func DeleteAllMyVariablesHandler(c *gin.Context) {
	c.JSON(200, gin.H{"hello": "world"})
}
