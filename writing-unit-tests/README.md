# Writing Unit Tests

## The Basics

Files ending in `_test.go` are test files. You place them in the same package directory as the files containing the production code that they cover.

Since these are go files, the first non-comment line should be a package declaration. A test file's package determines what code it can test. To test the package's public API, the test file's package name should be the same as that of the production package, but with `_test` appended. For example, to test the public API of package `foo`, the test file's package should be named `foo_test`. To test a package's private code, the test file's package should be the same as that containing the private code. For example, to test private code in package `foo`, the test file's package should also be `foo`.

When creating a test file to a package, the first thing to decide is whether it should test the package's public API or its private code. I recommend focusing on testing the public API at first. This helps you to think more carefully about API design. These tests can also be less brittle, giving you more freedom to iterate on implementation details without having to rewrite them. The implementation details live in private code, and you can always add more test files when you want closer coverage over specific functions.

## Assertions with github.com/stretchr/testify/assert

TODO

## Mocking with github.com/stretchr/testify/mock

TODO
