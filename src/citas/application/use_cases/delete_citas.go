package usecases

import (
	"apiHospital/src/citas/domain/repositories"
)

type DeleteCita struct {
	repo repositories.ICita
}

func NewDeleteCita(repo repositories.ICita) *DeleteCita {
	return &DeleteCita{repo: repo}
}

func (uc *DeleteCita) Execute(id int32) error {
	return uc.repo.Delete(id)
}
