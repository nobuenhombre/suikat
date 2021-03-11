CHUNKS - Кусочки
================

Модуль предоставляет одну единственную функцию, которая делит входящий слайс на кусочки определенной длины

.. code-block:: go
  :linenos:

    func Split(in []int64, limit int) [][]int64

На входе слайс из ``int64`` и число элементов в кусочках которые будут на выходе

Например

.. code-block:: go
  :linenos:

    import "github.com/nobuenhombre/suikat/pkg/chunks"

    // На входе
    in := []int64{1, 3, 5, 7, 9, 0, 2, 4, 6, 8}
    limit := 4

    out := chunks.Split(in, limit)

    // На выходе
    // out: [][]int64{
    //		{1, 3, 5, 7},
    //		{9, 0, 2, 4},
    //		{6, 8},
    //	},
