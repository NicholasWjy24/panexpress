package minuman

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func MinumanRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindAll() ([]Minuman, error) {

	var minuman []Minuman

	if err := r.db.Find(&minuman).Error; err != nil {
		return nil, err
	}

	return minuman, nil
}

func (r *Repository) CreateMinuman(minuman *Minuman) error {
	return r.db.Create(minuman).Error
}

func (r *Repository) UpdateMinuman(
	id string,
	minuman *Minuman,
) error {

	return r.db.
		Model(&Minuman{}).
		Where("mimnid = ?", id).
		Updates(minuman).
		Error
}

func (r *Repository) DeleteMinuman(
	id string,
) error {

	return r.db.
		Where("mimnid = ?", id).
		Delete(&Minuman{}).
		Error
}
