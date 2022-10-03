package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/razvanIncrys/simplegrpc/frontend/pkg/httpClient"
	"net/http"
)

func ListAllMyVariablesHandler(endpoints *httpClient.Endpoints) gin.HandlerFunc {
	return func(c *gin.Context) {
		target := c.Param("target")
		variables, err := endpoints.ListAllMyVariables(target)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":     "SimpleApp",
			"variables": variables.Variables,
			"error":     err,
		})
	}
}

func DeleteAllMyVariablesHandler(endpoints *httpClient.Endpoints) gin.HandlerFunc {
	return func(c *gin.Context) {
		target := c.Param("target")
		done, err := endpoints.DeleteAllMyVariables(target)
		if err != nil {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title": "SimpleApp",
				"error": err.Error(),
			})
		}
		message := "Deleted all variables"
		if !done {
			message = "No variables to delete"
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":   "SimpleApp",
			"Message": message,
		})
	}
}

func HomePageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "SimpleApp",
	})
}

func SetMyVariableHandler(endpoints *httpClient.Endpoints) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		value := c.Param("value")
		target := c.Param("target")
		variables, err := endpoints.SetMyVariable(name, value, target)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":     "SimpleApp",
			"variables": variables.Variables,
			"error":     err,
		})
	}
}

func GetMyVariableHandler(endpoints *httpClient.Endpoints) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		target := c.Param("target")
		variables, err := endpoints.GetMyVariable(name, target)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":     "SimpleApp",
			"variables": variables.Variables,
			"error":     err,
		})
	}
}
