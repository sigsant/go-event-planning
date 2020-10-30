/*
	Create a program that will allow you to enter events organizable by hour.
	There must be menu options of some form, and you must be able to easily edit, add, and delete events
	without directly changing the source code.

	source: https://www.reddit.com/r/dailyprogrammer/comments/pihtx/intermediate_challenge_1/
*/

package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"taskEvent/modules/evento"
)

var (
	errOptionInvalid = errors.New("ERROR: Selección inválida")
)
var planningEvent map[string]string

/*
 * crearEvento guarda el evento creado en un array de 2 dimensiones
 * y lo guarda posteriormente en un archivo CSV.
 */
func crearEvento(horario string, evento string) {
	var planningEvent [][]string
	var row []string

	row = []string{horario, evento}
	planningEvent = append(planningEvent, row)

	// return planningEvent
	// //TODO? Usar como debug aqui y reutilizarlo para el menú mostrar
	// for i := range planningEvent {
	// 	fmt.Println("\t", planningEvent[1][i])
	// }
	crearCSV(planningEvent)
}

func crearCSV(planning [][]string) {
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

/*
 * verificarOpcion comprueba si el usuario ha introducido la opción correcta del menú.
 * Sale del programa si se ha introducido una opción incorrecta.
 */
func verificarOpcion(sel string) (string, error) {
	// opciones := []string{"Mostrar", "Agregar", "Editar", "Borrar"}
	opciones := []string{"C", "E", "B", "M"}

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

	test := evento.Horario{8, 25}
	horario, evento := evento.NuevoEvento(test, "Mirar")
	crearEvento(horario, evento)

	// ioutil.WriteFile("Planning.txt", []byte(planning), 0644)

}
