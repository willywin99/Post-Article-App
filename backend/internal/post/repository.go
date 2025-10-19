package post

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) Create(p *Post) error {
	return r.DB.Create(p).Error
}

func (r *Repository) List(limit, offset int, status string) ([]Post, error) {
	var res []Post
	q := r.DB.Order("created_date desc").Limit(limit).Offset(offset)
	if status != "" {
		q = q.Where("status = ?", status)
	}
	if err := q.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Repository) GetByID(id uint) (*Post, error) {
	var p Post
	if err := r.DB.First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *Repository) Update(p *Post) error {
	return r.DB.Save(p).Error
}

func (r *Repository) Delete(id uint) error {
	return r.DB.Delete(&Post{}, id).Error
}
