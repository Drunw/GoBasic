package main

import (
	"strings"
)

var idRuta string = "ruta convertir mayuscula"

func convertirMayuscula(r Request) Request {
	mensaje := "Comienza la ruta convertir mayuscula"
	imprimirLogSimple(uuidString, idRuta, mensaje)
	nombre := r.Nombre
	apellido := r.Apellido
	r.Apellido = strings.ToUpper(string(apellido[0])) + apellido[1:]
	r.Nombre = strings.ToUpper(string(nombre[0])) + nombre[1:]
	imprimirLogSalidaApi(uuidString, idRuta, r.String())
	return r
}
