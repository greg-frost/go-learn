package main

import (
	"fmt"
)

// Интерфейс "сервис"
type Service interface {
	Get(key string) string
	Put(key, value string)
	Remove(key string)
}

// Структура "удаленная служба"
type RemoteService struct {
	m map[string]string
}

// Конструктор удаленной службы
func NewRemoteService() Service {
	return &RemoteService{
		m: make(map[string]string),
	}
}

// Получение значения
func (rs *RemoteService) Get(key string) string {
	return rs.m[key]
}

// Сохранение значения
func (rs *RemoteService) Put(key, value string) {
	rs.m[key] = value
}

// Удаление значения
func (rs *RemoteService) Remove(key string) {
	delete(rs.m, key)
}

// Структура "удаленный заместитель"
type RemoteProxy struct {
	service Service
}

// Конструктор удаленного заместителя
func NewRemoteProxy() Service {
	return &RemoteProxy{
		service: NewRemoteService(),
	}
}

// Получение значения
func (rp *RemoteProxy) Get(key string) string {
	return rp.service.Get(key)
}

// Сохранение значения
func (rp *RemoteProxy) Put(key, value string) {
	rp.service.Put(key, value)
}

// Удаление значения
func (rp *RemoteProxy) Remove(key string) {
	rp.service.Remove(key)
}

// Структура "заместитель доступа"
type AccessProxy struct {
	service Service
}

// Конструктор заместителя доступа
func NewAccessProxy(service Service) Service {
	return &AccessProxy{
		service: service,
	}
}

// Получение значения
func (ap *AccessProxy) Get(key string) string {
	return ap.service.Get(key)
}

// Сохранение значения
func (*AccessProxy) Put(key, value string) {
	fmt.Println("Put: доступ запрещен")
}

// Удаление значения
func (*AccessProxy) Remove(key string) {
	fmt.Println("Remove: доступ запрещен")
}

func main() {
	fmt.Println(" \n[ ЗАМЕСТИТЕЛЬ ]\n ")

	fmt.Println("Удаленный заместитель:")
	remote := NewRemoteProxy()
	remote.Put("key", "remote value")
	fmt.Printf("key = %q\n", remote.Get("key"))
	remote.Put("deleted", "value")
	remote.Remove("deleted")
	fmt.Printf("deleted = %q\n", remote.Get("deleted"))
	fmt.Println()

	fmt.Println("Заместитель доступа:")
	access := NewAccessProxy(remote)
	access.Put("newkey", "value")
	access.Remove("newkey")
	fmt.Printf("key = %q\n", access.Get("key"))
}
