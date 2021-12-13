ADAPT
=====

Package adapt provides functions
to check vars interface{} type
and convert interface{} to basic types: bool, int, string.


Check
-----
compares reflect.Value with a string representation of the expected type.
If the type does not match the expected one, the ge.MismatchError{} error is returned

.. code-block:: go
  :linenos:

    func Check(val reflect.Value, expectType string) error


Bool
----
convert v interface{} into bool.
If v can't be converted - return ge.MismatchError{}

.. code-block:: go
  :linenos:

    func Bool(v interface{}) (bool, error)


Int
---
convert v interface{} into int.
If v can't be converted - return ge.MismatchError{}

.. code-block:: go
  :linenos:

    func Int(v interface{}) (int, error)


String
------
convert v interface{} into string.
If v can't be converted - return ge.MismatchError{}

.. code-block:: go
  :linenos:

    func String(v interface{}) (string, error)


IsNil
-----
check interface{} is nil.
and return bool value

.. code-block:: go
  :linenos:

    func IsNil(i interface{}) bool
