package handler

import (
	"errors"
	"strconv"

	"final/internal/dentist"
	"final/internal/domain"
	"final/pkg/web"

	"github.com/gin-gonic/gin"
)

type dentistHandler struct {
	s dentist.Service
}

func NewDentistHandler(s dentist.Service) *dentistHandler {
	return &dentistHandler{
		s: s,
	}
}

func (h *dentistHandler) GetDentistByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}

		dentist, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentist not found"))
			return
		}

		web.Success(c, 200, dentist)
	}
}

func (h *dentistHandler) PostDentist() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dentist domain.Dentist
		err := c.ShouldBindJSON(&dentist)

		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}

		valid, err := validateEmptyFieldsDentist(&dentist)
		if !valid {
			web.Failure(c, 400, err)
			return
		}

		d, err := h.s.Create(dentist)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, d)
	}
}

func (h *dentistHandler) PutDentist() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		var dentist domain.Dentist
		err = c.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptyFieldsDentist(&dentist)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.s.Update(id, dentist)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, d)
	}
}

func (h *dentistHandler) PatchDentist() gin.HandlerFunc {
	return func(c *gin.Context) {
		var DentistNew domain.DentistDTO
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		if err := c.ShouldBindJSON(&DentistNew); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Dentist{
			Name:               DentistNew.Name,
			LastName:           DentistNew.LastName,
			RegistrationNumber: DentistNew.RegistrationNumber,
		}

		d, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, d)
	}
}

func (h *dentistHandler) DeleteDentist() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}
