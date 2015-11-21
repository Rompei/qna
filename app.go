package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func main() {
	// Database migration
	db := GetDB()
	db.AutoMigrate(&Question{}, &Comment{}, &Selection{})
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Map(db)
	m.Group("/questions", func(r martini.Router) {
		r.Get("/", IndexQuestion)
		r.Get("/:id", GetQuestion)
		r.Post("/create", binding.Bind(Question{}), CreateQuestion)
		r.Put("/update/:id", UpdateQuestion)
	})
	m.Group("/comments", func(r martini.Router) {
		r.Get("/", IndexComment)
		r.Get("/:id", GetComment)
		r.Post("/create", binding.Bind(Comment{}), CreateComment)
		r.Put("/update/:id", UpdateComment)
	})
	m.Group("/selections", func(r martini.Router) {
		r.Get("/", IndexSelection)
		r.Get("/:id", GetSelection)
		r.Post("/create", binding.Bind(Selection{}), CreateSelection)
		r.Put("/update/:id", UpdateSelection)
	})
	m.Run()
}
