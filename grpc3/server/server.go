package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net"
	"os"
	"sync"
	"time"

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
	routeNotes    map[string][]*pb.RouteNote
	mu            sync.Mutex
}

// Конструктор сервера
func NewRouteServer() *routeServer {
	rs := &routeServer{
		routeNotes: make(map[string][]*pb.RouteNote),
	}

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

// Запись маршрута по заданным точкам и возвращение суммарной информации
func (s *routeServer) RecordRoute(stream pb.Route_RecordRouteServer) error {
	var pointCount, featureCount, distance int32
	var prev *pb.Point
	startTime := time.Now()

	for {
		point, err := stream.Recv()
		if err == io.EOF {
			endTime := time.Now()
			return stream.SendAndClose(&pb.RouteSummary{
				PointCount:   pointCount,
				FeatureCount: featureCount,
				Distance:     distance,
				ElapsedTime:  int32(endTime.Sub(startTime).Seconds()),
			})
		}
		if err != nil {
			return err
		}

		pointCount++
		for _, feature := range s.savedFeatures {
			if proto.Equal(feature.Location, point) {
				featureCount++
			}
		}
		if prev != nil {
			distance += calcDistance(prev, point)
		}
		prev = point
	}
}

// Обмен сообщениями между клиентом и сервером по точкам маршрута
func (s *routeServer) RouteChat(stream pb.Route_RouteChatServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		key := serialize(in.Location)

		s.mu.Lock()
		s.routeNotes[key] = append(s.routeNotes[key], in)
		rn := make([]*pb.RouteNote, len(s.routeNotes[key]))
		copy(rn, s.routeNotes[key])
		s.mu.Unlock()

		for _, note := range rn {
			if err := stream.Send(note); err != nil {
				return err
			}
		}
	}
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

// Расчет расстояния между двумя точками
func calcDistance(p1 *pb.Point, p2 *pb.Point) int32 {
	const CordFactor float64 = 1e7
	const R = float64(6371000) // радиус Земли в метрах

	lat1 := toRadians(float64(p1.Latitude) / CordFactor)
	lat2 := toRadians(float64(p2.Latitude) / CordFactor)
	lng1 := toRadians(float64(p1.Longitude) / CordFactor)
	lng2 := toRadians(float64(p2.Longitude) / CordFactor)
	dlat := lat2 - lat1
	dlng := lng2 - lng1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c
	return int32(distance)
}

// Перевод градусов в радианы
func toRadians(num float64) float64 {
	return num * math.Pi / float64(180)
}

// Сериализация координат
func serialize(point *pb.Point) string {
	return fmt.Sprintf("%d %d", point.Latitude, point.Longitude)
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
