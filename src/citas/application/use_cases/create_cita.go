package usecases

import (
	"apiHospital/src/citas/domain/repositories"
	"apiHospital/src/citas/infastructure/adapters"
	"encoding/json"
	"log"
	"time"
)

type ICita interface {
	Execute(nombrePaciente, apellidoPaciente, numeroContacto, areaCita string, fecha, hora time.Time) error
}

type CreateCita struct {
	db repositories.ICita
}

func NewCreateCita(db repositories.ICita) *CreateCita {
	return &CreateCita{db: db}
}

func (cc *CreateCita) Execute(nombrePaciente, apellidoPaciente, numeroContacto, areaCita string, fecha, hora string) error {

	data := map[string]interface{}{
		"nombrePaciente":   nombrePaciente,
		"apellidoPaciente": apellidoPaciente,
		"numeroContacto":   numeroContacto,
		"areaCita":         areaCita,
		"fecha":            fecha,
		"hora":             hora,
	}

	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to marshal JSON: %s", err)
		return err
	}

	adapters.Execute(body)

	return cc.db.Save(nombrePaciente, apellidoPaciente, numeroContacto, areaCita, fecha, hora)
}
