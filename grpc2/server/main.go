package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "golearn/grpc2/protos/hello"

	"google.golang.org/grpc"
)

// Порт
var port = flag.Int("port", 8888, "Порт сервера")

// Структура "сервер"
type server struct {
	pb.UnimplementedGreeterServer
}

// Приветствие
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Получено: %v", in.GetName())
	return &pb.HelloReply{Message: "Привет, " + in.GetName() + "!"}, nil
}

// Поздравление
func (s *server) Congrats(ctx context.Context, in *pb.CongratRequest) (*pb.HelloReply, error) {
	log.Printf("Получено: имя - %v, возраст - %d", in.GetName(), in.GetAge())
	msg := fmt.Sprintf("%v, поздравляем с %d-летием!", in.GetName(), in.GetAge())
	return &pb.HelloReply{Message: msg}, nil
}

func main() {
	fmt.Println(" \n[ GRPC2 (СЕРВЕР) ]\n ")

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

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
