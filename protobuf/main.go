package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"go-learn/base"
	pb2 "go-learn/protobuf/v2/user"
	pbp "go-learn/protobuf/v3/people"
	pb3 "go-learn/protobuf/v3/user"

	"google.golang.org/protobuf/proto"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
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

// Новый пользователь Protobuf v2
func newUserPb2() pb2.User {
	return pb2.User{
		Name:  proto.String("Greg Frost"),
		Id:    proto.Int32(100021),
		Email: proto.String("greg-frost@yandex.ru"),
	}
}

// Новый пользователь Protobuf v3
func newUserPb3() pb3.User {
	return pb3.User{
		Name:  "Greg Frost",
		Id:    100021,
		Email: "greg-frost@yandex.ru",
	}
}

// Новый человек
func newPerson() *pbp.Person {
	return &pbp.Person{
		Name:  "Greg Frost",
		Id:    100021,
		Email: "greg-frost@yandex.ru",
		Phones: []*pbp.Person_PhoneNumber{
			{
				Number: "+7 987 65-43-21",
				Type:   pbp.PhoneType_PHONE_TYPE_HOME,
			},
		},
		LastUpdated: tspb.New(time.Now()),
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

// Обработчик Protocol Buffers v2
func handleProtobuf2(w http.ResponseWriter, r *http.Request) {
	u := newUserPb2()

	body, err := proto.Marshal(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/x-protobuf")
	w.Write(body)
}

// Обработчик Protocol Buffers v3
func handleProtobuf3(w http.ResponseWriter, r *http.Request) {
	u := newUserPb3()

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
		http.HandleFunc("/protobuf2", handleProtobuf2)
		http.HandleFunc("/protobuf3", handleProtobuf3)

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

	var jsonUser User
	err = json.Unmarshal(body, &jsonUser)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Имя:", jsonUser.Name)
	fmt.Println("ID:", jsonUser.Id)
	fmt.Println("E-mail:", jsonUser.Email)
	raw := string(body)
	fmt.Printf("RAW: %s (%d)\n\n", raw, len(raw))

	/* Protobuf v2 */

	fmt.Println("Protobuf v2:")
	res, err = http.Get("http://localhost:8080/protobuf2")
	if err != nil {
		log.Fatal(err)
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	var pbv2User pb2.User
	err = proto.Unmarshal(body, &pbv2User)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Имя:", pbv2User.GetName())
	fmt.Println("ID:", pbv2User.GetId())
	fmt.Println("E-mail:", pbv2User.GetEmail())
	raw = strings.TrimSpace(string(body))
	fmt.Printf("RAW: %s (%d)\n\n", raw, len(raw))

	/* Protobuf v3 */

	fmt.Println("Protobuf v3:")
	res, err = http.Get("http://localhost:8080/protobuf3")
	if err != nil {
		log.Fatal(err)
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	var pbv3User pb3.User
	err = proto.Unmarshal(body, &pbv3User)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Имя:", pbv3User.GetName())
	fmt.Println("ID:", pbv3User.GetId())
	fmt.Println("E-mail:", pbv3User.GetEmail())
	raw = strings.TrimSpace(string(body))
	fmt.Printf("RAW: %s (%d)\n\n", raw, len(raw))

	/* Сравнение */

	times := 100000
	fmt.Printf("Сравнение:\n(%d повторов)\n\n", times)

	// JSON
	jsonUser = newUser()
	start := time.Now()
	for i := 0; i < times; i++ {
		body, _ = json.Marshal(jsonUser)
		json.Unmarshal(body, &jsonUser)
	}
	fmt.Println("JSON:", time.Now().Sub(start))

	// Protobuf v2
	pbv2User = newUserPb2()
	start = time.Now()
	for i := 0; i < times; i++ {
		body, _ = proto.Marshal(&pbv2User)
		proto.Unmarshal(body, &pbv2User)
	}
	fmt.Println("Protobuf v2:", time.Now().Sub(start))

	// Protobuf v3
	pbv3User = newUserPb3()
	start = time.Now()
	for i := 0; i < times; i++ {
		body, _ = proto.Marshal(&pbv3User)
		proto.Unmarshal(body, &pbv3User)
	}
	fmt.Println("Protobuf v3:", time.Now().Sub(start))
	fmt.Println()

	/* Сложный пример */

	fmt.Println("Комплексный Protobuf:")
	fmt.Println()

	fmt.Println("Создание объектов...")
	people := &pbp.People{
		People: []*pbp.Person{newPerson(), newPerson(), newPerson()},
	}

	fmt.Println("Сериализация...")
	out, err := proto.Marshal(people)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Запись в файл...")
	path := base.Dir("protobuf")
	filename := filepath.Join(path, "file.txt")
	if err = ioutil.WriteFile(filename, out, 0644); err != nil {
		log.Fatal(err)
	}
	defer os.Remove(filename)

	fmt.Println("Чтение из файла...")
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Десериализация...")
	newPeople := &pbp.People{}
	if err = proto.Unmarshal(in, newPeople); err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	fmt.Print("Сериализованные данные идентичны: ")
	if equalBytes := bytes.Compare(in, out); equalBytes == 0 {
		fmt.Println("да")
	} else {
		fmt.Println("нет")
	}

	fmt.Print("Десериализованные структуры равны: ")
	if p1, p2 := fmt.Sprintf("%+v", people), fmt.Sprintf("%+v", newPeople); p1 == p2 {
		fmt.Println("да")
	} else {
		fmt.Println("нет")
	}

	fmt.Print("Десериализованные структуры идентичны: ")
	if reflect.DeepEqual(people, newPeople) {
		fmt.Println("да")
	} else {
		fmt.Println("нет")
	}
	fmt.Println()

	fmt.Println("Имя:", newPeople.People[0].GetName())
	fmt.Println("ID:", newPeople.People[0].GetId())
	fmt.Println("E-mail:", newPeople.People[0].GetEmail())
	fmt.Print("Телефоны: ")
	for _, phone := range newPeople.People[0].GetPhones() {
		var phoneType string
		switch phone.Type {
		case pbp.PhoneType_PHONE_TYPE_MOBILE:
			phoneType = "сот."
		case pbp.PhoneType_PHONE_TYPE_HOME:
			phoneType = "дом."
		case pbp.PhoneType_PHONE_TYPE_WORK:
			phoneType = "раб."
		default:
			phoneType = "?"
		}
		fmt.Printf("%s (%s) ", phone.Number, phoneType)
	}
	fmt.Println()
	fmt.Println("Дата:", people.People[0].GetLastUpdated().AsTime())
	fmt.Println()

	fmt.Println("Количество:", len(newPeople.People))
}
