package main

import (
	"strings"

	"github.com/sirupsen/logrus"
)

func imprimirLogSimple(idTransaccion string, idRuta string, mensaje string) {
	logger := logrus.New()
	idTransaccion = strings.ToUpper(idTransaccion)
	idRuta = strings.ToUpper(idRuta)
	logger.Infof(`{route: %s, traceId: %s, mensaje: %s}`, idRuta, idTransaccion, mensaje)
}

func imprimirLogEntradaApi(idTransaccion string, body string) {
	logger := logrus.New()
	idTransaccion = strings.ToUpper(idTransaccion)
	idRuta := "RUTA INICIAL"
	mensaje := "Peticion de entrada al API"
	logger.Infof(`{route: %s, traceId: %s, mensaje: %s, entrada:{%s} }`, idRuta, idTransaccion, mensaje, body)

}

func imprimirLogSalidaApi(idTransaccion string, idRuta string, body string) {
	logger := logrus.New()
	idTransaccion = strings.ToUpper(idTransaccion)
	idRuta = strings.ToUpper(idRuta)
	mensaje := "Body de salida del API"
	logger.Infof(`{route: %s, traceId: %s, mensaje: %s, salida:{%s} }`, idRuta, idTransaccion, mensaje, body)

}
