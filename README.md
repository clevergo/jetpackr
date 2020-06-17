# Jet Template Engine Packr Loader
[![Build Status](https://travis-ci.org/clevergo/jetpackr.svg?branch=master)](https://travis-ci.org/clevergo/jetpackr)
[![Coverage Status](https://coveralls.io/repos/github/clevergo/jetpackr/badge.svg?branch=master)](https://coveralls.io/github/clevergo/jetpackr?branch=master)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/clevergo.tech/jetpackr?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/jetpackr)](https://goreportcard.com/report/github.com/clevergo/jetpackr)
[![Sourcegraph](https://sourcegraph.com/github.com/clevergo/jetpackr/-/badge.svg)](https://sourcegraph.com/github.com/clevergo/jetpackr?badge)
[![Release](https://img.shields.io/github/release/clevergo/jetpackr.svg?style=flat-square)](https://github.com/clevergo/jetpackr/releases)

## Install

```shell
$ go get github.com/clevergo/jetpackr
```

## Usage

See [example](example).

```go
package main

import (
	"github.com/CloudyKit/jet/v3"
	"github.com/clevergo/jetpackr"
	"github.com/gobuffalo/packr/v2"
)

func main() {
	box := packr.New("views", "./views")
	view := jet.NewHTMLSetLoader(jetpackr.New(box))
	// ...
}
```
