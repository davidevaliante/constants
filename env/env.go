package env

import "fmt"

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

func EnvFromString(s string) (*Env, error) {
	env, exists := envMap[s]
	if !exists {
		return nil, fmt.Errorf("unsupported Env string format: %s", s)
	}
	return &env, nil
}
