# Usar la imagen oficial de Go como base
FROM golang:alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /go/src/app

# Copiar los archivos necesarios
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Dar permisos de ejecución al archivo main
RUN chmod +x main

# Segunda etapa para reducir el tamaño de la imagen
FROM alpine:latest

# Instalar dependencias
RUN apk --no-cache add ca-certificates

# Establecer el directorio de trabajo
WORKDIR /root/

# Copiar el ejecutable desde la etapa de compilación
COPY --from=builder /go/src/app/main .

# Exponer el puerto en el que se ejecuta la aplicación
EXPOSE 8080

# Ejecutar la aplicación
CMD ["./main"]
