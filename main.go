package main

import (
	"encoding/csv"
	"fmt"
	"os"
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

	fmt.Println(csvLineas)

	//Mostrar y capturar resultados
	for i, v := range csvLineas {
		fmt.Printf("Problem #%d: %s = \n", i+1, v)
		var election int
		fmt.Scan(&election)
	}

}
