package makanan

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func MenuRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindAll() ([]Makanan, error) {

	var makanan []Makanan

	if err := r.db.Find(&makanan).Error; err != nil {
		return nil, err
	}

	return makanan, nil
}

func (r *Repository) CreateMakanan(makanan *Makanan) error {
	return r.db.Create(makanan).Error
}
