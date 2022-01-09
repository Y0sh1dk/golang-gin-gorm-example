package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Post struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Author    string    `gorm:"not null" json:"author"`
	Body      string    `gorm:"not null" json:"body"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Post) Prepare() {
	p.CreatedAt = time.Now()
}

func GetPosts(db *gorm.DB, posts *[]Post) error {
	if err := db.Find(posts).Error; err != nil {
		return err
	}
	return nil
}

func (p *Post) GetPostByID(db *gorm.DB, post *Post, id string) error {
	return nil
}

func (p *Post) CreatePost(db *gorm.DB, post *Post) error {
	return nil
}

func (p *Post) UpdatePost(db *gorm.DB, post *Post) error {
	// Set Post.UpdatedAt
	return nil
}

// Delete post AND comments
func (p *Post) DeletePost(db *gorm.DB, post *Post) error {
	return nil
}
