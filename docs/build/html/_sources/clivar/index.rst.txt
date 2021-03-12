CLIVAR - Сахаренные флаги
=========================

Есть такой пакет в GO - flag. Он позволяет прочитать параметры приложения заданные через командную строку

Например

.. code-block:: bash
  :linenos:

    myapp -file=readme.txt -action=transform -out=pdf

    // тут флаги
    -file=readme.txt
    -action=transform
    -out=pdf

Но все время писать разбор входящих параметров через флаги в каждом новом приложении утомительно.
Да и однообразно, не хочется писать все время одинаковый код отличающийся только буквами.

И вот они

Сахаренные флаги
----------------

Суть в том, чтобы не писать все время однотипный код, флаги описываются в виде тегов структуры
Т.е. вы задаете структуру поля которой и будут флагами вашего приложения,
а как парсить эти флаги описано в тегах ``cli:``

тег состоит из следующих частей

.. code-block:: go
  :linenos:

    cli:"<Имя флага>[<Текст подсказки если флаг указан с опечаткой>]:<Тип>=<Значение по умолчанию>"

Вот как выглядит пример приложения с сахаренными флагами

.. code-block:: go
  :linenos:

    package main

    import (
      "fmt"
      "log"

      "github.com/nobuenhombre/suikat/pkg/clivar"
    )

    type CliConfig struct {
      Path           string  `cli:"PATH[Путь куда то]:string=/some/default/path"`
      Port           int     `cli:"PORT[Порт портальный]:int=8080"`
      Coefficient    float64 `cli:"COEFFICIENT[Коэффициент тунгусский]:float64=75.31"`
      MakeSomeAction bool    `cli:"MSA[Можно мне?]:bool=false"`
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

компонент использует рефлексию