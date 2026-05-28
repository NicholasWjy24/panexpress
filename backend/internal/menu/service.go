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
