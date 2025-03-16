// CreateCitaUseCase.go
package application

import (
	repositories "apiHospital/src/citas/domain"
	"apiHospital/src/citas/domain/entities"
)

type CreateCitaUseCase struct {
	repo repositories.ICita
}

func NewCreateCitaUseCase(repo repositories.ICita) *CreateCitaUseCase {
	return &CreateCitaUseCase{repo: repo}
}

func (uc *CreateCitaUseCase) Run(cita *entities.Cita) (*entities.Cita, error) {
	err := uc.repo.Save(cita)
	if err != nil {
		return nil, err
	}
	return cita, nil
}
