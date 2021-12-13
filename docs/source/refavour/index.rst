REFAVOUR
========
Package refavour provides interface to work with tags of struct and several functions for working with reflect

TagProcessor
------------
interface to provide parsing tags of struct


GetReflectValue
---------------
return reflect.Value of interface{}


CheckKind
---------
check is that interface{} equal expected Kind
if equal then return nil
else return KindNotMatchedError


CheckStructure
--------------
check is that interface{} struct
if struct then return nil
else return KindNotMatchedError


CheckMap
--------
check is that interface{} map
if map then return nil
else return KindNotMatchedError


CheckSlice
----------
check is that interface{} slice
if slice then return nil
else return KindNotMatchedError


CheckCanBeChanged
-----------------
Check whether the data receiver can be changed


FieldsInfo
----------
Types of structure fields


GetStructureFieldsTypes
-----------------------
reads the structure from the interface and generates a list of field names and their types