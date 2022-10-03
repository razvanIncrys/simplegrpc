package models

type Variable struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Variables struct {
	Variables []Variable `json:"variables"`
}
