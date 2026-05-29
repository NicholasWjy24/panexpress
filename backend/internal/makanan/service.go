package makanan

type Service struct {
	repository *Repository
}

func MenuService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetAllMakanan() ([]Makanan, error) {
	return s.repository.FindAll()
}

type RegisterInputMakanan struct {
	MknnNm string  `json:"mknnnm" binding:"required"`
	MknnTp string  `json:"mknntp" binding:"required"`
	MknnPr float64 `json:"mknnpr" binding:"required"`
	MknnQt int     `json:"mknnqt" binding:"required"`
	MknnSt bool    `json:"mknnst"`
}

func (s *Service) Register(
	input RegisterInputMakanan,
) (*Makanan, error) {

	makanan := &Makanan{
		MknnNm: input.MknnNm,
		MknnTp: input.MknnTp,
		MknnPr: input.MknnPr,
		MknnQt: input.MknnQt,
		MknnSt: input.MknnSt,
	}

	if err := s.repository.CreateMakanan(makanan); err != nil {
		return nil, err
	}

	return makanan, nil
}

func (s *Service) UpdateMakanan(
	id string,
	makanan *Makanan,
) error {

	return s.repository.UpdateMakanan(
		id,
		makanan,
	)
}

func (s *Service) DeleteMakanan(
	id string,
) error {

	return s.repository.DeleteMakanan(id)
}
