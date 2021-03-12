COLOROG - Цветной Лог
=====================

Захотелось мне как то раз сделать цветной лог.
Xто-бы время в логе одним цветом, а строчки тоже разноцветные, но не какие угодно, а следующих типов:

#. Успех - Зеленого цвета
#. Информация - Голубого цвета
#. Сообщение - Белого цвета
#. Ошибка незначительная - Желтого цвета
#. Фатальная ошибка - Красного цвета
#. и Паника - Красного цвета на белом фоне

и чтобы время можно было не показывать если хочется
и что-бы формат даты времени можно было настроить
взял и написал

Конструктор
-----------

.. code-block:: go
  :linenos:

    func NewColoredLog(showTime bool, timeFormat string) *ColoredLog

#. showTime - показывать или не показывать время
#. timeFormat - шаблон даты-времени в логе - например "2006-01-02 15:04:05"

пример

.. code-block:: go
  :linenos:

    import "github.com/nobuenhombre/suikat/pkg/colorog"

    ...

    // Без даты и времени
    colorLog := colorog.NewColoredLog(false, "")

    // Только дата
    colorLog := colorog.NewColoredLog(true, "2006-01-02")

    // И дата и время
    colorLog := colorog.NewColoredLog(true, "2006-01-02 15:04:05")


Методы
------

.. code-block:: go
  :linenos:

    // Сообщение об успехе - Зеленым цветом
    // - для вывода нескольких переменных без переноса строки
    func (cl *ColoredLog) Success(v ...interface{})
    // - для вывода нескольких переменных внутри отформатированного текста
    func (cl *ColoredLog) Successf(format string, v ...interface{})
    // - для вывода нескольких переменных с переносом строки
    func (cl *ColoredLog) Successln(v ...interface{})

    // Сообщение об ошибке которая не приводит к выходу из программы - Желтым цветом
    func (cl *ColoredLog) Error(v ...interface{})
    func (cl *ColoredLog) Errorf(format string, v ...interface{})
    func (cl *ColoredLog) Errorln(v ...interface{})

    // Информационное сообщение - Голубым цветом
    func (cl *ColoredLog) Info(v ...interface{})
    func (cl *ColoredLog) Infof(format string, v ...interface{})
    func (cl *ColoredLog) Infoln(v ...interface{})

    // Простое сообщение - Белым цветом
    func (cl *ColoredLog) Message(v ...interface{})
    func (cl *ColoredLog) Messagef(format string, v ...interface{})
    func (cl *ColoredLog) Messageln(v ...interface{})

    // Сообщение об ошибке которая приводит к выходу из программы - Красным цветом
    // кроме сообщения - происходит выход из программы с кодом 1
    func (cl *ColoredLog) Fatal(v ...interface{})
    func (cl *ColoredLog) Fatalf(format string, v ...interface{})
    func (cl *ColoredLog) Fatalln(v ...interface{})

    // Паника - красное на белом фоне
    // кроме сообщения - создает панику
    func (cl *ColoredLog) Panic(v ...interface{})
    func (cl *ColoredLog) Panicf(format string, v ...interface{})
    func (cl *ColoredLog) Panicln(v ...interface{})

.. _colorog-example:

Пример
------

.. code-block:: go
  :linenos:

    package main

    import (
        "flag"
        "github.com/nobuenhombre/suikat/pkg/colorog"
        "github.com/nobuenhombre/suikat/pkg/fico"
    )

    func main() {
      log := colorog.NewColoredLog(false, "")

      log.Infoln("B64")
      log.Messageln("this program read file.txt and create file.txt.b64 with base64 encoded content")

      fileName := flag.String(
        "file",
        "file.txt",
        "файл для перекодирования в b64")

      flag.Parse()

      txtFile := fico.TxtFile(*fileName)

      err := txtFile.B64()
      if err != nil {
        log.Fatalf("B64 error [%v]", err)
      }

      log.Success("Success")
    }
