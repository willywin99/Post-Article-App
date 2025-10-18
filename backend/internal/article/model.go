package article

import "time"

type Article struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Title       string    `json:"title" gorm:"size:200;not null"`
    Content     string    `json:"content" gorm:"type:text;not null"`
    Category    string    `json:"category" gorm:"size:100;not null"`
    Status      string    `json:"status" gorm:"size:100;not null"`
    CreatedDate time.Time `json:"created_date" gorm:"autoCreateTime"`
    UpdatedDate time.Time `json:"updated_date" gorm:"autoUpdateTime"`
}
