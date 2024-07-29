package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	pb "golearn/protobuf/v2/user"

	"google.golang.org/protobuf/proto"
)

// Структура "пользователь"
type User struct {
	Name  string `json:"name"`
	Id    int32  `json:"id"`
	Email string `json:"email,omitempty"`
}

// Новый пользователь
func newUser() User {
	return User{
		Name:  "Greg Frost",
		Id:    100021,
		Email: "greg-frost@yandex.ru",
	}
}

// Новый protobuf-пользователь
func newPbUser() pb.User {
	return pb.User{
		Name:  proto.String("Greg Frost"),
		Id:    proto.Int32(100021),
		Email: proto.String("greg-frost@yandex.ru"),
	}
}

// Обработчик JSON
func handleJSON(w http.ResponseWriter, r *http.Request) {
	u := newUser()

	body, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

// Обработчик Protocol Buffers
func handleProtobuf(w http.ResponseWriter, r *http.Request) {
	u := newPbUser()

	body, err := proto.Marshal(&u)
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
		http.HandleFunc("/json", handleJSON)
		http.HandleFunc("/protobuf", handleProtobuf)

		fmt.Println("Ожидаю обновлений...")
		fmt.Println("(на http://localhost:8080)")
		log.Fatal(http.ListenAndServe("localhost:8080", nil))
	}()

	time.Sleep(250 * time.Millisecond)
	fmt.Println()

	/* JSON */

	fmt.Println("JSON:")
	res, err := http.Get("http://localhost:8080/json")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	var jsUser User
	err = json.Unmarshal(body, &jsUser)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Имя:", jsUser.Name)
	fmt.Println("ID:", jsUser.Id)
	fmt.Println("E-mail:", jsUser.Email)
	raw := string(body)
	fmt.Printf("RAW: %s (%d)\n\n", raw, len(raw))

	/* Protobuf */

	fmt.Println("Protobuf:")
	res, err = http.Get("http://localhost:8080/protobuf")
	if err != nil {
		log.Fatal(err)
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	var pbUser pb.User
	err = proto.Unmarshal(body, &pbUser)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Имя:", pbUser.GetName())
	fmt.Println("ID:", pbUser.GetId())
	fmt.Println("E-mail:", pbUser.GetEmail())
	raw = strings.TrimSpace(string(body))
	fmt.Printf("RAW: %s (%d)\n\n", raw, len(raw))

	/* Сравнение */

	times := 100000
	fmt.Printf("Сравнение (%d повторов):\n", times)

	// JSON
	jsUser = newUser()
	start := time.Now()
	for i := 0; i < times; i++ {
		body, _ = json.Marshal(jsUser)
		json.Unmarshal(body, &jsUser)
	}
	fmt.Println("JSON:", time.Now().Sub(start))

	// Protobuf
	pbUser = newPbUser()
	start = time.Now()
	for i := 0; i < times; i++ {
		body, _ = proto.Marshal(&pbUser)
		proto.Unmarshal(body, &pbUser)
	}
	fmt.Println("Protobuf:", time.Now().Sub(start))
}
