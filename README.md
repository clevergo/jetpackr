# Jet Template Engine Packr Loader
[![Build Status](https://travis-ci.org/clevergo/jet-packrloader.svg?branch=master)](https://travis-ci.org/clevergo/jet-packrloader)
[![Coverage Status](https://coveralls.io/repos/github/clevergo/jet-packrloader/badge.svg?branch=master)](https://coveralls.io/github/clevergo/jet-packrloader?branch=master)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/clevergo/jet-packrloader)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/jet-packrloader)](https://goreportcard.com/report/github.com/clevergo/jet-packrloader)
[![Release](https://img.shields.io/github/release/clevergo/jet-packrloader.svg?style=flat-square)](https://github.com/clevergo/jet-packrloader/releases)

## Install

```shell
$ go get github.com/clevergo/jet-packrloader
```

## Usage

See [example](example).

```go
package main

import (
	"github.com/CloudyKit/jet/v3"
	packrloader "github.com/clevergo/jet-packrloader"
	"github.com/gobuffalo/packr/v2"
)

func main() {
	box := packr.New("views", "./views")
	view := jet.NewHTMLSetLoader(packrloader.New(box))
	// ...
}
```
