package store

import (
	"final/internal/domain"
)

type StoreInterfaceDentist interface {
	ReadDentist(id int) (domain.Dentist, error)
	CreateDentist(dentist domain.Dentist) error
	UpdateDentist(dentist domain.Dentist) error
	DeleteDentist(id int) error
	}

	type StoreInterfacePatient interface {
		ReadPatient(id int) (domain.Patient, error)
		CreatePatient(patient domain.Patient) error
		UpdatePatient(patient domain.Patient) error
		DeletePatient(id int) error
		}

		type StoreInterfaceAppointment interface {
			ReadAppointment(id int) (domain.Appointment, error)
			CreateAppointment(appointment domain.Appointment) error
			UpdateAppointment(appointment domain.Appointment) error
			DeleteAppointment(id int) error
			ReadAppointmentByPersonalIdNumber(personal_id_number int) (domain.Appointment, error)
			}
