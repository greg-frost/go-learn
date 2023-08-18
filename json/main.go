package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

// Структура "человек"
type Person struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"`
}

// Композитная структура "Ответ"
type (
	Response struct {
		Header Header `json:"header"`
		Data   []Data `json:"data,omitempty"`
	}

	Header struct {
		Code    int    `json:"code"`
		Message string `json:"message,omitempty"`
	}

	Data struct {
		Type       string     `json:"type"`
		ID         int        `json:"id"`
		Attributes Attributes `json:"attributes"`
	}

	Attributes struct {
		Email       string `json:"email"`
		Article_ids []int  `json:"article_ids"`
	}
)

// Чтение JSON-строки
func readJson(respJson string) (Response, error) {
	resp := Response{}
	if err := json.Unmarshal([]byte(respJson), &resp); err != nil {
		return Response{}, fmt.Errorf("JSON десериализация: %w", err)
	}

	return resp, nil
}

func main() {
	fmt.Println(" \n[ JSON ]\n ")

	/* Сериализация */

	man := Person{
		Name:        "Greg",
		Email:       "greg-frost@yandex.ru",
		DateOfBirth: time.Now(),
	}

	jsMan, err := json.Marshal(man)
	if err != nil {
		log.Fatalln("не удалось сериализовать в json")
	}

	fmt.Printf("Сериализация:\n%v", string(jsMan))
	fmt.Println(" \n ")

	/* Анонимная структура */

	req := struct {
		NameContains string `json:"name_contains"`
		Offset       int    `json:"offset"`
		Limit        int    `json:"limit"`
	}{
		NameContains: "Григорий",
		Limit:        50,
	}

	jsReq, _ := json.Marshal(req)
	fmt.Printf("Анонимная структура:\n%v", string(jsReq))
	fmt.Println(" \n ")

	/* Композитный JSON (вручную) */

	resp1 := Response{
		Header: Header{
			Code:    0,
			Message: "",
		},
		Data: []Data{
			{
				Type: "user",
				ID:   100,
				Attributes: Attributes{
					Email:       "greg-frost@yandex.ru",
					Article_ids: []int{5, 10, 100},
				},
			},
		},
	}

	jsResp1, _ := json.Marshal(resp1)
	fmt.Printf("Структура ответа (вручную):\n%+v\n%v\n", resp1, string(jsResp1))
	fmt.Println()

	/* Композитный JSON (десериализация) */

	const respJson = `
	{
	    "header": {
	        "code": 200,
	        "message": "OK"
	    },
	    "data": [{
	        "type": "admin",
	        "id": 525,
	        "attributes": {
	            "email": "bob@yandex.ru",
	            "article_ids": [10, 11, 12]
	        }
	    }]
	}
	`

	resp2, _ := readJson(respJson)
	jsResp2, _ := json.Marshal(resp2)
	fmt.Printf("Структура ответа (десериализация):\n%+v\n%v\n", resp2, string(jsResp2))

	/* Encoder и Decoder */

	const data = `
		{"name": "Fred", "age": 40}
		{"name": "Mary", "age": 21}
		{"name": "Pat", "age": 30}
	`
	var p struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var b bytes.Buffer
	dec := json.NewDecoder(strings.NewReader(data))
	enc := json.NewEncoder(&b)

	for dec.More() {
		err := dec.Decode(&p)
		if err != nil {
			panic(err)
		}

		p.Name = p.Name + " " + p.Name

		err = enc.Encode(p)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println()
	fmt.Println("Encode и Decode:")
	fmt.Print(b.String())
}
