package evento

import (
	"fmt"
	"strconv"
)

func init() {
	fmt.Println("\n\t Se ha inicializado el paquete Eventos correctamente")
}

type Horario struct {
	Horas   int
	Minutos int
}
type Evento struct {
	Horario     Horario
	Nombre      string
	Descripcion string
}

//Crear
func NuevoEvento(h Horario, n string) (string, string) {
	horaCero := "0" + strconv.Itoa(h.Horas)
	minutoCero := "0" + strconv.Itoa(h.Minutos)
	var horaPlanning, nombrePlanning string

	if h.Horas < 10 && h.Minutos < 10 {
		horaPlanning = fmt.Sprintf("Hora: %s:%s", horaCero, minutoCero)
		nombrePlanning = fmt.Sprintf("Actividad: %s", n)
	} else if h.Horas < 10 {
		horaPlanning = fmt.Sprintf("Hora: %s:%d", horaCero, h.Minutos)
		nombrePlanning = fmt.Sprintf("Actividad: %s", n)

	} else if h.Minutos < 10 {
		horaPlanning = fmt.Sprintf("Hora: %d:%s", h.Horas, minutoCero)
		nombrePlanning = fmt.Sprintf("Actividad: %s", n)

	} else {
		horaPlanning = fmt.Sprintf("Hora: %d:%d", h.Horas, h.Minutos)
		nombrePlanning = fmt.Sprintf("Actividad: %s", n)
	}

	// texto := fmt.Sprintf("%s%s", horaPlanning, nombrePlanning)

	return horaPlanning, nombrePlanning
}
