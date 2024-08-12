package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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
	routes := []struct {
		Name     string
		Location struct {
			Latitude  int32
			Longitude int32
		}
	}{}
	err = json.Unmarshal(b, &routes)
	if err != nil {
		return rs
	}

	// Запись объектов
	var sf []*pb.Feature
	for _, r := range routes {
		sf = append(sf, &pb.Feature{
			Name: r.Name,
			Location: &pb.Point{
				Latitude:  r.Location.Latitude,
				Longitude: r.Location.Longitude,
			},
		})
	}
	rs.savedFeatures = sf

	return rs
}

// Приветствие
func (s *routeServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range s.savedFeatures {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}
	return &pb.Feature{Location: point}, nil
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
