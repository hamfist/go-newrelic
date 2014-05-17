package main

import (
	"log"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/meatballhat/go-newrelic/martini/nr"
)

func main() {
	m := martini.Classic()
	err := nr.Configure(m)
	if err != nil {
		log.Fatal(err)
	}

	m.Use(render.Renderer())

	m.Get(`/`, func(r render.Render) {
		r.JSON(200, map[string]string{"whatever": "hosers"})
	})

	m.Run()
}
