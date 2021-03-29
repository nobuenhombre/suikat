FITREE - Files Tree
===================

Компонент предназначен для работы с файловыми деревьями.

Для этого я завел две структуры

#. Узел или ``TreeNodeStruct``
#. Список узлов или ``TreeNodeListStruct``

.. code-block:: go
  :linenos:

    // Узел
    type TreeNodeStruct struct {
      Path         string         // Путь
      Name         string         // Имя
      Depth        int            // Глубина от корня
      Files        []os.FileInfo  // Список файлов в данном узле
      FilesCount   int            // Число файлов
      SubDirs      []os.FileInfo  // Список подкаталогов
      SubDirsCount int            // Число подкаталогов
    }

.. code-block:: go
  :linenos:

    // Список узлов
    type TreeNodeListStruct struct {
      List    []TreeNodeStruct    // Список узлов TreeNodeStruct
      Reverse map[string]int      // Обратный индекс для быстрых путешествий по дереву (подробнее см. ниже)
    }

Узел
----

Узел имеет только один метод - Fill.

.. code-block:: go
  :linenos:

    func (node *TreeNodeStruct) Fill(path string, depth int) error

Здесь два входящих параметра

#. path - Путь
#. depth - Стартовая Глубина (точка отсчета)

например путь ``"/home/hello"``
и глубина - обозначим ее 0
после исполнения будет заполнена указанная структура.
соответственно можно узнать список файлов в данном каталоге, их количество, список вложенных каталогов и их количество.

Список узлов
------------

Имеет уже три метода

.. code-block:: go
  :linenos:

    // Добавляет в список новый Узел
    func (list *TreeNodeListStruct) Add(node TreeNodeStruct)

    // Получает Узел для указанного пути со стартовой глубиной
    // и добавляет Узел в список
    // Далее Получает узлы вложенных каталогов и так рекурсивно спускается до самого глубокого каталога.
    func (list *TreeNodeListStruct) Scan(path string, depth int, ignoreErr bool) error

    // Возвращает Узел по его индексу
    func (list *TreeNodeListStruct) GetNode(index int) (*TreeNodeStruct, error)
