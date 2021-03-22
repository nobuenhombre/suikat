FICO - FIles COntent
====================

Методы модуля
-------------

.. code-block:: go
  :linenos:

    // Текстовый (бинарный) файл - по сути это его имя - полный путь
    type TxtFile string

    // Читает содержимое в виде слайса байтов
    func (f *TxtFile) ReadBytes() ([]byte, error)

    // Читает содержимое в виде строки
    func (f *TxtFile) Read() (string, error)

    // Записывает строку в файл
    func (f *TxtFile) Write(content string) error

    // Записывает строку в упакованный файл .gz
    func (f *TxtFile) WriteGZ(content string) error

    // Читает исходный файл и упаковывает его в .gz рядом с исходным файлом
    func (f *TxtFile) GZ() error

    // Читает исходный файл в виде base64 строки
    func (f *TxtFile) ReadAsB64String() (string, error)

    // Читает исходный файл в виде base64 строки записывает эту строку в файл с расширением .b64 рядом с исходным файлом
    func (f *TxtFile) B64() error

    // Берет обычную строку - и превращает ее HEX представление с разделителями glue
    func StrBytes(in, glue string, isUpper bool) string

    // Читает исходный файл и возвращает строку его HEX представления
    func (f *TxtFile) ReadAsHexString() (string, error)

    // Читает исходный файл - превращает ее HEX представление и записывает в файл .hex рядом с исходным файлом
    func (f *TxtFile) Hex() error

    // Читает исходный файл, получает mime-type, и base64 представление и возвращает строку вида "data:<MIME-TYPE>;base64,<BASE64-DATA>"
    func (f *TxtFile) DataURI() (string, error)

.. code-block:: go
  :linenos:

    // Пачка файлов [имя файла]контент
    type TxtFilesPack map[string]string

    // Читает контент всех файлов в пачке
    func (p *TxtFilesPack) Read() error

    // Пишет контент в файлы пачки
    func (p *TxtFilesPack) Write() error

    // Пишет контент пачки в упакованные файлы
    func (p *TxtFilesPack) WriteGZ() error

    // Удаляет все файлы пачки
    func (p *TxtFilesPack) Remove() error
