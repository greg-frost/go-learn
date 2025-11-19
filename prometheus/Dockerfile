# Этап 1: Компиляция

FROM golang:alpine as build

RUN apk update && apk add --no-cache git

RUN mkdir $GOPATH/src/server
ADD ./main.go $GOPATH/src/server
WORKDIR $GOPATH/src/server

RUN go mod init
RUN go mod tidy
RUN go mod download

RUN mkdir /prometheus
RUN go build -o /prometheus/server main.go

# Этап 2: Сборка

FROM alpine:latest

RUN mkdir /prometheus
COPY --from=build /prometheus/server /prometheus/server

EXPOSE 8080

WORKDIR /prometheus

CMD ["/prometheus/server"]