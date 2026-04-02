package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type (
	// Карта "поля структуры"
	FieldsInfo map[string]FieldInfo

	// Структура "информация о поле структуры"
	FieldInfo struct {
		Type     string     `json:"type"`
		Tags     TagsInfo   `json:"tags,omitempty"`
		Embedded FieldsInfo `json:"embedded,omitempty"`
	}

	// Карта "информация о тегах"
	TagsInfo map[string][]string
)

// Строковое представление карты полей структуры
func (f FieldsInfo) String() string {
	b, _ := json.MarshalIndent(f, "", "   ")
	return string(b)
}

// Информация о поле структуры
func GetStructTags(obj interface{}) (info FieldsInfo) {
	info = make(FieldsInfo)

	// Уточнение типа
	var objType reflect.Type
	if t, ok := obj.(reflect.Type); ok {
		objType = t
	} else {
		objType = reflect.ValueOf(obj).Type()
	}

	// Проверка на указатель
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	// Проверка на структуру
	if objType.Kind() != reflect.Struct {
		return
	}

	// Обход полей
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		info[field.Name] = FieldInfo{
			Type:     field.Type.String(),               // Тип
			Tags:     parseTagString(string(field.Tag)), // Теги
			Embedded: GetStructTags(field.Type),         // Подструктура
		}
	}
	return
}

// Разбор строки тегов
func parseTagString(tagRaw string) (info TagsInfo) {
	info = make(TagsInfo)

	// Пример: json:"name" pg:"nullable,sortable"
	for _, tag := range strings.Split(tagRaw, " ") {
		if tag = strings.TrimSpace(tag); tag == "" {
			continue
		}

		// Разделение тега на части
		tagParts := strings.Split(tag, ":")
		if len(tagParts) != 2 {
			continue
		}

		// Имя тега
		tagName := strings.TrimSpace(tagParts[0])
		if _, ok := info[tagName]; ok {
			continue
		}

		// Значения тега
		tagValuesRaw, _ := strconv.Unquote(tagParts[1])
		var tagValues []string
		for _, value := range strings.Split(tagValuesRaw, ",") {
			if value := strings.TrimSpace(value); value != "" {
				tagValues = append(tagValues, value)
			}
		}

		info[tagName] = tagValues
	}
	return
}

// Тестовая структура
type (
	TestStruct struct {
		ID        string `json:"id" format:"uuid" example:"68b69bd2-8db6-4b7f-b7f0-7c78739046c6"`
		Name      string `json:"name" example:"Bob"`
		Group     Group  `json:"group"`
		CreatedAt int64  `json:"created_at" format:"unix" example:"1622647813"`
	}
	Group struct {
		ID             uint64   `json:"id"`
		PermsOverrides []string `json:"overrides" example:"USERS_RW,COMPANY_RWC"`
	}
)

func main() {
	fmt.Println(" \n[ РЕФЛЕКСИЯ ПО ТЕГАМ ]\n ")

	var s *TestStruct
	fmt.Println(GetStructTags(s))
}
