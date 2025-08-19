package config

import "flag"

type Argser interface {
	ParseFlags()
}

type ArgsCLI struct {
	Port string // StartPort is the port for start application
}

func NewArgsCLI() *ArgsCLI {
	ArgsCLI := new(ArgsCLI)
	ArgsCLI.ParseFlags()
	return ArgsCLI
}

func (a *ArgsCLI) ParseFlags() {
	flag.StringVar(&a.Port, "port", "localhost:8080", "Start adress for application")
	flag.Parse()
}
