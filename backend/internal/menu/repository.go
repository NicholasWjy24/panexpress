package menu

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func MenuRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindByRoleLevel(roleLevel int) ([]Menu, error) {
	var menus []Menu
	if err := r.db.Table("menus").
		Where("min_role_level <= ? AND max_role_level >= ?", roleLevel, roleLevel).
		Find(&menus).Error; err != nil {
		return nil, err
	}

	return menus, nil
}

func (r *Repository) FindAll() ([]Menu, error) {

	var menus []Menu

	if err := r.db.Find(&menus).Error; err != nil {
		return nil, err
	}

	return menus, nil
}

func (r *Repository) CreateMenu(menu *Menu) error {
	return r.db.Create(menu).Error
}
