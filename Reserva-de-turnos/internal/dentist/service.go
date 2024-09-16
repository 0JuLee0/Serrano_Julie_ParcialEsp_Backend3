package dentist

import (
	"final/internal/domain"
)

type Service interface {
	GetByID(id int) (domain.Dentist, error)
	Create(dentist domain.Dentist) (domain.Dentist, error)
	Update(id int, dentist domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domain.Dentist, error) {
	dentist, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *service) Create(dentist domain.Dentist) (domain.Dentist, error) {
	dentist, err := s.r.Create(dentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *service) Update(id int, dentist domain.Dentist) (domain.Dentist, error) {
	dentist, err := s.r.Update(id, dentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
