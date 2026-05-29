package minuman

type Service struct {
	repository *Repository
}

func MinumanService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetAllMinuman() ([]Minuman, error) {
	return s.repository.FindAll()
}

type RegisterInputMinuman struct {
	MimnNm string  `json:"mimnnm" binding:"required"`
	MimnTp string  `json:"mimntp" binding:"required"`
	MimnPr float64 `json:"mimnpr" binding:"required"`
	MimnQt int     `json:"mimnqt" binding:"required"`
	MimnSt bool    `json:"mimnst"`
}

func (s *Service) Register(
	input RegisterInputMinuman,
) (*Minuman, error) {

	minuman := &Minuman{
		MimnNm: input.MimnNm,
		MimnTp: input.MimnTp,
		MimnPr: input.MimnPr,
		MimnQt: input.MimnQt,
		MimnSt: input.MimnSt,
	}

	if err := s.repository.CreateMinuman(minuman); err != nil {
		return nil, err
	}

	return minuman, nil
}

func (s *Service) UpdateMinuman(
	id string,
	makanan *Minuman,
) error {

	return s.repository.UpdateMinuman(
		id,
		makanan,
	)
}

func (s *Service) DeleteMinuman(
	id string,
) error {

	return s.repository.DeleteMinuman(id)
}
