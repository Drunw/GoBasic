package main

import "strings"

func convertirMayuscula(r Request) Request {
	nombre := r.Nombre
	apellido := r.Apellido
	r.Apellido = strings.ToUpper(string(apellido[0])) + apellido[1:]
	r.Nombre = strings.ToUpper(string(nombre[0])) + nombre[1:]
	return r
}
