FUTI - File Utils
=================

Функции для частых операций над файлами

.. code-block:: go
  :linenos:

    // Существует ли файл (именно файл, не каталог)
    func FileExists(filename string) bool

    // Существует ли каталог
    func DirExists(dirname string) bool

    // Удалить файл или каталог
    func Delete(fileName string) error

    // Создать временный файл в каталоге временных файлов если dir=''
    // или в конкретном каталоге dir,
    // имя формируется по шаблону pattern - в нем символ * заменяется на случайный набор символов
    // - например "in-*.pdf" создаст файл "in-hdy8g8yg3ufc.pdf"
    // Записать в него данные data и вернуть имя файла
    func CreateTempFile(dir, pattern string, data *[]byte) (string, error)

    // скопировать файл из inFile в outFile
    func Copy(inFile, outFile string) error

    // переместить файл из inFile в outFile
    func Move(inFile, outFile string) error