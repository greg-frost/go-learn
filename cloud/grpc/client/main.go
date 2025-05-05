package client

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "go-learn/cloud/grpc/proto"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println(" \n[ GRPC-КЛИЕНТ ]\n ")

	// Установка соединения
	conn, err := grpc.Dial(
		"localhost:5050",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Получение клиента
	client := pb.NewCloudClient(conn)

	// Чтение параметров
	var action, key, value string
	if len(os.Args) > 2 {
		action, key = os.Args[1], os.Args[2]
		value = strings.Join(os.Args[3:], " ")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Выполнение действия
	switch action {
	case "get":
		r, err := client.Get(ctx, &pb.GetRequest{Key: key})
		if err != nil {
			log.Fatalf("не удалось получить значение для ключа %s: %v", key, err)
		}
		log.Printf("Get %s: %s", key, r.Value)
	case "put":
		_, err := client.Put(ctx, &pb.PutRequest{Key: key, Value: value})
		if err != nil {
			log.Fatalf("не удалось сохранить ключ %s: %v", key, err)
		}
		log.Printf("Put %s", key)
	default:
		log.Fatal("пример: go run client.go [get|put] KEY VALUE")
	}
}
