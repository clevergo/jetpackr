package main

import (
	"net/http"

	"github.com/CloudyKit/jet/v3"
	"github.com/clevergo/clevergo"
	"github.com/clevergo/jetpackr"
	"github.com/gobuffalo/packr/v2"
)

func main() {
	box := packr.New("views", "./views")
	view := jet.NewHTMLSetLoader(jetpackr.New(box))
	router := clevergo.NewRouter()
	router.Get("/", func(ctx *clevergo.Context) error {
		tmpl, err := view.GetTemplate("index.tmpl")
		if err != nil {
			return err
		}

		return tmpl.Execute(ctx.Response, nil, nil)
	})
	http.ListenAndServe(":8080", router)
}
