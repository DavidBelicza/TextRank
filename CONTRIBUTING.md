# Contributing

Contributions are welcome. This page describes what is expected before a pull
request is merged.

## Before you start

* For a bug, open an issue with code that reproduces it.
* For a new feature, open an issue first and wait for an answer. Features are
  discussed before any code is written, so nobody works on something that will
  not be merged.
* Fixing a logical error in the ranking needs no discussion, just send the pull
  request.

## Pull requests

* Create a branch for the issue and keep the pull request small. Several small
  pull requests are reviewed faster than one big one.
* Do not add third party dependencies. The library depends only on the standard
  library, and the test dependency on testify is the single exception.
* Do not change exported functions, types or interfaces without describing it in
  the pull request. The module path ends with /v2, so a breaking change needs a
  new major version.

## Tests

Run the tests and the vet before sending the pull request:

```
go test ./...
go vet ./...
```

New and changed code needs tests. Coverage of the project is at 100 percent and
the build fails when it drops, so keep it there.

Code has to be formatted with gofmt:

```
gofmt -l .
```

## Code style

Follow the style of the surrounding code. Exported functions and types have a
comment, unexported ones usually do not.
