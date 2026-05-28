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

func (r *Repository) UpdateUser(
	id string,
	user *User,
) error {

	return r.db.
		Model(&User{}).
		Where("id = ?", id).
		Updates(user).
		Error
}

func (r *Repository) DeleteUser(
	id string,
) error {

	return r.db.
		Delete(&User{}, id).
		Error
}
