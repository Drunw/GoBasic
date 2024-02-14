# Usar la imagen oficial de Go como base
FROM golang:1.21.6-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-docker-demo

EXPOSE 8080

CMD [ "/go-docker-demo" ]
