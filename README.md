# Scantastic!

A scanner controller with a very simple REST interface.

## Setup

This package depends on a few libraries that do a lot of the heavy lifting to keep this codebase small:
- SANE: a library for granting controller access to scanners via C
- github.com/tjgq/sane: a library that provides Golang bindings to SANE
- github.com/gorilla/mux: a library that provides a simple interface for creating REST interfaces, skipping over a lot of boilerplate

These must be installed before this package can be built:
```
$ sudo apt-get install libsane-dev # NOTE THAT IT MUST BE THE DEV LIBRARY
$ go get github.com/tjgq/sane
$ go get github.com/gorilla/mux
```

## To Build

Simply build as you normally would:
```
$ go build
```

## To Run

Once build, run the following:
```
$ ./scantastic
```

The REST interface will be exposed on port 8000

## Known issues:

There's a problem with cross-compiling from a Mac to a Pi; for now, I have to pull the code onto a Pi and compile it there.  It's not great.