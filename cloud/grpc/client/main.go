package main

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
			log.Fatalf("не удалось получить значение ключа %s: %v", key, err)
		}
		log.Printf("Get %s: %q", key, r.Value)
	case "put":
		_, err := client.Put(ctx, &pb.PutRequest{Key: key, Value: value})
		if err != nil {
			log.Fatalf("не удалось сохранить значение %q ключа %s: %v", value, key, err)
		}
		log.Printf("Put %s: %q", key, value)
	case "delete":
		_, err := client.Delete(ctx, &pb.DeleteRequest{Key: key})
		if err != nil {
			log.Fatalf("не удалось удалить значение ключа %s: %v", key, err)
		}
		log.Printf("Delete %s", key)
	default:
		fmt.Println("Синтаксис: go run ... [get|put|delete] KEY (VALUE)")
		return
	}
}
