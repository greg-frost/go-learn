package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"go-learn/avito/internal/model"
	"go-learn/avito/internal/storage"
	u "go-learn/avito/internal/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Структура "обработчик"
type handler struct {
	storage storage.Storage
}

// Конструктор обработчика
func NewHandler(storage storage.Storage) Handler {
	return &handler{storage: storage}
}

// Регистрация эндпоинтов в обработчике
func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/dummyLogin", h.DummyLogin).Methods("POST")
	router.HandleFunc("/pvz", h.CreatePVZ).Methods("POST")
	router.HandleFunc("/receptions", h.CreateReception).Methods("POST")
	router.HandleFunc("/pvz/{pvzId}/close_last_reception", h.CloseLastReception).Methods("POST")
	router.HandleFunc("/products", h.CreateProduct).Methods("POST")
	router.HandleFunc("/pvz/{pvzId}/delete_last_product", h.DeleteLastProduct).Methods("POST")
	router.HandleFunc("/", h.GetSummary).Methods("GET")
}

// Псевдо-логин (получение токена)
func (h *handler) DummyLogin(w http.ResponseWriter, r *http.Request) {
	var user model.UserDTO
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		u.RespondWithError(w, model.Error{
			Code:    http.StatusBadRequest,
			Message: "неправильный запрос",
		})
		return
	}

	role := user.Role
	if !role.Valid() {
		u.RespondWithError(w, model.Error{
			Code:    http.StatusBadRequest,
			Message: "неправильная роль",
		})
		return
	}

	token := model.Token{Role: role}
	tokenString, err := token.SignedString()
	if err != nil {
		u.RespondWithError(w, model.Error{
			Code:    http.StatusInternalServerError,
			Message: "токен не был создан",
		})
		return
	}

	u.Respond(w, http.StatusOK, tokenString)
}

// Создание ПВЗ
func (h *handler) CreatePVZ(w http.ResponseWriter, r *http.Request) {
	role := u.GetRoleFromContext(r.Context())
	if role != model.RoleModerator {
		u.RespondWithError(w, model.Error{
			Code:    http.StatusForbidden,
			Message: "только для модераторов",
		})
		return
	}

	var pvz model.PvzDTO
	if err := json.NewDecoder(r.Body).Decode(&pvz); err != nil {
		u.RespondWithError(w, model.Error{
			Code:    http.StatusBadRequest,
			Message: "неправильный запрос",
		})
		return
	}

	city := pvz.City
	if !city.Valid() {
		u.RespondWithError(w, model.Error{
			Code:    http.StatusBadRequest,
			Message: "неправильный город",
		})
		return
	}

	newPVZ := model.PVZ{
		ID:               uuid.NewString(),
		RegistrationDate: time.Now(),
		City:             city,
	}

	if err := h.storage.CreatePVZ(newPVZ); err != nil {
		log.Println("ошибка создания ПВЗ:", err)
		u.RespondWithError(w, model.Error{
			Code:    http.StatusInternalServerError,
			Message: "ПВЗ не был создан",
		})
		return
	}

	u.Respond(w, http.StatusCreated, newPVZ)
}

// Создание приемки
func (h *handler) CreateReception(w http.ResponseWriter, r *http.Request) {
	role := u.GetRoleFromContext(r.Context())
	if role != model.RoleEmployee {
		u.RespondWithError(w, model.Error{
			Code:    http.StatusForbidden,
			Message: "только для сотрудников",
		})
		return
	}

	var reception model.ReceptionDTO
	if err := json.NewDecoder(r.Body).Decode(&reception); err != nil {
		u.RespondWithError(w, model.Error{
			Code:    http.StatusBadRequest,
			Message: "неправильный запрос",
		})
		return
	}

	pvzID := reception.PvzID

	if _, err := h.storage.FindPVZ(pvzID); err != nil {
		u.RespondWithError(w, model.Error{
			Code:    http.StatusBadRequest,
			Message: "ПВЗ не найден",
		})
		return
	}

	lastReception, _ := h.storage.FindLastReception(pvzID)
	if lastReception.Status == model.StatusInProgress {
		u.RespondWithError(w, model.Error{
			Code:    http.StatusBadRequest,
			Message: "другая приемка активна",
		})
		return
	}

	newReception := model.Reception{
		ID:       uuid.NewString(),
		DateTime: time.Now(),
		PvzID:    pvzID,
		Status:   model.StatusInProgress,
	}

	if err := h.storage.CreateReception(newReception); err != nil {
		log.Println("ошибка создания приемки:", err)
		u.RespondWithError(w, model.Error{
			Code:    http.StatusInternalServerError,
			Message: "приемка не была создана",
		})
		return
	}

	u.Respond(w, http.StatusCreated, newReception)
}

// Закрытие последней (активной) приемки
func (h *handler) CloseLastReception(w http.ResponseWriter, r *http.Request) {
	role := u.GetRoleFromContext(r.Context())
	if role != model.RoleEmployee {
		u.RespondWithError(w, model.Error{
			Code:    http.StatusForbidden,
			Message: "только для сотрудников",
		})
		return
	}

	pvzID := mux.Vars(r)["pvzId"]

	lastReception, _ := h.storage.FindLastReception(pvzID)
	if lastReception.Status != model.StatusInProgress {
		u.RespondWithError(w, model.Error{
			Code:    http.StatusBadRequest,
			Message: "нет активной приемки",
		})
		return
	}

	if err := h.storage.CloseReception(lastReception.ID); err != nil {
		log.Println("ошибка обновления приемки:", err)
		u.RespondWithError(w, model.Error{
			Code:    http.StatusInternalServerError,
			Message: "приемка не была закрыта",
		})
		return
	}
	lastReception.Status = model.StatusClose

	u.Respond(w, http.StatusOK, lastReception)
}

// Добавление товара
func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	role := u.GetRoleFromContext(r.Context())
	if role != model.RoleEmployee {
		u.RespondWithError(w, model.Error{
			Code:    http.StatusForbidden,
			Message: "только для сотрудников",
		})
		return
	}

	var product model.ProductDTO
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		u.RespondWithError(w, model.Error{
			Code:    http.StatusBadRequest,
			Message: "неправильный запрос",
		})
		return
	}

	productType := product.Type
	if !productType.Valid() {
		u.RespondWithError(w, model.Error{
			Code:    http.StatusBadRequest,
			Message: "неправильный тип",
		})
		return
	}

	pvzID := product.PvzID

	lastReception, _ := h.storage.FindLastReception(pvzID)
	if lastReception.Status != model.StatusInProgress {
		u.RespondWithError(w, model.Error{
			Code:    http.StatusBadRequest,
			Message: "нет активной приемки",
		})
		return
	}

	newProduct := model.Product{
		ID:          uuid.NewString(),
		DateTime:    time.Now(),
		Type:        productType,
		ReceptionID: lastReception.ID,
	}

	if err := h.storage.CreateProduct(newProduct); err != nil {
		log.Println("ошибка добавления товара:", err)
		u.RespondWithError(w, model.Error{
			Code:    http.StatusInternalServerError,
			Message: "товар не был добавлен",
		})
		return
	}

	u.Respond(w, http.StatusCreated, newProduct)
}

// Удаление последнего добавленного товара
func (h *handler) DeleteLastProduct(w http.ResponseWriter, r *http.Request) {
	role := u.GetRoleFromContext(r.Context())
	if role != model.RoleEmployee {
		u.RespondWithError(w, model.Error{
			Code:    http.StatusForbidden,
			Message: "только для сотрудников",
		})
		return
	}

	pvzID := mux.Vars(r)["pvzId"]

	lastReception, _ := h.storage.FindLastReception(pvzID)
	if lastReception.Status != model.StatusInProgress {
		u.RespondWithError(w, model.Error{
			Code:    http.StatusBadRequest,
			Message: "нет активной приемки",
		})
		return
	}

	lastProduct, err := h.storage.FindLastProduct(lastReception.ID)
	if err != nil {
		u.RespondWithError(w, model.Error{
			Code:    http.StatusBadRequest,
			Message: "нет товаров для удаления",
		})
		return
	}

	if err := h.storage.DeleteProduct(lastProduct.ID); err != nil {
		log.Println("ошибка удаления товара:", err)
		u.RespondWithError(w, model.Error{
			Code:    http.StatusInternalServerError,
			Message: "товар не был удален",
		})
		return
	}

	u.Respond(w, http.StatusOK, nil)
}

const (
	summaryPageDefault  = 1  // Значение номера страницы по умолчанию
	summaryPageMin      = 1  // Минимальное значение номера страницы
	summaryLimitDefault = 10 // Ограничение выборки по умолчанию
	summaryLimitMin     = 1  // Минимальное значение ограничения выборки
	summaryLimitMax     = 30 // Максимальное значение ограничения выборки
)

// Получение сводной информации
func (h *handler) GetSummary(w http.ResponseWriter, r *http.Request) {
	role := u.GetRoleFromContext(r.Context())
	if role != model.RoleEmployee && role != model.RoleModerator {
		u.RespondWithError(w, model.Error{
			Code:    http.StatusForbidden,
			Message: "только для сотрудников и модераторов",
		})
		return
	}

	pageInput := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageInput)
	if err != nil {
		page = summaryPageDefault
	}
	if page < summaryPageMin {
		page = summaryPageMin
	}

	limitInput := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitInput)
	if err != nil {
		limit = summaryLimitDefault
	}
	if limit < summaryLimitMin {
		limit = summaryLimitMin
	}
	if limit > summaryLimitMax {
		limit = summaryLimitMax
	}

	var filterByDate bool
	startDateInput := r.URL.Query().Get("startDate")
	startDate, err := time.Parse(time.RFC3339, startDateInput)
	if err == nil {
		filterByDate = true
	}
	endDateInput := r.URL.Query().Get("endDate")
	endDate, err := time.Parse(time.RFC3339, endDateInput)
	if err != nil {
		endDate = time.Now()
	} else {
		filterByDate = true
	}

	pvzs, err := h.storage.ListPVZ(page, limit, startDate, endDate, filterByDate)
	if err != nil {
		log.Println("ошибка выдачи суммарной информации по ПВЗ:", err)
		u.RespondWithError(w, model.Error{
			Code:    http.StatusInternalServerError,
			Message: "внутренняя ошибка сервера",
		})
		return
	}
	pvzIDs := make([]string, 0, len(pvzs))
	for _, pvz := range pvzs {
		pvzIDs = append(pvzIDs, pvz.ID)
	}

	receptions, err := h.storage.ListReceptions(pvzIDs, startDate, endDate)
	if err != nil {
		log.Println("ошибка выдачи суммарной информации по приемкам:", err)
		u.RespondWithError(w, model.Error{
			Code:    http.StatusInternalServerError,
			Message: "внутренняя ошибка сервера",
		})
		return
	}
	var receptionsIDs []string
	for _, rcps := range receptions {
		for _, reception := range rcps {
			receptionsIDs = append(receptionsIDs, reception.ID)
		}
	}

	products, err := h.storage.ListProducts(receptionsIDs)
	if err != nil {
		log.Println("ошибка выдачи суммарной информации по товарам:", err)
		u.RespondWithError(w, model.Error{
			Code:    http.StatusInternalServerError,
			Message: "внутренняя ошибка сервера",
		})
		return
	}

	results := make([]model.PVZResult, 0, len(pvzs))
	for _, pvz := range pvzs {
		pvzResult := model.PVZResult{
			PVZ:        pvz,
			Receptions: make([]model.ReceptionResult, 0, len(receptions[pvz.ID])),
		}
		for _, reception := range receptions[pvz.ID] {
			receptionResult := model.ReceptionResult{
				Reception: reception,
				Products:  make([]model.Product, 0),
			}
			if receptionProducts, ok := products[reception.ID]; ok {
				receptionResult.Products = receptionProducts
			}
			pvzResult.Receptions = append(pvzResult.Receptions, receptionResult)
		}
		results = append(results, pvzResult)
	}

	u.Respond(w, http.StatusOK, results)
}
