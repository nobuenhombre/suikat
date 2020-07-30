package main

import (
	"fmt"
	"log"

	"github.com/nobuenhombre/suikat/pkg/envar"
)

type AppConfig struct {
	Path           string  `env:"PATH:string=/some/default/path"`
	Port           int     `env:"PORT:int=8080"`
	Coefficient    float64 `env:"COEFFICIENT:float64=75.31"`
	MakeSomeAction bool    `env:"MSA:bool=false"`
}

func (cfg *AppConfig) Load() error {
	return envar.Load(cfg)
}

func main() {
	cfg := &AppConfig{}

	err := cfg.Load()
	if err != nil {
		log.Fatalf("ENV config error [%v]", err)
	}

	fmt.Printf("%#v", cfg)
}
