package main

import (
	"github.com/CloudyKit/jet/v3"
	"clevergo.tech/clevergo"
	"clevergo.tech/jetpackr"
	"github.com/gobuffalo/packr/v2"
)

func main() {
	box := packr.New("views", "./views")
	view := jet.NewHTMLSetLoader(jetpackr.New(box))
	app := clevergo.New()
	app.Get("/", func(ctx *clevergo.Context) error {
		tmpl, err := view.GetTemplate("index.tmpl")
		if err != nil {
			return err
		}

		return tmpl.Execute(ctx.Response, nil, nil)
	})
	app.Run(":8080")
}
