package menu

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindAll() ([]Menu, error) {
	var menus []Menu
	if err := r.db.Table("menus").Find(&menus).Error; err != nil {
		return nil, err
	}

	return menus, nil
}
