# Using Interfaces

An interface is a set of method signatures. As in other languages that have them, an interface in Go can't do anything by itself. At some point, the program must create an object of a type that implements the interface that can be passed around and have its methods called.

One cool feature of Go interfaces is that they are "implicitly" implemented. Nowhere in a Go program will you see anything like `SomeFooThing implements DoesFoo` except perhaps in auto-generated comments. Instead, as long as `SomeFooThing` has all of the methods defined on the `DoesFoo` interface, `SomeFooThing` implements `DoesFoo`, and you can use it anywhere that calls for a `DoesFoo`.

Since there is no explicit declaration that a struct type implements a specific interface, people coming from Python should be pretty familiar with this pattern. It's how Go supports "duck typing".

This leads to two important guidelines:

- Prefer interfaces to structs in API definitions.
- Provide functions that return interfaces.

## Prefer interfaces to structs in API definitions

TODO

## Provide functions that return interfaces

TODO
