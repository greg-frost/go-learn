package main

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
	var value string
	var err error

	return &pb.GetResponse{Value: value}, err
}

// Сохранение значения
func (s *server) Put(ctx context.Context, r *pb.PutRequest) (*pb.PutResponse, error) {
	log.Printf("Получен запрос PUT key=%v value=%v", r.Key, r.Value)

	// err := Put(r.Key, r.Value)
	var err error

	return &pb.PutResponse{}, err
}

// Удаление значения
func (s *server) Delete(ctx context.Context, r *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	log.Printf("Получен запрос DELETE key=%v", r.Key)

	// err := Delete(r.Key)
	var err error

	return &pb.DeleteResponse{}, err
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
