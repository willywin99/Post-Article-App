package article

import "gorm.io/gorm"

type Repository struct {
    DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
    return &Repository{DB: db}
}

func (r *Repository) Create(a *Article) error {
    return r.DB.Create(a).Error
}

func (r *Repository) List(limit, offset int, status string) ([]Article, error) {
    var res []Article
    q := r.DB.Order("created_date desc").Limit(limit).Offset(offset)
    if status != "" {
        q = q.Where("status = ?", status)
    }
    if err := q.Find(&res).Error; err != nil {
        return nil, err
    }
    return res, nil
}

func (r *Repository) GetByID(id uint) (*Article, error) {
    var a Article
    if err := r.DB.First(&a, id).Error; err != nil {
        return nil, err
    }
    return &a, nil
}

func (r *Repository) Update(a *Article) error {
    return r.DB.Save(a).Error
}

func (r *Repository) Delete(id uint) error {
    return r.DB.Delete(&Article{}, id).Error
}
