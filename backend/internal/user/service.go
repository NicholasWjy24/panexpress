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

func (s *Service) UpdateUser(
	id string,
	user *User,
) error {

	return s.repository.UpdateUser(
		id,
		user,
	)
}

func (s *Service) DeleteUser(
	id string,
) error {

	return s.repository.DeleteUser(id)
}
