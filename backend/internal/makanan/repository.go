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

func (r *Repository) FindByID(id string) (*Makanan, error) {
	var makanan Makanan

	if err := r.db.
		Where("mknnid = ?", id).
		First(&makanan).
		Error; err != nil {
		return nil, err
	}

	return &makanan, nil
}

func (r *Repository) CreateMakanan(makanan *Makanan) error {
	return r.db.Create(makanan).Error
}

func (r *Repository) UpdateMakanan(
	id string,
	makanan *Makanan,
) error {

	return r.db.
		Model(&Makanan{}).
		Where("mknnid = ?", id).
		Updates(makanan).
		Error
}

func (r *Repository) UpdateImage(
	id string,
	image []byte,
) error {

	return r.db.
		Model(&Makanan{}).
		Where("mknnid = ?", id).
		Update("mknnig", image).
		Error
}

func (r *Repository) DeleteMakanan(
	id string,
) error {

	return r.db.
		Where("mknnid = ?", id).
		Delete(&Makanan{}).
		Error
}
