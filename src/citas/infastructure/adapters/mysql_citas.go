package adapters

import (
	"apiHospital/src/citas/domain/entities"
	"apiHospital/src/citas/domain/repositories"

	"gorm.io/gorm"
)

type MySQLRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) repositories.ICita {
	return &MySQLRepository{db: db}
}

func (repo *MySQLRepository) Save(nombrePaciente, apellidoPaciente, numeroContacto, areaCita string, fecha, hora string) error {
	cita := entities.NewCita(nombrePaciente, apellidoPaciente, numeroContacto, areaCita, fecha, hora)
	result := repo.db.Create(cita)
	return result.Error
}

func (repo *MySQLRepository) GetAll() ([]entities.Cita, error) {
	var citas []entities.Cita
	result := repo.db.Find(&citas)
	return citas, result.Error
}

func (repo *MySQLRepository) GetById(id int32) (*entities.Cita, error) {
	var cita entities.Cita
	result := repo.db.First(&cita, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &cita, nil
}

func (repo *MySQLRepository) Delete(id int32) error {
	var cita entities.Cita
	result := repo.db.First(&cita, id)
	if result.Error != nil {
		return result.Error
	}
	result = repo.db.Delete(&cita)
	return result.Error
}
