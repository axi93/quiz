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

	finals := parseLines(csvLineas)

	isOK := 0
	//Mostrar y capturar resultados
	for i, v := range finals {
		fmt.Printf("Problem #%d: %s = \n", i+1, v.a)
		var respuesta string
		fmt.Scanf("%s\n", &respuesta)
		if respuesta == v.b {
			isOK++
		}

	}
	fmt.Printf("You scored %d out of %d.\n", isOK, len(finals))
}

func parseLines(csvLineas [][]string) []final {
	ret := make([]final, len(csvLineas))
	for i, linea := range csvLineas {
		ret[i] = final{
			a: linea[0],
			b: strings.TrimSpace(linea[1]),
		}
	}
	return ret
}

type final struct {
	a string
	b string
}
