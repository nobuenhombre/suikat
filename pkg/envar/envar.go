package envar

import (
	"os"
	"strconv"
)

type EnvVar struct {
	Key          string
	DefaultValue interface{}
}

func (ev *EnvVar) GetString() string {
	if value, exists := os.LookupEnv(ev.Key); exists {
		return value
	}

	return ev.DefaultValue.(string)
}

func (ev *EnvVar) GetInt() int {
	if valueStr, exists := os.LookupEnv(ev.Key); exists {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}

	return ev.DefaultValue.(int)
}

func (ev *EnvVar) GetBool() bool {
	if valueStr, exists := os.LookupEnv(ev.Key); exists {
		if value, err := strconv.ParseBool(valueStr); err == nil {
			return value
		}
	}

	return ev.DefaultValue.(bool)
}
