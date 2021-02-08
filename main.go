package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Abro el fichero
	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Error al abrir documento" + err.Error())
		return
	}
	//Leo las lineas
	csvLineas, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Error al leer las lineas" + err.Error())
		return
	}

	final := parseLines(csvLineas)

	fmt.Println(final)

	//Mostrar y capturar resultados
	for i, v := range csvLineas {
		fmt.Printf("Problem #%d: %s = \n", i+1, v)
		var election int
		fmt.Scan(&election)
	}

}

func parseLines(csvLineas [][]string) []final {
	devolvemos := make([]final, len(csvLineas))
	for i, linea := range csvLineas {
		devolvemos[i] = final{
			a: linea[0],
			b: strings.TrimSpace(linea[1]),
		}
	}
	return devolvemos
}

type final struct {
	a string
	b string
}
