package httpClient

import (
	"encoding/json"
	"github.com/razvanIncrys/simplegrpc/frontend/models"
	"net/url"
)

type Endpoints struct {
	client *Client
}

func NewEndpoints(baseURL string) *Endpoints {
	client := NewClient(baseURL)
	return &Endpoints{client: client}
}

func (c *Endpoints) SetMyVariable(variableName, variableValue string, target ...string) (variables models.Variables, err error) {
	respB, err := c.client.MakeReq("set", url.Values{"variableName": {variableName}, "variableValue": {variableValue}}, "GET", target...)
	if err != nil {
		return
	}
	variableResp := variable{}
	err = json.Unmarshal(respB, &variableResp)
	if err != nil {
		return
	}
	variables.Variables = append(variables.Variables, models.Variable{Name: variableResp.Variable.Name, Value: variableResp.Variable.Value})
	return
}

type variable struct {
	Variable models.Variable `json:"variable"`
}

func (c *Endpoints) GetMyVariable(variableName string, target ...string) (variables models.Variables, err error) {
	respB, err := c.client.MakeReq("get", url.Values{"variableName": {variableName}}, "GET", target...)
	if err != nil {
		return
	}
	variableResp := variable{}
	err = json.Unmarshal(respB, &variableResp)
	if err != nil {
		return
	}
	variables.Variables = append(variables.Variables, models.Variable{Name: variableResp.Variable.Name, Value: variableResp.Variable.Value})
	return
}
func (c *Endpoints) ListAllMyVariables(target ...string) (variables models.Variables, err error) {
	respB, err := c.client.MakeReq("list", url.Values{}, "GET", target...)
	if err != nil {
		return
	}
	err = json.Unmarshal(respB, &variables)
	if err != nil {
		return
	}
	return
}
func (c *Endpoints) DeleteAllMyVariables(target ...string) (bool, error) {
	_, err := c.client.MakeReq("delete", url.Values{}, "GET", target...)
	if err != nil {
		return false, err
	}
	return true, nil
}
