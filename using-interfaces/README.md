# Using Interfaces

An interface is a set of method signatures. Just like in other languages that have them, an interface in Go can't do anything by itself. At some point there needs to be variable of a type that implements the interface, which gets passed around in the program and has its methods called.

One cool feature of Go interfaces is that they are "implicitly" implemented. Nowhere in a Go program will you see anything like `type SomeFooThing struct implements DoesFoo` except perhaps in auto-generated comments. Instead, as long as `SomeFooThing` has the methods in `DoesFoo`'s definition, it implements `DoesFoo`, and you can use it anywhere that a `DoesFoo` is called for, whether it be a local variable in a function or as an argument to a function.

This leads to two important guidelines:

- Prefer interfaces to structs in API definitions.
- Provide functions that return interfaces.

## Prefer interfaces to structs in API definitions

TODO

## Provide functions that return interfaces

TODO
