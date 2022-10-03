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

// SetMyVariableToEnvironment sets a variable in the environment
func SetMyVariableToEnvironment(variableName string, variableValue string) (string, error) {
	if variableName == "" {
		return "", errors.New("variable name is empty")
	}
	//verifies if the variable is set in the environment
	sysVariableValue, ok := os.LookupEnv(composeVariableName(variableName))
	fmt.Printf("Variable %s is  present with value ", sysVariableValue)
	//if not set, generate a random string and setting it in the environment
	if !ok {
		fmt.Printf("Variable %s is not present, set variable", variableName)
		if len(variableValue) == 0 {
			variableValue = generateRandomString(10)
			fmt.Printf("Variable %s has no value, set variable random value", variableName)
		}
		err := os.Setenv(composeVariableName(variableName), variableValue)
		if err != nil {
			return "", err
		}
		return variableValue, nil
	}

	if sysVariableValue != variableValue {
		err := os.Setenv(composeVariableName(variableName), variableValue)
		if err != nil {
			return "", err
		}
	}
	return variableValue, nil
}

// GetMyVariableFromEnvironment gets a variable from the environment
func GetMyVariableFromEnvironment(variableName string) (string, error) {

	if variableName == "" {
		return "", errors.New("variable name is empty")
	}
	//verifies if the variable is set in the environment
	sysVariableValue, ok := os.LookupEnv(composeVariableName(variableName))
	if !ok {
		return "", errors.New("variable not found")
	}
	return sysVariableValue, nil
}

// GetAllMyVariablesFromEnvironment gets all variables from the environment
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

// DeleteAllMyVariablesFromEnvironment deletes all variables from the environment
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
