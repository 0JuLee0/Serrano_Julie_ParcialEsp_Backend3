package appointment

import (
	"final/internal/domain"
)

type Service interface {
	GetByID(id int) (domain.Appointment, error)
	GetByPersonalIdNumber(personal_id_number int) (domain.Appointment, error)
	Create(appointment domain.Appointment) (domain.Appointment, error)
	Update(id int, appointment domain.Appointment) (domain.Appointment, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domain.Appointment, error) {
	appointment, err := s.r.GetByID(id)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (s *service) GetByPersonalIdNumber(personal_id_number int) (domain.Appointment, error) {
	appointment, err := s.r.GetByPersonalIdNumber(personal_id_number)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (s *service) Create(appointment domain.Appointment) (domain.Appointment, error) {
	appointment, err := s.r.Create(appointment)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (s *service) Update(id int, appointment domain.Appointment) (domain.Appointment, error) {
	appointment, err := s.r.Update(id, appointment)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
