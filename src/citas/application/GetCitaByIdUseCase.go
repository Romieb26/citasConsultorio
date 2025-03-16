// GetCitaByIdUseCase.go
package application

import (
	repositories "apiHospital/src/citas/domain"
	"apiHospital/src/citas/domain/entities"
)

type GetCitaByIdUseCase struct {
	repo repositories.ICita
}

func NewGetCitaByIdUseCase(repo repositories.ICita) *GetCitaByIdUseCase {
	return &GetCitaByIdUseCase{repo: repo}
}

func (uc *GetCitaByIdUseCase) Run(id int32) (*entities.Cita, error) {
	cita, err := uc.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return cita, nil
}
