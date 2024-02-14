# Usar la imagen oficial de Go como base
FROM golang:1.21.6-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download
RUN go get myapp

COPY *.go ./

RUN go build -o main .

EXPOSE 8080

CMD [ "./main" ]
