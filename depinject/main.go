package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

// Функция логирования
func LogLog(message string) {
	log.Println(message)
}

// Другая функция логирования
func LogPrint(message string) {
	fmt.Println(message)
}

// Интерфейс "хранилище данных"
type DataStore interface {
	NameById(userID string) (string, bool)
}

// Фабрика хранилища
func NewDataStore(store string) DataStore {
	switch store {
	case "simple":
		return NewSimpleDataStore()
	case "complex":
		return NewComplexDataStore()
	default:
		return NewSimpleDataStore()
	}
}

// Структура "простое хранилище данных"
type SimpleDataStore struct {
	userData map[string]string
}

// Конструктор простого хранилища
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		map[string]string{
			"1": "Greg",
			"2": "John",
			"3": "Ada",
		},
	}
}

// Получение имени по идентификатору из простого хранилища
func (sds SimpleDataStore) NameById(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

// Структура "сложное хранилище данных"
type ComplexDataStore struct {
	userData map[int]string
}

// Конструктор сложного хранилища
func NewComplexDataStore() ComplexDataStore {
	return ComplexDataStore{
		map[int]string{
			2: "Greg Frost",
			4: "John Smith",
			8: "Ada Wong",
		},
	}
}

// Получение имени по идентификатору из сложного хранилища
func (cds ComplexDataStore) NameById(userID string) (string, bool) {
	userIDNum, _ := strconv.Atoi(userID)
	name, ok := cds.userData[int(math.Pow(2, float64(userIDNum)))]
	return name, ok
}

// Интерфейс "логгер"
type Logger interface {
	Log(message string)
}

// Адаптер логгера
type LoggerAdapter func(message string)

// Логирование через функцию адаптера
func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

// Интерфейс "логика"
type Logic interface {
	SayHello(userID string) (string, error)
	SayGoodbye(userID string) (string, error)
}

// Фабрика логики
func NewLogic(logic string, lg Logger, ds DataStore) Logic {
	switch logic {
	case "simple":
		return NewSimpleLogic(lg, ds)
	case "complex":
		return NewComplexLogic(lg, ds)
	default:
		return NewSimpleLogic(lg, ds)
	}
}

// Структура "простая логика"
type SimpleLogic struct {
	lg Logger
	ds DataStore
}

// Конструктор простой логики
func NewSimpleLogic(lg Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		lg: lg,
		ds: ds,
	}
}

// Приветствие по простой логике
func (sl SimpleLogic) SayHello(userID string) (string, error) {
	name, ok := sl.ds.NameById(userID)
	if !ok {
		return "", errors.New("Неизвестный пользователь")
	}

	sl.lg.Log("В функции SayHello для пользователя " + userID)
	return "Привет, " + name, nil
}

// Прощание по простой логике
func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	name, ok := sl.ds.NameById(userID)
	if !ok {
		return "", errors.New("Неизвестный пользователь")
	}

	sl.lg.Log("В функции SayGoodbye для пользователя " + userID)
	return "Пока, " + name, nil
}

// Структура "сложная логика"
type ComplexLogic struct {
	lg Logger
	ds DataStore
}

// Конструктор сложной логики
func NewComplexLogic(lg Logger, ds DataStore) ComplexLogic {
	return ComplexLogic{
		lg: lg,
		ds: ds,
	}
}

// Приветствие по сложной логике
func (cl ComplexLogic) SayHello(userID string) (string, error) {
	name, ok := cl.ds.NameById(userID)
	if !ok {
		return "", errors.New("Unknown user")
	}

	cl.lg.Log("In function SayHello for user " + userID)
	return "Hello then, " + name + ", nice to meet you!", nil
}

// Прощание по сложной логике
func (cl ComplexLogic) SayGoodbye(userID string) (string, error) {
	name, ok := cl.ds.NameById(userID)
	if !ok {
		return "", errors.New("Unknown user")
	}

	cl.lg.Log("In function SayGoodbye for user " + userID)
	return "Bye, " + name + ", sorry that you leave...", nil
}

// Структура "контроллер"
type Controller struct {
	lg Logger
	lc Logic
}

// Конструктор контроллера
func NewController(lg Logger, lc Logic) Controller {
	return Controller{
		lg: lg,
		lc: lc,
	}
}

// Приветствие через контроллер
func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.lg.Log("В контроллере SayHello")
	userID := r.URL.Query().Get("user_id")
	message, err := c.lc.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

// Прощание через контроллер
func (c Controller) SayGoodbye(w http.ResponseWriter, r *http.Request) {
	c.lg.Log("В контроллере SayGoodbye")
	userID := r.URL.Query().Get("user_id")
	message, err := c.lc.SayGoodbye(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func main() {
	fmt.Println(" \n[ ВНЕДРЕНИЕ ЗАВИСИМОСТИ ]\n ")

	// Простое
	simpleLogger := LoggerAdapter(LogPrint)
	simpleDataStore := NewDataStore("simple")
	simpleLogic := NewLogic("simple", simpleLogger, simpleDataStore)
	simpleController := NewController(simpleLogger, simpleLogic)

	// Сложное
	complexLogger := LoggerAdapter(LogLog)
	complexDataStore := NewDataStore("complex")
	complexLogic := NewLogic("complex", complexLogger, complexDataStore)
	complexController := NewController(complexLogger, complexLogic)

	// Обработчики
	http.HandleFunc("/hi", simpleController.SayHello)
	http.HandleFunc("/bye", simpleController.SayGoodbye)
	http.HandleFunc("/hello", complexController.SayHello)
	http.HandleFunc("/goodbye", complexController.SayGoodbye)

	// Запуск сервера
	fmt.Println("Ожидаю соединений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
