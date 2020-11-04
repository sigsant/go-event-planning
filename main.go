package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	csvline "taskEvent/modules/csvline"
	evento "taskEvent/modules/evento"
)

var errOptionInvalid = errors.New("ERROR: Selección inválida")

var fichero = flag.String("csv", "planning.csv", "Nombre del archivo csv donde se guarda los eventos")
var planningEvent [][]string

// verificarOpcion comprueba si el usuario ha introducido la opción correcta del menú.
// Sale del programa si se ha introducido una opción incorrecta.
func verificarOpcion(sel string) (string, error) {
	opciones := []string{"C", "E", "B", "M", "S"}

	for _, v := range opciones {
		if v == sel {
			return sel, nil
		}
	}
	return " ", errOptionInvalid
}

// mostrarMenu recuerda al usuario las opciones que puede introducir el usuario.
func mostrarMenu() string {
	fmt.Println("\n\tTask Event Manager")
	fmt.Println("\t================")

	fmt.Println("\n\t*(C)rear nuevo evento")
	fmt.Println("\t*(E)ditar evento (Placeholder)")
	fmt.Println("\t*(B)orrar evento (Placeholder)")
	fmt.Println("\t*(M)ostrar eventos guardados")
	fmt.Println("\t*(S)alir del programa")

	lector := bufio.NewScanner(os.Stdin)
	fmt.Print("\n\tSelecciona una de las siguientes opciones: ")
	lector.Scan()
	seleccion := lector.Text()
	fmt.Print("\n")

	return strings.ToUpper(seleccion)
}

func main() {
	flag.Parse()
	opcion := mostrarMenu()

	_, err := verificarOpcion(opcion)
	if err != nil {
		log.Fatal(err)
	}

	switch opcion {
	case "C":
		horario, actividad := evento.CrearActividad()
		horarioFormato, actividadFormato := evento.FormatoEvento(horario, actividad)
		planningEvent = evento.CrearEvento(horarioFormato, actividadFormato)
		evento.CrearCSV(planningEvent, fichero)
	case "M":
		lineasEvento := csvline.ProcesarArchivo(fichero)
		csvline.MostrarEvento(lineasEvento)
	case "S":
		os.Exit(0)
	}
}
