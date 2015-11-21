package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"net/http"
	"strconv"
)

// IndexQuestion returns list of questions
func IndexQuestion(req *http.Request, params martini.Params, r render.Render, db *gorm.DB) {
	limit := 40
	offset := 0
	query := req.URL.Query()
	rawPage, rawMaxResults := query.Get("page"), query.Get("maxResults")
	page, err := strconv.Atoi(rawPage)
	maxResults, err := strconv.Atoi(rawMaxResults)
	if err != nil {
		fmt.Println(err)
		r.JSON(400, Error{400, "page and maxResults must be integer."})
	}

	limit = maxResults
	offset = (page - 1) * maxResults

	var questions []Question
	db.Limit(limit).Offset(offset).Find(&questions)
	for i, v := range questions {
		db.Model(&v).Related(&questions[i].Comments)
		db.Model(&v).Related(&questions[i].Selections)
	}
	r.JSON(200, questions)
}

// GetQuestion returns one question
func GetQuestion(params martini.Params, r render.Render, db *gorm.DB) {
	id := params["id"]

	var question Question
	db.First(&question, id)
	db.Model(&question).Related(&question.Comments)
	db.Model(&question).Related(&question.Selections)

	if question.ID == 0 {
		r.JSON(404, Error{404, "Question was not found."})
	}

	r.JSON(200, question)
}

// CreateQuestion creates a question
func CreateQuestion(question Question, r render.Render, db *gorm.DB) {
	db.Create(&question)
	r.JSON(201, question)
}

// UpdateQuestion update a question
func UpdateQuestion(question Question, r render.Render, db *gorm.DB) {
	db.Save(&question)
	r.JSON(200, question)
}
