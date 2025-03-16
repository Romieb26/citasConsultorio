// GetAllCitasUseCase.go
package application

import (
	repositories "apiHospital/src/citas/domain"
	"apiHospital/src/citas/domain/entities"
)

type GetAllCitasUseCase struct {
	repo repositories.ICita
}

func NewGetAllCitasUseCase(repo repositories.ICita) *GetAllCitasUseCase {
	return &GetAllCitasUseCase{repo: repo}
}

func (uc *GetAllCitasUseCase) Run() ([]entities.Cita, error) {
	citas, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return citas, nil
}
