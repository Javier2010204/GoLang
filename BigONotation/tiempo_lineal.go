package main

import "fmt"

// O(n) => el tiempo de ejecucion va aumentando linealmente segun la cantidad de elementos

func main() {
	
	var arr [5]string =["Lunes", "Martes", "Miercoles", "Jueves", "Viernes"]

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

}