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

// Композитная структура "ответ"
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

// Чтение JSON
func readJSON(respJson string) (Response, error) {
	var resp Response
	if err := json.Unmarshal([]byte(respJson), &resp); err != nil {
		return Response{}, fmt.Errorf("JSON десериализация: %w", err)
	}

	return resp, nil
}

// Печать JSON
func printJSON(v interface{}) {
	switch v := v.(type) {
	case string:
		fmt.Println("строка:", v)
	case float64:
		fmt.Println("число:", v)
	case bool:
		fmt.Println("число:", v)
	case []interface{}:
		fmt.Println("массив:")
		for _, data := range v {
			fmt.Print("- ")
			printJSON(data)
		}
	case map[string]interface{}:
		fmt.Println("объект:")
		for field, data := range v {
			fmt.Printf("%q ", field)
			printJSON(data)
		}
	default:
		fmt.Println("неизвестный тип")
	}
}

func main() {
	fmt.Println(" \n[ JSON ]\n ")

	/* Сериализация */

	man := Person{
		Name:        "Greg",
		Email:       "greg-frost@yandex.ru",
		DateOfBirth: time.Now(),
	}

	jsonMan, err := json.Marshal(man)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Сериализация:\n%v\n\n", string(jsonMan))

	/* Анонимная структура */

	req := struct {
		NameContains string `json:"name_contains"`
		Offset       int    `json:"offset"`
		Limit        int    `json:"limit"`
	}{
		NameContains: "Григорий",
		Limit:        50,
	}

	jsonReq, _ := json.Marshal(req)
	fmt.Printf("Анонимная структура:\n%v\n\n", string(jsonReq))

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

	jsonResp1, _ := json.Marshal(resp1)
	fmt.Printf("Структура ответа (вручную):\n%+v\n%v\n\n", resp1, string(jsonResp1))

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

	resp2, _ := readJSON(respJson)
	jsonResp2, _ := json.Marshal(resp2)
	fmt.Printf("Структура ответа (десериализация):\n%+v\n%v\n\n", resp2, string(jsonResp2))

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
			log.Fatal(err)
		}

		p.Name = p.Name + " " + p.Name

		err = enc.Encode(p)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Encode и Decode:")
	fmt.Println(b.String())

	/* Произвольный формат */

	fmt.Println("Произвольный формат:")
	jsonCustom := `
	{
		"first_name": "Greg",
		"last_name": "Frost",
		"age": 36,
		"is_active": true,
		"langs":
		{
			"items": [
				"Go",
				"PHP",
				"JavaScript"
			],
			"rate": 1.21
		}
	}
	`

	var custom interface{}
	err = json.Unmarshal([]byte(jsonCustom), &custom)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(custom)
	fmt.Println()

	printJSON(custom)
}
