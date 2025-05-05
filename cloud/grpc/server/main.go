package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "go-learn/cloud/grpc/proto"

	"google.golang.org/grpc"
)

// Структура "сервер"
type server struct {
	pb.UnimplementedCloudServer
}

// Получение значения
func (s *server) Get(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {
	log.Printf("Получен запрос GET key=%v", r.Key)

	// value, err := Get(r.Key)
	value, err := "", fmt.Errorf("не реализовано")

	return &pb.GetResponse{Value: value}, err
}

func main() {
	fmt.Println(" \n[ GRPC-СЕРВЕР ]\n ")

	// Новый gRPC сервер
	s := grpc.NewServer()

	// Регистрация Cloud-сервера
	pb.RegisterCloudServer(s, &server{})

	// Прослушивание TCP-порта
	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatal(err)
	}

	// Запуск сервера
	fmt.Println("Ожидаю TCP-подключений...")
	fmt.Println("(на localhost:5050)")
	log.Fatal(s.Serve(listener))
}
