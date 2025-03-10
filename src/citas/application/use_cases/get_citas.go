package usecases

import (
	"apiHospital/src/citas/domain/entities"
	"apiHospital/src/citas/domain/repositories"
)

type GetCita struct {
	db repositories.ICita
}

func NewViewCita(db repositories.ICita) *GetCita {
	return &GetCita{db: db}
}

func (vc *GetCita) Execute() ([]entities.Cita, error) {
	return vc.db.GetAll()
}
