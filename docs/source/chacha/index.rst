CHACHA - Дерево Валидации
=========================

Дерево валидации выполняет как ни странно валидацию чего нибудь.
Например нам нужна многоуровневая проверка на определенные условия.

Например можно составить такой план валидации какого то json:

1) Проверить, что входящие данные (строка) это вообще json
2) Проверить, что в данном json есть определенный набор полей
3) Проверить, что определенное поле содержит допустимое значение

здесь каждая последующая проверка зависит от предыдущей
ведь не имеет смысла искать набор полей в данных которые не являются json

также бывают параллельные проверки от одного корня
например

4) Проверить, что число букв в поле в нужном диапазоне
5) Проверить, что есть такие буквы и нет других букв

и т.п.

т.е. получилось вот такое дерево валидаторов

.. code-block:: bash
  :linenos:

    1
    └─2
      └─3
        ├─4
        └─5

Валидатор
---------

Для того чтобы создать валидатор нужно воспользоваться типом ``Validator`` из данного пакета

.. code-block:: go
  :linenos:

    type Validator func(params ...interface{}) (interface{}, error)


как мы видим функция валидатор, это Вариативная функция, т.е. она на вход принимает любое количество аргументов

Например

.. code-block:: go
  :linenos:

    import "github.com/nobuenhombre/suikat/pkg/chacha"

    ...

    type YourJSON struct {
      Hello string
    }

    type JsonError struct {
	    Parent error
    }

    func (e *JsonError) Error() string {
	    return fmt.Sprintf("invalid json [%v]", e.Parent)
    }

    // Вот он наш валидатор
    func IsValidJSON(params ...interface{}) (interface{}, error) {
      jsonStr := params[0]

      testJson := &YourJSON{}

      err := json.Unmarshal(jsonStr, testJson)
      if err != nil {
		return nil, &JsonError{
		  Parent: err,
	    }
      }

      return testJson, nil
    }

Поучилась такая функция валидатор, которая пытается разобрать json строку и превратить ее в нашу структуру YourJSONStruct
если это не удается - функция возвращает ошибку ``JsonError``
если это удается - она возвращает получившуюся структуру

вообще в валидаторе можно вернуть следующие варианты ответов

.. code-block:: go
  :linenos:

    import "github.com/nobuenhombre/suikat/pkg/chacha"

    ...

    func IsValidSomething(params ...interface{}) (interface{}, error) {

      ...

      // 1. Какую то ошибку - это означает что валидация не прошла
      //    т.е. исследуемые данные невалидны
      //    Это значит что дочерние валидаторы не будут выполнятся
      return nil, someError

      // 2. Нет ошибки и нет никаких данных
      //    если есть дочерние валидаторы они будут вызваны
      return nil, nil

      // 3. Нет ошибки, есть данные
      //    если есть дочерние валидаторы они будут вызваны
      //    Кроме того эти данные будут переданы в дочерние валидаторы
      //    что-бы дочерние валидаторы могли их проверить
      return someData, nil

      // 4. Нет ошибки, но и дочерние валидаторы я не хочу запускать
      //    дочерние валидаторы не будут выполнятся
      return new(chacha.DontCheckChildrens), nil

    }

Дерево Валидации
----------------

Вот такой код можно написать чтобы составить дерево валидации

.. code-block:: go
  :linenos:

    import "github.com/nobuenhombre/suikat/pkg/chacha"

    ...

    // Объявляем валидаторы
    func IsValidRoot(params ...interface{}) (interface{}, error) { ... }
    func IsValidSortPresent(params ...interface{}) (interface{}, error) { ... }
    func IsValidSortLength(params ...interface{}) (interface{}, error) { ... }
    func IsValidNameLength(params ...interface{}) (interface{}, error) { ... }
    func IsValidDepth(params ...interface{}) (interface{}, error) { ... }
    func IsValidCount(params ...interface{}) (interface{}, error) { ... }
    func IsValidNames(params ...interface{}) (interface{}, error) { ... }

    ...

    // Объявляем Дерево Валидаторов
    type MyValidator struct {
      Rules *chacha.Tree
    }

    func (v *MyValidator) CreateRules() {
      v.Rules = &chacha.Tree{
        Validator: IsValidRoot,
        Childrens: []chacha.Tree{
          {
            Validator: IsValidSortPresent,
            Childrens: []chacha.Tree{
              {Validator: IsValidSortLength},
              {Validator: IsValidNameLength},
            },
          },
          {Validator: IsValidDepth},
          {Validator: IsValidCount},
          {Validator: IsValidNames},
        },
      }
    }

    ...

    func ValidateData(data) {
      for index := range data {

        validator := &MyValidator{}
        validator.CreateRules()
        validator.Rules.Validate(data[index])

        if validator.Rules.Valid {
          // Что то делать если валидно
          // Например считать счетчик
        } else {
          // Что то делать если невалидно
          // Например считать счетчик и выводить на экран ошибку валидации
        }
      }
    }