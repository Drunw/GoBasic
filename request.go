package main

import "fmt"

type Request struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

func (r *Request) String() string {
	return fmt.Sprintf("Nombre: %s, Apellido: %s", r.Nombre, r.Apellido)
}
