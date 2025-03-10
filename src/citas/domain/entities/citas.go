package entities

type Cita struct {
	CitaID           int32  `json:"citaId" gorm:"column:cita_id;primaryKey;autoIncrement"`
	NombrePaciente   string `json:"nombrePaciente" gorm:"column:nombre_paciente;not null"`
	ApellidoPaciente string `json:"apellidoPaciente" gorm:"column:apellido_paciente;not null"`
	NumeroContacto   string `json:"numeroContacto" gorm:"column:numero_contacto;not null"`
	AreaCita         string `json:"areaCita" gorm:"type:enum('Medicina General', 'Pediatría', 'Oftalmología', 'Dermatología', 'Cardiología');not null"`
	Fecha            string `json:"fecha" gorm:"column:fecha;not null"`
	Hora             string `json:"hora" gorm:"column:hora;not null"`
}

func NewCita(nombrePaciente, apellidoPaciente, numeroContacto, areaCita, fecha, hora string) *Cita {
	return &Cita{
		NombrePaciente:   nombrePaciente,
		ApellidoPaciente: apellidoPaciente,
		NumeroContacto:   numeroContacto,
		AreaCita:         areaCita,
		Fecha:            fecha,
		Hora:             hora,
	}
}
