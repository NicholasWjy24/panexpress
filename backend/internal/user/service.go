package user

type Service struct {
	repository *Repository
}

func UserService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetUsers() ([]User, error) {
	return s.repository.FindAll()
}
