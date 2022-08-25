# tomedo

A Go-package and binary to access the database of a tomedo server.

## Installation

### To install the `tomedo` binary

The easiest way to install the `tomedo` command is by running...

    $ go install github.com/bjoernalbers/tomedo/cmd/...@latest

Alternatively you can clone this repository, `cd` into it and execute:

    $ go install ./...

Both methods will install the binary into `$GOBIN` which defaults to
`$GOPATH/bin` or `$HOME/go/bin`.

### To install the Go library / package

In case you just want to use the "tomedo" Go-package, install it like this
within your module:

    $ go get github.com/bjoernalbers/

## Usage

Running `tomedo` displays a list of users.
