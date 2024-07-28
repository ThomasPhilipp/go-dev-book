+++
archetype = "chapter"
title = "Tipps"
weight = 8
+++

## Export variable

To export a variable and make it *public*, simply write the variable- and function-name with a capital letter. The public elements define the interface and the API of a package.

## Blank identifier

To ignore a value within e.g. a loop, we can use the *_* (blank identifier).

## String length

A string is representated internally as bytes. When using `len()` functions, it returns the number of bytes. The characters A-Z requires 1 byte, special characters can use more. To geht the number of characters we should use `RuneCountInString()` of the `utf8` package.