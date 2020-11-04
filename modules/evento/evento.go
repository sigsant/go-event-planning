package evento

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Horario struct {
	Horas   int
	Minutos int
}
type Evento struct {
	Horario Horario
	Nombre  string
}

// CrearActividad solicita al usuario los datos de inicio (horas, minutos) y nombre de la actividad
func CrearActividad() (Horario, string) {
	entradaDatos := bufio.NewScanner(os.Stdin)
	fmt.Print("\n\tIntroduce la hora de inicio de la actividad: ")
	entradaDatos.Scan()
	horaInt, _ := strconv.Atoi(entradaDatos.Text())
	fmt.Print("\tIntroduce los minutos de inicio de la actividad: ")
	entradaDatos.Scan()
	minutoInt, _ := strconv.Atoi(entradaDatos.Text())
	fmt.Print("\tIntroduce la actividad a realizar: ")
	entradaDatos.Scan()
	actividad := entradaDatos.Text()
	return Horario{horaInt, minutoInt}, actividad
}

// FormatoEvento modifica el horario según el sistema de 24 horas (00:00 - 23:59).
// Crea una plantilla base para mostrar en el fichero CSV.
// Devuelve los valores horario y nombre del evento en formato string.
func FormatoEvento(h Horario, n string) (string, string) {
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

	return horaPlanning, nombrePlanning
}

// CrearEvento guarda el evento creado en un array de 2 dimensiones.
// Devuelve el array multidimensional {hora, evento}
func CrearEvento(horario string, evento string) [][]string {
	var planningEvent [][]string
	var row []string

	row = []string{horario, evento}
	planningEvent = append(planningEvent, row)

	// //TODO? Usar como debug aqui y reutilizarlo para el menú mostrar
	// for i := range planningEvent {
	// 	fmt.Println("\t", planningEvent[1][i])
	// }
	CrearCSV(planningEvent)
	return planningEvent
}

// CrearCSV crea un fichero CSV con los datos de los eventos guardados en un array multidmensional.
// En caso de existir previamente, añade la información al fichero.
func CrearCSV(planning [][]string) {
	csvFile, err := os.OpenFile("planning.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("No ha sido posible crear el fichero %s.", err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	for _, filaCsv := range planning {
		_ = csvwriter.Write(filaCsv)
	}
	csvwriter.Flush()
}
