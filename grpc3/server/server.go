package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net"
	"os"

	pb "golearn/grpc3/protos/route"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

// Порт и путь
var port = flag.Int("port", 8888, "Порт сервера")
var path = os.Getenv("GOPATH") + "/src/golearn/grpc3/"

// Структура "сервер"
type routeServer struct {
	pb.UnimplementedRouteServer
	savedFeatures []*pb.Feature
}

// Конструктор сервера
func NewRouteServer() *routeServer {
	rs := &routeServer{}

	// Чтение файла
	b, err := ioutil.ReadFile(path + "data/routes.json")
	if err != nil {
		return rs
	}

	// Демаршаллинг
	err = json.Unmarshal(b, &rs.savedFeatures)
	if err != nil {
		return rs
	}

	return rs
}

// Получение объекта по координатам
func (s *routeServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range s.savedFeatures {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}
	return &pb.Feature{Location: point}, nil
}

// Получение объектов в области, заданной прямоугольником координат
func (s *routeServer) ListFeatures(rect *pb.Rectangle, stream pb.Route_ListFeaturesServer) error {
	for _, feature := range s.savedFeatures {
		if inRange(feature.Location, rect) {
			if err := stream.Send(feature); err != nil {
				return err
			}
		}
	}
	return nil
}

// Проверка вхождения объекта в прямоугольник координат
func inRange(point *pb.Point, rect *pb.Rectangle) bool {
	left := math.Min(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	right := math.Max(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	top := math.Max(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))
	bottom := math.Min(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))

	if float64(point.Longitude) >= left &&
		float64(point.Longitude) <= right &&
		float64(point.Latitude) >= bottom &&
		float64(point.Latitude) <= top {
		return true
	}
	return false
}

func main() {
	fmt.Println(" \n[ GRPC3 (СЕРВЕР) ]\n ")

	// Парсинг флагов
	flag.Parse()

	// Запуск сервера
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Запуск на", listener.Addr())
	fmt.Println("(ожидаю обновлений...)")
	fmt.Println()

	rs := NewRouteServer()
	s := grpc.NewServer()
	pb.RegisterRouteServer(s, rs)
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
