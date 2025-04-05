package main

import (
	"fmt"
	"net/http"
	"server/utils"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/paquetes", utils.RecibirPaquetes)
	mux.HandleFunc("/mensaje", utils.RecibirMensaje)

	fmt.Println("Servidor escuchando en https://localhost:8080")

	//panic("no implementado!")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
