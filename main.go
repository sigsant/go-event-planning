package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	evento "taskEvent/modules/evento"
)

var (
	errOptionInvalid = errors.New("ERROR: Selección inválida")
)
var (
	planningEvent [][]string
)

/*
 * verificarOpcion comprueba si el usuario ha introducido la opción correcta del menú.
 * Sale del programa si se ha introducido una opción incorrecta.
 */
func verificarOpcion(sel string) (string, error) {
	// opciones := []string{"Mostrar", "Agregar", "Editar", "Borrar", "Salir"}
	opciones := []string{"C", "E", "B", "M", "S"}

	for _, v := range opciones {
		if v == strings.ToUpper(sel) {
			return sel, nil
		}
	}
	return " ", errOptionInvalid
}

/*
 * mostrarMenu recuerda al usuario las opciones que puede introducir el usuario.
 */
func mostrarMenu() {
	fmt.Println("\n\tTask Event Manager")
	fmt.Println("\t================")

	fmt.Println("\n\t*(C)rear nuevo evento (Placeholder)")
	fmt.Println("\t*(E)ditar evento (Placeholder)")
	fmt.Println("\t*(B)orrar evento (Placeholder)")
	fmt.Println("\t*(M)ostrar eventos guardados (Placeholder)")
	fmt.Println("\t*(S)alir del programa")
}

func main() {
	mostrarMenu()

	lector := bufio.NewScanner(os.Stdin)
	fmt.Print("\n\tSelecciona una de las siguientes opciones: ")
	lector.Scan()
	opcion := lector.Text()
	fmt.Print("\n")

	_, err := verificarOpcion(opcion)
	if err != nil {
		log.Fatal(err)
	}

	switch opcion {
	case "C":
		horario, actividad := evento.CrearActividad()
		horarioFormato, actividadFormato := evento.FormatoEvento(horario, actividad)
		planningEvent = evento.CrearEvento(horarioFormato, actividadFormato)
	case "S":
		os.Exit(0)
	}
}
