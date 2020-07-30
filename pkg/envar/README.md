# ENVAR

компонент предназначен для чтения env (настроек окружения)

## Пример использования

```go
package main

import (
	"fmt"
	"github.com/nobuenhombre/suikat/pkg/envar"
)

type AppConfig struct {
	Path           string  `env:"PATH:string=/some/default/path"`
	Port           int     `env:"PORT:int=8080"`
	Coefficient    float64 `env:"COEFFICIENT:float64=75.31"`
	MakeSomeAction bool    `env:"MSA:bool=false"`
}

func main() {
	cfg := &AppConfig{}
	envar.GetENV(cfg)

	fmt.Printf("%#v", cfg)
}

``` 