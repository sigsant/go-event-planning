package csvline

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
)

type Csvline struct {
	Hora   string
	Nombre string
}

var (
	errFicheroNoExistente = errors.New("ERROR: El fichero no existe")
	errSinLineas          = errors.New("ERROR: Imposible leer líneas")
)

func ProcesarArchivo(fichero *string) [][]string {
	file, err := os.Open(*fichero)
	if err != nil {
		log.Fatal(errFicheroNoExistente)
	}
	defer file.Close()

	lineas, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(errSinLineas)
	}
	return lineas
}

// MostrarEvento muestra linea a linea los eventos guardados en el fichero csv.
func MostrarEvento(lineas [][]string) {
	for _, linea := range lineas {
		data := Csvline{
			Hora:   linea[0],
			Nombre: linea[1],
		}
		fmt.Println("\n\t", data.Hora)
		fmt.Println("\t", data.Nombre)
	}
}
