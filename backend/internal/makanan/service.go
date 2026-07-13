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
	MknnIg []byte  `json:"-"`
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
		MknnIg: input.MknnIg,
	}

	if err := s.repository.CreateMakanan(makanan); err != nil {
		return nil, err
	}

	return makanan, nil
}

func (s *Service) GetMakananImage(id string) ([]byte, error) {
	makanan, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return makanan.MknnIg, nil
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

func (s *Service) UpdateMakananImage(
	id string,
	image []byte,
) error {

	return s.repository.UpdateImage(
		id,
		image,
	)
}

func (s *Service) DeleteMakanan(
	id string,
) error {

	return s.repository.DeleteMakanan(id)
}
