package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "golearn/grpc3/protos/route"

	"google.golang.org/grpc"
)

// Адрес сервера
var addr = flag.String("addr", "localhost:8888", "Адрес сервера")

// Получение объекта
func GetFeature(client pb.RouteClient, point *pb.Point) {
	fmt.Printf("Получение объекта для координат (%d, %d)\n", point.Latitude, point.Longitude)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	feature, err := client.GetFeature(ctx, point)
	if err != nil {
		log.Fatal(err)
	}

	name := feature.GetName()
	if name != "" {
		fmt.Println("Найдено:", feature.GetName())
	} else {
		fmt.Println("Не найдено...")
	}
}

func main() {
	fmt.Println(" \n[ GRPC3 (КЛИЕНТ) ]\n ")

	// Парсинг флагов
	flag.Parse()

	// Соединение с сервером
	conn, err := grpc.NewClient(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Создание клиента
	c := pb.NewRouteClient(conn)

	// Вызов gRPC-метода GetFeature

	fmt.Println("GetFeature:")
	GetFeature(c, &pb.Point{Latitude: 409146138, Longitude: -746188906})
	GetFeature(c, &pb.Point{Latitude: 0, Longitude: 0})
}
