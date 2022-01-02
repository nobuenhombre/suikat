package main

import (
	"fmt"
	"log"

	"github.com/nobuenhombre/suikat/pkg/clivar"
)

type CliConfig struct {
	Path           string   `cli:"PATH[Путь куда то]:string=/some/default/path"`
	Port           int      `cli:"PORT[Порт портальный]:int=8080"`
	Coefficient    float64  `cli:"COEFFICIENT[Коэффициент тунгусский]:float64=75.31"`
	MakeSomeAction bool     `cli:"MSA[Можно мне?]:bool=false"`
	Languages      []string `cli:"LANGUAGES[Languages]:SliceStrings=ru,en,de,fr"`
}

func (cfg *CliConfig) Load() error {
	return clivar.Load(cfg)
}

func main() {
	cfg := &CliConfig{}

	err := cfg.Load()
	if err != nil {
		log.Fatalf("CLI config error [%v]", err)
	}

	fmt.Printf("%#v", cfg)
}
