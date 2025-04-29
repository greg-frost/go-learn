//go:build integration

package integration

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"go-learn/avito/internal/auth"
	"go-learn/avito/internal/db/postgres"
	"go-learn/avito/internal/handler"
	"go-learn/avito/internal/model"
	"go-learn/avito/internal/storage"

	"github.com/gorilla/mux"
)

var (
	r *mux.Router
	h handler.Handler
	s storage.Storage
)

func init() {
	os.Setenv("DB_NAME", "avito")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASS", "admin")

	r = mux.NewRouter()

	pgParams := postgres.ConnectionParams{
		DbName:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}
	var err error
	s, err = postgres.NewStorage(pgParams)
	if err != nil {
		log.Fatal(err)
	}

	h = handler.NewHandler(s)
	h.Register(r)

	r.Use(auth.JwtAuthentication)
}

func TestIntegration(t *testing.T) {
	t.Log("Начато")

	moderatorToken, err := getToken("moderator")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Токен модератора получен:", "..."+
		moderatorToken[len(moderatorToken)-16:])

	pvzID, err := createPVZ(moderatorToken, randomCity())
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ПВЗ создан:", pvzID)

	defer func() {
		err := s.DeletePVZ(pvzID)
		if err != nil {
			t.Fatalf("ПВЗ и информация не были удалены (ID: %s)", pvzID)
		}

		t.Log("Товары удалены")
		t.Log("Приемка удалена")
		t.Log("ПВЗ удален")

		t.Log("Закончено")
	}()

	employeeToken, err := getToken("employee")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Токен сотрудника получен:", "..."+
		employeeToken[len(employeeToken)-16:])

	receptionID, err := createReception(employeeToken, pvzID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Приемка создана:", receptionID)

	productsCount := 50
	for i := 0; i < productsCount; i++ {
		_, err := createProduct(employeeToken, randomType(), pvzID)
		if err != nil {
			t.Fatal(err)
		}
	}
	t.Log("Число добавленных товаров:", productsCount)

	productID, err := createProduct(employeeToken, randomType(), pvzID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Еще один товар добавлен:", productID)

	err = deleteLastProduct(employeeToken, pvzID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Последний добавленный товар удален")

	err = closeLastReception(employeeToken, pvzID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Приемка закрыта")
}

// Получение токена
func getToken(role string) (string, error) {
	body := `{"role": "` + role + `"}`
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"POST", "/dummyLogin",
		strings.NewReader(body))

	r.ServeHTTP(rec, req)
	if rec.Code != 200 {
		return "", fmt.Errorf("не получен токен для роли %s: %w", role, getError(rec))
	}

	var token string
	json.NewDecoder(rec.Body).Decode(&token)
	return token, nil
}

// Создание ПВЗ
func createPVZ(token, city string) (string, error) {
	body := `{"city": "` + city + `"}`
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"POST", "/pvz",
		strings.NewReader(body))

	req.Header.Add("Authorization", "Bearer "+token)

	r.ServeHTTP(rec, req)
	if rec.Code != 201 {
		return "", fmt.Errorf("не создан ПВЗ для города %s: %w", city, getError(rec))
	}

	var pvz model.PVZ
	json.NewDecoder(rec.Body).Decode(&pvz)
	return pvz.ID, nil
}

// Создание приемки
func createReception(token, pvzID string) (string, error) {
	body := `{"pvzId": "` + pvzID + `"}`
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"POST", "/receptions",
		strings.NewReader(body))

	req.Header.Add("Authorization", "Bearer "+token)

	r.ServeHTTP(rec, req)
	if rec.Code != 201 {
		return "", fmt.Errorf("не создана приемка для ПВЗ %s: %w", pvzID, getError(rec))
	}

	var reception model.Reception
	json.NewDecoder(rec.Body).Decode(&reception)
	return reception.ID, nil
}

// Добавление товара
func createProduct(token, productType, pvzID string) (string, error) {
	body := `{"type": "` + productType + `", "pvzId": "` + pvzID + `"}`
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"POST", "/products",
		strings.NewReader(body))

	req.Header.Add("Authorization", "Bearer "+token)

	r.ServeHTTP(rec, req)
	if rec.Code != 201 {
		return "", fmt.Errorf("не добавлен товар для ПВЗ %s: %w", pvzID, getError(rec))
	}

	var product model.Product
	json.NewDecoder(rec.Body).Decode(&product)
	return product.ID, nil
}

// Удаление последнего товара
func deleteLastProduct(token, pvzID string) error {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"POST", "/pvz/"+pvzID+"/delete_last_product", nil)

	req.Header.Add("Authorization", "Bearer "+token)

	r.ServeHTTP(rec, req)
	if rec.Code != 200 {
		return fmt.Errorf("не удален товар для ПВЗ %s: %w", pvzID, getError(rec))
	}

	return nil
}

// Закрытие активной приемки
func closeLastReception(token, pvzID string) error {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"POST", "/pvz/"+pvzID+"/close_last_reception", nil)

	req.Header.Add("Authorization", "Bearer "+token)

	r.ServeHTTP(rec, req)
	if rec.Code != 200 {
		return fmt.Errorf("не закрыта приемка для ПВЗ %s: %w", pvzID, getError(rec))
	}

	return nil
}

// Получение ошибки
func getError(rec *httptest.ResponseRecorder) model.Error {
	var err model.Error
	json.NewDecoder(rec.Body).Decode(&err)
	err.Code = rec.Code
	return err
}

// Случайный город
func randomCity() string {
	cities := []model.City{
		model.CityMoscow,
		model.CitySaintPetersburg,
		model.CityKazan,
	}
	i := rand.Intn(len(cities))
	return string(cities[i])
}

// Случайный тип товара
func randomType() string {
	types := []model.Type{
		model.TypeElectronics,
		model.TypeClothes,
		model.TypeShoes,
	}
	i := rand.Intn(len(types))
	return string(types[i])
}
