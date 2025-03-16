// citas_repository.go
package repositories

import (
	"apiHospital/src/citas/domain/entities"
)

type ICita interface {
	Save(cita *entities.Cita) error
	Update(cita *entities.Cita) error
	Delete(id int32) error
	GetById(id int32) (*entities.Cita, error)
	GetAll() ([]entities.Cita, error)
}
