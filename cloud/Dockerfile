# Этап 1: Компиляция

FROM golang:1.16 as build

COPY . /src

WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux go build -o cloud

# Этап 2: Сборка

FROM scratch

COPY --from=build /src/cloud .
COPY --from=build /src/*.pem .

EXPOSE 8080

CMD ["/cloud"]