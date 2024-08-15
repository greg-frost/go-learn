package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	pb "golearn/grpc3/protos/route"

	"google.golang.org/grpc"
)

// Адрес сервера
var addr = flag.String("addr", "localhost:8888", "Адрес сервера")

// Поиск объекта по координатам
func GetFeature(client pb.RouteClient, point *pb.Point) {
	fmt.Printf("Получение объекта для координат (%d, %d):\n", point.Latitude, point.Longitude)

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

// Поиск объектов в области координат
func ListFeatures(client pb.RouteClient, rect *pb.Rectangle) {
	fmt.Printf("Поиск объектов в области координат (%d, %d) - (%d, %d):\n",
		rect.Lo.Latitude, rect.Lo.Longitude, rect.Hi.Latitude, rect.Hi.Longitude)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.ListFeatures(ctx, rect)
	if err != nil {
		log.Fatal(err)
	}

	var count int
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		name := feature.GetName()
		if name == "" {
			name = "?"
		}
		fmt.Printf("Найдено: %q (%v, %v)\n", name,
			feature.GetLocation().GetLatitude(), feature.GetLocation().GetLongitude())

		count++
	}
	fmt.Println("Найдено объектов:", count)
}

// Запись маршрута и выдача результата
func RecordRoute(client pb.RouteClient) {
	// Слаучайное число (но не меньше двух) случайных точек
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pointCount := int(r.Int31n(100)) + 2
	var points []*pb.Point
	for i := 0; i < pointCount; i++ {
		points = append(points, randomPoint(r))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.RecordRoute(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, point := range points {
		if err := stream.Send(point); err != nil {
			log.Fatal(err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Результат маршрута:\nТочек - %d\nОбъектов - %d\nРасстояние - %d\n",
		reply.GetPointCount(), reply.GetFeatureCount(), reply.GetDistance())
}

// Генерация случайной точки
func randomPoint(r *rand.Rand) *pb.Point {
	lat := (r.Int31n(180) - 90) * 1e7
	long := (r.Int31n(360) - 180) * 1e7
	return &pb.Point{Latitude: lat, Longitude: long}
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

	// Метод GetFeature
	fmt.Printf("GetFeature\n----------\n")
	GetFeature(c, &pb.Point{Latitude: 409146138, Longitude: -746188906})
	GetFeature(c, &pb.Point{Latitude: 0, Longitude: 0})
	fmt.Println()

	// Метод ListFeatures
	fmt.Printf("ListFeatures\n------------\n")
	ListFeatures(c, &pb.Rectangle{
		Lo: &pb.Point{Latitude: 400000000, Longitude: -750000000},
		Hi: &pb.Point{Latitude: 405000000, Longitude: -745000000},
	})
	fmt.Println()

	// Метод RecordRoute
	fmt.Printf("RecordRoute\n-----------\n")
	RecordRoute(c)
}
