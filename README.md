# Jet Template Engine Packr Loader
[![Build Status](https://img.shields.io/travis/clevergo/jetpackr?style=for-the-badge)](https://travis-ci.org/clevergo/jetpackr)
[![Coverage Status](https://img.shields.io/coveralls/github/clevergo/jetpackr?style=for-the-badge)](https://coveralls.io/github/clevergo/jetpackr?branch=master)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/clevergo.tech/jetpackr?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/jetpackr?style=for-the-badge)](https://goreportcard.com/report/github.com/clevergo/jetpackr)
[![Release](https://img.shields.io/github/release/clevergo/jetpackr.svg?style=for-the-badge)](https://github.com/clevergo/jetpackr/releases)
[![Downloads](https://img.shields.io/endpoint?url=https://pkg.clevergo.tech/api/badges/downloads/month/clevergo.tech/jetpackr&style=for-the-badge)](https://pkg.clevergo.tech/)

## Install

```shell
$ go get -u clevergo.tech/jetpackr
```

## Usage

See [example](https://github.com/clevergo/examples/tree/master/jetrenderer).

```go
package main

import (
	"github.com/CloudyKit/jet/v5"
	"github.com/clevergo/jetpackr"
	"github.com/gobuffalo/packr/v2"
)

func main() {
	box := packr.New("views", "./views")
	view := jet.NewHTMLSetLoader(jetpackr.New(box))
	// ...
}
```
