package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/razvanIncrys/simplegrpc/frontend/pkg/gClient"
	"github.com/razvanIncrys/simplegrpc/pb"
	"net/http"
)

type mVariables struct {
	Name  string
	Value string
}

func ListAllMyVariablesHandler(gClient *gClient.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		gResponse, err := gClient.ListAllMyVariables(context.Background(), &pb.ListMyAllVariableRequest{})
		if err != nil {
			c.JSON(500, gin.H{"error -----------": err.Error()})
		}
		bResp, err := json.Marshal(gResponse.GetVariables())
		if err != nil {
			c.JSON(500, gin.H{"error -----------": err.Error()})
		}
		variables := []mVariables{}
		err = json.Unmarshal(bResp, &variables)
		if err != nil {
			c.JSON(500, gin.H{"error -----------": err.Error()})
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":     "SimpleApp",
			"variables": variables,
		})
	}
}

func DeleteAllMyVariablesHandler(gClient *gClient.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		gResponse, err := gClient.DeleteAllMyVariables(context.Background(), &pb.DeleteAllMyVariableRequest{})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		message := "Deleted all variables"
		if !gResponse.GetDeleted() {
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

func GetMyVariableHandler(gClient *gClient.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		gResponse, err := gClient.GetMyVariable(context.Background(), &pb.GetMyVariableRequest{VariableName: name})
		if err != nil {
			c.JSON(500, gin.H{"error -----------": err.Error()})
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":   "SimpleApp",
			"message": fmt.Sprintf("Variable %s created successfully with value <b>%s</b>", name, gResponse.GetVariableValue()),
		})
	}
}
