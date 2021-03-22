ENVAR - Сахаренные переменные окружения
=======================================

Умеет GO читать переменные окружения. Вот так ``os.LookupEnv(...``

И можно использовать хоть миллион этих переменных, однако код пухнет.

поскольку сахаренные флаги уже у меня были, я решил посахарить и переменные окружения.

Сахаренные переменные окружения
-------------------------------

Суть в том, чтобы не писать все время однотипный код, переменные окружения описываются в виде тегов структуры
Т.е. вы задаете структуру поля которой и будут переменными окружения для вашего приложения,
а как парсить эти флаги описано в тегах ``env:``

тег состоит из следующих частей

.. code-block:: go
  :linenos:

    env:"<Имя переменной окружения>:<Тип>=<Значение по умолчанию>"

Вот как выглядит пример приложения с сахаренными переменными окружения

.. code-block:: go
  :linenos:

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


компонент использует рефлексию