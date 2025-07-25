package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "go-learn/cloud/frontend/grpc/proto"

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

	// Новый клиент
	client := pb.NewCloudClient(conn)

	// Параметры и настройки
	var action, key, value string
	if len(os.Args) > 2 {
		action, key = os.Args[1], os.Args[2]
		value = strings.Join(os.Args[3:], " ")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch action {
	// Получение значения
	case "get":
		r, err := client.Get(ctx, &pb.GetRequest{Key: key})
		if err != nil {
			log.Fatalf("не удалось получить значение ключа %s: %v", key, err)
		}
		log.Printf("GET %s: %q", key, r.Value)

	// Запись значения
	case "put":
		_, err := client.Put(ctx, &pb.PutRequest{Key: key, Value: value})
		if err != nil {
			log.Fatalf("не удалось сохранить значение %q ключа %s: %v", value, key, err)
		}
		log.Printf("PUT %s: %q", key, value)

	// Удаление значения
	case "delete":
		_, err := client.Delete(ctx, &pb.DeleteRequest{Key: key})
		if err != nil {
			log.Fatalf("не удалось удалить значение ключа %s: %v", key, err)
		}
		log.Printf("DELETE %s", key)

	default:
		fmt.Println("Синтаксис: go run ... [get|put|delete] KEY (VALUE)")
		return
	}
}
