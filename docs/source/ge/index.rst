GE - Global Errors
==================

Ошибки - у меня было несколько модификаций этого модуля.
Хотелось по ошибке быстро находить место в программе.
Хотелось чтобы ошибка содержала достаточно информации о том что произошло и с какими данными.

И вот оно получилось наконец.

Путь к ошибке
-------------

Чтобы понимать, где произошла ошибка - я сделал вот такую структуру

.. code-block:: go
  :linenos:

    type Way struct {
      Package string // Имя Пакета
      Caller  string // Функция внутри которой произошла ошибка
      File    string // Имя файла - ибо в пакете может быть несколько файлов
      Line    int    // Номер строки на которой произошла ошибка
    }

Теперь нужно заполнить эту структуру.

для этого я написал внутреннюю функцию модуля - снаружи она недоступна

.. code-block:: go
  :linenos:

    func getWay() *Way

Ошибка
------

Итак наша ошибка определена как IdentityError

Хотелось

#. Прокидывать родительскую ошибку - для этого завел поле ``Parent error``
#. Проверять ошибку методом errors.Is() - для этого завел метод ``func (e *IdentityError) Unwrap(err error) error``
#. Иметь ``Путь ошибки`` - Завел поле Way
#. Знать значения переменных которые могли привести к ошибке - завел тип ``type Params map[string]interface{}``

И в итоге придумал две функции

.. code-block:: go
  :linenos:

    // Возвращает новую ошибку IdentityError содержащую текст message и Путь Way
    // При желании можно добавить список значений переменных
    func New(message string, params ...Params) error

    // Возвращает новую ошибку IdentityError содержащую ссылку на родительскую ошибку и Путь Way
    // При желании можно добавить список значений переменных
    func Pin(parent error, params ...Params) error

Примеры использования
---------------------

.. code-block:: go
  :linenos:

    err := SomeCode(a, b)
    if err != nil {
        return ge.New("some code was error")
    }


.. code-block:: go
  :linenos:

    err := SomeCode(a, b)
    if err != nil {
        return ge.New("some code was error", ge.Params{"a": a, "b": b})
    }

.. code-block:: go
  :linenos:

    err := SomeCode(a, b)
    if err != nil {
        return ge.Pin(err)
    }

.. code-block:: go
  :linenos:

    err := SomeCode(a, b)
    if err != nil {
        return ge.Pin(&YourCustomError{})
    }

.. code-block:: go
  :linenos:

    err := SomeCode(a, b)
    if err != nil {
        return ge.Pin(err, ge.Params{"a": a, "b": b})
    }

.. code-block:: go
  :linenos:

    err := SomeCode(a, b)
    if err != nil {
        return ge.Pin(&YourCustomError{}, ge.Params{"a": a, "b": b})
    }
