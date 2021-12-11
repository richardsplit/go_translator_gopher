package env

import (
	"os"
	"strconv"

	"github.com/pkg/errors"
)

func OptionalString(envName string, defaultValue string) string {
	value, ok := os.LookupEnv(envName)
	if !ok {
		return defaultValue
	}
	return value
}

func OptionalInt(envName string, defaultValue int) (int, error) {
	value, ok := os.LookupEnv(envName)
	if !ok {
		return defaultValue, nil
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.Errorf("environment variable %q present but not of type int", envName)
	}

	return intValue, nil
}
