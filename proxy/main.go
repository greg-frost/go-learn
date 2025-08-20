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

func main() {
	fmt.Println(" \n[ ЗАМЕСТИТЕЛЬ ]\n ")
}
