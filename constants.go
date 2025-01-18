package constants

type Env int

const (
	Local Env = iota
	Development
	Staging
	Production
)
