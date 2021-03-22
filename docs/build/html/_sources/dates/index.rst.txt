DATES - Даты
============

Шаблоны кода которые я часто использую для работы с датами

Доступные константы
-------------------

Для форматирования
~~~~~~~~~~~~~~~~~~

.. code-block:: go
  :linenos:

    const (
	  DateFormatDashYYYYMMDD           = "2006-01-02"
	  DateFormatPointDDMMYYYY          = "02.01.2006"
	  DateTimeFormatDashYYYYMMDDHHmmss = "2006-01-02 15:04:05"
	  DateTimeFormat1C                 = "2006-01-02T15:04:05"
    )

Собственно, писать шаблон даты через цифры - особенно когда привык к буквам, медленно. Приходится вспоминать,
что там за цифры, искать похожий код или статью из документации языка. поэтому завел себе вот такие константы.
IDE показывает имя константы и сразу понятно - что там будет.
Плюс для взаимодействия с 1С нужен был определенный формат даты.

Для расчетов и циклов
~~~~~~~~~~~~~~~~~~~~~

.. code-block:: go
  :linenos:

    const (
	  WeekDays  = 7
	  MonthDays = 31
	  YearDays  = 365
    )

    const (
	  SecondsInMinute = 60
	  MinutesInHour   = 60
	  HourInDay       = 24
	  MonthInYear     = 12
    )

Методы модуля
-------------

.. code-block:: go
  :linenos:

    // Вычислить разность дат со временем
    func Diff(a, b time.Time) *DateTimeDiff

    // Начало дня
    func BeginOfDay(t time.Time) time.Time

    // Начало предыдущего дня
    func BeginOfPrevDay(t time.Time) time.Time

    // Начало завтрашнего дня
    func BeginOfNextDay(t time.Time) time.Time

    // Начало дня неделю назад
    func BeginOfPrevWeek(t time.Time) time.Time

    // Начало дня неделю вперед
    func BeginOfNextWeek(t time.Time) time.Time

    // Некоторое время назад
    func BeforePeriod(t time.Time, period int64, measure time.Duration) time.Time

    // Некоторое время вперед
    func AfterPeriod(t time.Time, period int64, measure time.Duration) time.Time
