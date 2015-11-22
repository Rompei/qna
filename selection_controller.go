package main

import (
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"net/http"
	"strconv"
	"time"
)

// IndexSelection returns list of comments
func IndexSelection(req *http.Request, params martini.Params, r render.Render, db *gorm.DB) {
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

	var comments []Selection
	db.Order("id desc").Limit(limit).Offset(offset).Find(&comments)
	r.JSON(200, comments)
}

// GetSelection returns one selection
func GetSelection(params martini.Params, r render.Render, db *gorm.DB) {
	id := params["id"]

	var selection Selection
	db.First(&selection, id)

	if selection.ID == 0 {
		r.JSON(404, Error{404, "Selection was not found."})
	}

	r.JSON(200, selection)
}

// CreateSelection creates a selection
func CreateSelection(selection Selection, r render.Render, db *gorm.DB) {
	db.Create(&selection)
	r.JSON(201, selection)
}

// UpdateSelection update a selection
func UpdateSelection(selection Selection, r render.Render, db *gorm.DB) {
	selection.UpdatedAt = time.Now()
	db.Save(&selection)
	r.JSON(200, selection)
}
