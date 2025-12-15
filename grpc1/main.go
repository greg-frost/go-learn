package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "go-learn/grpc1/protos/hello"

	"google.golang.org/grpc"
)

// Структура "сервер"
type server struct{}

// Общение через gRPC
func (s *server) Say(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	msg := "Привет, " + in.Name + "!"
	return &pb.HelloResponse{Message: msg}, nil
}

func main() {
	fmt.Println(" \n[ GRPC 1 ]\n ")

	// Сервер
	fmt.Println("Сервер:")
	go func() {
		fmt.Println("Запуск на localhost:8888")

		listener, err := net.Listen("tcp", "localhost:8888")
		if err != nil {
			log.Fatal(err)
		}

		s := grpc.NewServer()
		pb.RegisterHelloServer(s, &server{})
		s.Serve(listener)
	}()
	time.Sleep(250 * time.Millisecond)
	fmt.Println()

	// Клиент
	fmt.Println("Клиент:")
	addr := "localhost:8888"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)
	name := "Greg Frost"
	hr := &pb.HelloRequest{Name: name}
	r, err := c.Say(context.Background(), hr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Получено %q\n", r.Message)
}
