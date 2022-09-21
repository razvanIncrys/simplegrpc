package variable

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Variable struct {
	Name  string
	Value string
}

func GetMyVariableFromEnvironment(variableName string) (string, error) {

	if variableName == "" {
		return "", errors.New("variable name is empty")
	}
	//verifies if the variable is set in the environment
	sysVariable, ok := os.LookupEnv(composeVariableName(variableName))
	//if not set, generate a random string and setting it in the environment
	if !ok {
		fmt.Printf("Variable %s is not present, set default value", variableName)
		sysVariable = generateRandomString(10)
		err := os.Setenv(composeVariableName(variableName), sysVariable)
		if err != nil {
			return "", err
		}
	}
	return sysVariable, nil
}

func GetAllMyVariablesFromEnvironment() []Variable {
	var variables []Variable
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if strings.Contains(pair[0], "SIMPLEGRPC_") {
			varName := strings.Replace(pair[0], "SIMPLEGRPC_", "", 1)
			variables = append(variables, Variable{Name: varName, Value: pair[1]})
		}
	}
	return variables
}

func DeleteAllMyVariablesFromEnvironment() (bool, error) {
	var variables []Variable
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if strings.Contains(pair[0], "SIMPLEGRPC_") {
			err := os.Unsetenv(pair[0])
			if err != nil {
				return false, err
			}
			variables = append(variables, Variable{Name: pair[0], Value: pair[1]})
		}
	}
	return true, nil
}

func composeVariableName(variableName string) string {
	return fmt.Sprintf("SIMPLEGRPC_%s", variableName)
}

func generateRandomString(n int) string {
	//create a slice of runes from alphabet
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	//create a slice of runes with length n
	s := make([]rune, n)
	//populate the slice with random runes
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
