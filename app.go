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
		r.Get("", IndexQuestion)
		r.Get("/:id", GetQuestion)
		r.Post("", binding.Bind(Question{}), CreateQuestion)
		r.Put("/:id", binding.Bind(Question{}), UpdateQuestion)
	})
	m.Group("/comments", func(r martini.Router) {
		r.Get("", IndexComment)
		r.Get("/:id", GetComment)
		r.Post("", binding.Bind(Comment{}), CreateComment)
		r.Put("/:id", binding.Bind(Comment{}), UpdateComment)
	})
	m.Group("/selections", func(r martini.Router) {
		r.Get("", IndexSelection)
		r.Get("/:id", GetSelection)
		r.Post("", binding.Bind(Selection{}), CreateSelection)
		r.Put("/:id", binding.Bind(Selection{}), UpdateSelection)
	})
	m.Run()
}
