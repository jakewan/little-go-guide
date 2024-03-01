# Using Interfaces

An interface is a set of method signatures. As in other languages that have them, an interface in Go can't do anything by itself. At some point, the program must create an object of a type that implements the interface that can be passed around and have its methods called.

One feature of Go interfaces is that struct types implement them "implicitly". Nowhere in a Go program will you see anything like `SomeFooThing implements DoesFoo` except perhaps in auto-generated comments. Instead, as long as struct `SomeFooThing` has all of the methods defined on the `DoesFoo` interface, `SomeFooThing` is said to implement `DoesFoo`, and you can use it anywhere that calls for a `DoesFoo`.

Since there is no explicit declaration that a struct type implements a specific interface, people coming from Python should be pretty familiar with this pattern. It's how Go supports "duck typing".

Interfaces play a significant role in testability and public API design. See [Writing Unit Tests][] and [Project Organization - Public API Design][] for more on these topics.

## Examples

The `example` subdirectory contains a small program illustrating basic interface usage.

[project organization - public api design]: ../project-organization/README.md#public-api-design
[writing unit tests]: ../writing-unit-tests/README.md
