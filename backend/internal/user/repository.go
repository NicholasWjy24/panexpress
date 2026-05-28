package user

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func UserRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindAll() ([]User, error) {

	var users []User

	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
