package main

import (
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"net/http"
	"strconv"
	"time"
)

// IndexComment returns list of comments
func IndexComment(req *http.Request, params martini.Params, r render.Render, db *gorm.DB) {
	limit := 40
	offset := 0
	query := req.URL.Query()
	rawPage, rawMaxResults := query.Get("page"), query.Get("maxResults")
	page, err := strconv.Atoi(rawPage)
	maxResults, err := strconv.Atoi(rawMaxResults)
	if err != nil {
		r.JSON(400, Error{400, "page and maxResults must be integer."})
	}

	limit = maxResults
	offset = (page - 1) * maxResults

	var comments []Comment
	db.Order("id desc").Limit(limit).Offset(offset).Find(&comments)
	r.JSON(200, comments)
}

// GetComment returns one comment
func GetComment(params martini.Params, r render.Render, db *gorm.DB) {
	id := params["id"]

	var comment Comment
	db.First(&comment, id)

	if comment.ID == 0 {
		r.JSON(404, Error{404, "Comment was not found."})
	}

	r.JSON(200, comment)
}

// CreateComment creates a comment
func CreateComment(comment Comment, r render.Render, db *gorm.DB) {
	db.Create(&comment)
	r.JSON(201, comment)
}

// UpdateComment update a comment
func UpdateComment(comment Comment, r render.Render, db *gorm.DB) {
	comment.UpdatedAt = time.Now()
	db.Save(&comment)
	r.JSON(200, comment)
}
