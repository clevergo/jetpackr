package main

import (
	"bytes"
	"log"

	"github.com/CloudyKit/jet/v3"
	packrloader "github.com/clevergo/jet-packrloader"
	"github.com/gobuffalo/packr/v2"
)

func main() {
	box := packr.New("views", "./views")
	view := jet.NewHTMLSetLoader(packrloader.New(box))
	wr := bytes.Buffer{}
	tmpl, err := view.GetTemplate("index.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	if err = tmpl.Execute(&wr, nil, nil); err != nil {
		log.Fatal(err)
	}

	log.Println(wr.String())
}
