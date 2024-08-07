package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "golearn/grpc2/protos/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:8888", "Адрес сервера")
	name = flag.String("name", "Greg Frost", "Имя для приветствия")
	age  = flag.Int("age", 100, "Возраст для поздравления")
)

func main() {
	fmt.Println(" \n[ GRPC2 (КЛИЕНТ) ]\n ")

	// Парсинг флагов
	flag.Parse()

	// Соединение с сервером
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Создание клиента
	c := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Вызов gRPC-метода SayHello
	r1, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Метод SayHello: %q\n", r1.GetMessage())

	// Вызов gRPC-метода Congrats
	r2, err := c.Congrats(ctx, &pb.CongratRequest{Name: *name, Age: int32(*age)})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Метод Congrats: %q\n", r2.GetMessage())
}
