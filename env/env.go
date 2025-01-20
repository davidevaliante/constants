package env

import (
	"fmt"
)

type Env int

const (
	Local Env = iota
	Development
	Staging
	Production
)

var envMap = map[string]Env{
	"local":       Local,
	"development": Development,
	"staging":     Staging,
	"production":  Production,
}

func EnvFromString(s string) *Env {
	env, exists := envMap[s]
	if !exists {
		fmt.Println(fmt.Sprintf("unable to map value %s to env, defaulting to env.Local", s))
		env, _ = envMap[s]
		return &env
	}
	return &env
}
