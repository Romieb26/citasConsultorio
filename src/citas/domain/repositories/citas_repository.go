package repositories

import (
	"apiHospital/src/citas/domain/entities"
)

type ICita interface {
	Save(nombrePaciente, apellidoPaciente, numeroContacto, areaCita string, fecha, hora string) error
	GetAll() ([]entities.Cita, error)
	GetById(id int32) (*entities.Cita, error)
	Delete(id int32) error
}
