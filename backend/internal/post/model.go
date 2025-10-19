package post

import (
	"strings"
	"time"
)

type Post struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title        string    `json:"title" gorm:"type:varchar(200)"`
	Content      string    `json:"content" gorm:"type:text"`
	Category     string    `json:"category" gorm:"type:varchar(100)"`
	Status       string    `json:"status" gorm:"type:varchar(100)"`
	Created_date time.Time `json:"created_date" gorm:"autoCreateTime"`
	Updated_date time.Time `json:"updated_date" gorm:"autoUpdateTime"`
}

// Validate checks if the post meets all requirements
func (p *Post) Validate() map[string]string {
	errors := make(map[string]string)

	// Title validation
	if strings.TrimSpace(p.Title) == "" {
		errors["title"] = "Title is required"
	} else if len(p.Title) < 20 {
		errors["title"] = "Title must be at least 20 characters"
	}

	// Content validation
	if strings.TrimSpace(p.Content) == "" {
		errors["content"] = "Content is required"
	} else if len(p.Content) < 200 {
		errors["content"] = "Content must be at least 200 characters"
	}

	// Category validation
	if strings.TrimSpace(p.Category) == "" {
		errors["category"] = "Category is required"
	} else if len(p.Category) < 3 {
		errors["category"] = "Category must be at least 3 characters"
	}

	// Status validation
	validStatus := map[string]bool{"publish": true, "draft": true, "thrash": true}
	if strings.TrimSpace(p.Status) == "" {
		errors["status"] = "Status is required"
	} else if !validStatus[strings.ToLower(p.Status)] {
		errors["status"] = "Status must be 'publish', 'draft', or 'thrash'"
	}

	return errors
}
