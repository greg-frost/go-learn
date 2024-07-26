package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	pb "golearn/protobuf/v2/user"

	"google.golang.org/protobuf/proto"
)

// Обработчик protobuf
func handleProtobuf(w http.ResponseWriter, r *http.Request) {
	u := &pb.User{
		Name:  proto.String("Greg Frost"),
		Id:    proto.Int32(100021),
		Email: proto.String("greg-frost@yandex.ru"),
	}

	body, err := proto.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/x-protobuf")
	w.Write(body)
}

func main() {
	fmt.Println(" \n[ PROTOBUF ]\n ")

	/* Сервер */

	fmt.Println("Сервер:")
	go func() {
		http.HandleFunc("/proto", handleProtobuf)

		fmt.Println("Ожидаю обновлений...")
		fmt.Println("(на http://localhost:8080)")
		log.Fatal(http.ListenAndServe("localhost:8080", nil))
	}()

	time.Sleep(250 * time.Millisecond)
	fmt.Println()

	/* Клиент */

	fmt.Println("Клиент:")
	res, err := http.Get("http://localhost:8080/proto")
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	var u pb.User
	err = proto.Unmarshal(b, &u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Имя:", u.GetName())
	fmt.Println("ID:", u.GetId())
	fmt.Println("E-mail:", u.GetEmail())
	fmt.Println()

	fmt.Println("RAW:", strings.TrimSpace(string(b)))
}
