package csvline

import (
	"fmt"
)

type CsvLine struct {
	Hora   string
	Nombre string
}

// MostrarEvento muestra linea a linea los eventos guardados en el fichero csv.
func MostrarEvento(lineas [][]string) {
	for _, linea := range lineas {
		data := CsvLine{
			Hora:   linea[0],
			Nombre: linea[1],
		}
		fmt.Println("\n\t", data.Hora)
		fmt.Println("\t", data.Nombre)
	}
}
