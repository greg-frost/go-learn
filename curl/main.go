package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println(" \n[ CURL ]\n ")

	/* Получение HTML-страницы */

	resp, _ := http.Get("http://gregfrostmusic.ru/")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()
}
