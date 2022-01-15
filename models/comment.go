package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Comment struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Author    string    `gorm:"not null" json:"author"`
	Body      string    `gorm:"not null" json:"body"`
	Post      Post      `json:"post"`
	PostID    uint      `gorm:"not null" json:"post_id"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Comment) Prepare() {
	c.Post = Post{}
	c.CreatedAt = time.Now()
}

func GetComments(db *gorm.DB, comments *[]Comment) error {
	return nil
}

func GetCommentByID(db *gorm.DB, comment *Comment) error {
	return nil
}

func CreateComment(db *gorm.DB, comment *Comment) error {
	return nil
}

func UpdateComment(db *gorm.DB, comment *Comment) error {
	return nil
}

func DeleteComment(db *gorm.DB, comment *Comment) error {
	return nil
}
