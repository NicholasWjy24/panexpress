package menu

type Service struct {
	repository *Repository
}

func MenuService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetMenus() ([]Menu, error) {
	return s.repository.FindAll()
}
