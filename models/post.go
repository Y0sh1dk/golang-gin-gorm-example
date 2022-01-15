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

// TODO(yoshi): Take int not string!
func GetPostByID(db *gorm.DB, post *Post, id string) error {
	if err := db.Where("id = ?", id).Find(post).Error; err != nil {
		return err
	}
	return nil
}

func CreatePost(db *gorm.DB, post *Post) error {
	if err := db.Create(post).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePost(db *gorm.DB, post *Post) error {
	post.UpdatedAt = time.Now()
	if err := db.Model(&post).Updates(post).Error; err != nil {
		return err
	}
	return nil
}

// Delete post AND comments
func DeletePost(db *gorm.DB, post *Post, id string) error {
	// TODO(yoshi): Delete comments too
	if err := GetPostByID(db, post, id); err != nil {
		return err
	}
	db.Delete(post)
	return nil
}
