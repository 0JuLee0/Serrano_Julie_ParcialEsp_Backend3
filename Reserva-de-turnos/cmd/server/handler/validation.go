package handler

import (
	"errors"
	"final/internal/domain"
)

func validateEmptyFieldsDentist(dentist *domain.Dentist) (bool, error) {
	if dentist.Name == "" || dentist.LastName == "" || dentist.RegistrationNumber == "" {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

func validateEmptyFieldsPatient(patient *domain.Patient) (bool, error) {
	if patient.Name == "" || patient.LastName == "" || patient.Address == "" || patient.Personal_id_number == 0 || patient.Creation_date == "" {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

func validateEmptyFieldsAppointment(appointment *domain.Appointment) (bool, error) {
	if appointment.Dentist.RegistrationNumber == "" || appointment.Patient.Personal_id_number == 0 || appointment.Appointment_date == "" || appointment.Appointment_time == "" || appointment.Description == "" {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}
