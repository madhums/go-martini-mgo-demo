package main

import (
	"github.com/go-martini/martini"
	"github.com/madhums/go-martini-mgo-demo/handlers/articles"
	"github.com/madhums/go-martini-mgo-demo/middlewares"
	"github.com/madhums/go-martini-mgo-demo/models"
	"github.com/martini-contrib/binding"
)

/**
 * Main
 *
 * - use middlewares
 * - define routes
 * - listen and serve
 */

func main() {

	// Initialize
	m := martini.Classic()

	// Connect to mongo
	m.Use(middlewares.Connect())

	// Templating support
	m.Use(middlewares.Templates())

	// Routes
	m.Get("/", articles.List)
	m.Get("/new", articles.AddEdit)
	m.Post("/articles", binding.Form(models.Article{}), articles.Add)
	m.Get("/articles/:_id", articles.Show)

	// Start listening
	m.Run()
}
