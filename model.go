package main

import "time"

// Question : model of question
type Question struct {
	ID         int         `json:"id" gorm:"primary_key"`
	Title      string      `json:"title"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
	Selections []Selection `json:"selections"`
	Comments   []Comment   `json:"comments"`
}

// Selection : model of selection
type Selection struct {
	ID         int       `json:"id" gorm:"primary_key"`
	QuestionID int       `json:"questionId" sql:"index"`
	Content    string    `json:"content"`
	Count      int       `json:"count"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

// Comment : model of comment
type Comment struct {
	ID         int       `json:"id" gorm:"primary_key"`
	QuestionID int       `json:"questionId" sql:"index"`
	Content    string    `json:"content" sql:"type:text"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
