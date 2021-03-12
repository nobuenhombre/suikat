CONVERTER - Строка в другие типы
================================

GO конечно умеет все, но зачем же я написал этот модуль?
Например strconv.ParseFloat возвращает float64, а указать в параметрах можно и 32,
но вернется все равно 64 и надо будет писать обертку float32(v) для преобразования типа.
И c int32 int64 тоже самое.

Надо помнить параметры этих функций, а ведь всегда хочется попроще.
И я подумал путь будут имена функций StringTo<TypeName>.
И пусть они, эти функции возвращают имеено указанный тип безо всяких доп конверсий.

А ошибки - вон надо написать хотябы 5 конверсий рядом и идентифицировать ошибку
и я подумал пусть ошибка от strconv.Parse... будет в общей обертке ParserError

Вот что из этого вышло

Ошибка ParserError
------------------

.. code-block:: go
  :linenos:

    type ParserError struct {
      ParserType string
      Value      string
      Parent     error
    }

Функции модуля
--------------

.. code-block:: go
  :linenos:

    func StringToInt(s string) (int, error)
    func StringToInt8(s string) (int8, error)
    func StringToInt16(s string) (int16, error)
    func StringToInt32(s string) (int32, error)
    func StringToInt64(s string) (int64, error)
    func StringToBool(s string) (bool, error)
    func StringToFloat32(s string) (float32, error)
    func StringToFloat64(s string) (float64, error)
    func StringToTime(s, format string) (time.Time, error)
