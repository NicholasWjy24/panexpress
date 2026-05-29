package menu

type Service struct {
	repository *Repository
}

func MenuService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetMenus(roleLevel int) ([]Menu, error) {
	return s.repository.FindByRoleLevel(roleLevel)
}

func (s *Service) GetAllMenus() ([]Menu, error) {
	return s.repository.FindAll()
}

type RegisterInputMenu struct {
	MenuName     string `json:"menu_name" binding:"required"`
	Route        string `json:"route" binding:"required"`
	MinRoleLevel int    `json:"min_role_level" binding:"required"`
	MaxRoleLevel int    `json:"max_role_level" binding:"required"`
}

func (s *Service) Register(input RegisterInputMenu) (*Menu, error) {

	menu := &Menu{
		MenuName:     input.MenuName,
		Route:        input.Route,
		MinRoleLevel: input.MinRoleLevel,
		MaxRoleLevel: input.MaxRoleLevel,
	}

	if err := s.repository.CreateMenu(menu); err != nil {
		return nil, err
	}

	return menu, nil
}
