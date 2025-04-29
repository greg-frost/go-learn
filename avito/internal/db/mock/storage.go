package mock

import (
	"errors"
	"time"

	"go-learn/avito/internal/model"
	s "go-learn/avito/internal/storage"

	"github.com/google/uuid"
)

// ПВЗ
var PVZ = model.PVZ{
	ID:               uuid.NewString(),
	RegistrationDate: time.Now(),
	City:             model.CityMoscow,
}

// ПВЗ с активной приемкой
var PvzInProgress = model.PVZ{
	ID:               uuid.NewString(),
	RegistrationDate: time.Now(),
	City:             model.CitySaintPetersburg,
}

// ПВЗ с закрытой приемкой
var PvzClosed = model.PVZ{
	ID:               uuid.NewString(),
	RegistrationDate: time.Now(),
	City:             model.CityKazan,
}

// ПВЗ с товарами
var PvzWithProducts = model.PVZ{
	ID:               uuid.NewString(),
	RegistrationDate: time.Now(),
	City:             model.CityMoscow,
}

// Активная приемка
var ReceptionInProgress = model.Reception{
	ID:       uuid.NewString(),
	DateTime: time.Now(),
	PvzID:    PvzInProgress.ID,
	Status:   model.StatusInProgress,
}

// Закрытая приемка
var ReceptionClosed = model.Reception{
	ID:       uuid.NewString(),
	DateTime: time.Now(),
	PvzID:    PvzClosed.ID,
	Status:   model.StatusClose,
}

// Приемка с товарами
var ReceptionWithProducts = model.Reception{
	ID:       uuid.NewString(),
	DateTime: time.Now(),
	PvzID:    PvzInProgress.ID,
	Status:   model.StatusInProgress,
}

// Товар
var Product = model.Product{
	ID:          uuid.NewString(),
	DateTime:    time.Now(),
	ReceptionID: ReceptionWithProducts.ID,
	Type:        model.TypeElectronics,
}

// Массив ПВЗ
var PVZs = []model.PVZ{
	PVZ,
	PvzInProgress,
	PvzClosed,
	PvzWithProducts,
	{
		ID:               uuid.NewString(),
		RegistrationDate: time.Now(),
		City:             model.CitySaintPetersburg,
	},
}

// Массив приемок
var Receptions = []model.Reception{
	ReceptionInProgress,
	ReceptionClosed,
	ReceptionWithProducts,
}

// Массив товаров
var Products = []model.Product{
	Product,
	Product,
	{
		ID:          uuid.NewString(),
		DateTime:    time.Now(),
		ReceptionID: ReceptionInProgress.ID,
		Type:        model.TypeElectronics,
	},
	{
		ID:          uuid.NewString(),
		DateTime:    time.Now(),
		ReceptionID: ReceptionClosed.ID,
		Type:        model.TypeClothes,
	},
	{
		ID:          uuid.NewString(),
		DateTime:    time.Now(),
		ReceptionID: ReceptionWithProducts.ID,
		Type:        model.TypeShoes,
	},
}

// Структура "хранилище"
type storage struct{}

// Конструктор хранилища
func NewStorage() s.Storage {
	return &storage{}
}

// Создание ПВЗ
func (s *storage) CreatePVZ(pvz model.PVZ) error {
	return nil
}

// Поиск ПВЗ
func (s *storage) FindPVZ(pvzID string) (model.PVZ, error) {
	switch pvzID {
	case PVZ.ID:
		return PVZ, nil
	case PvzInProgress.ID:
		return PvzInProgress, nil
	case PvzClosed.ID:
		return PvzClosed, nil
	case PvzWithProducts.ID:
		return PvzWithProducts, nil
	default:
		return model.PVZ{}, errors.New("ПВЗ не найден")
	}
}

// Список ПВЗ
func (s *storage) ListPVZ(page, limit int, startDate, endDate time.Time, filterByDate bool) (
	[]model.PVZ, error) {
	results := make([]model.PVZ, 0, limit)
	for _, pvz := range PVZs {
		if filterByDate {
			var receptionsCount int
			for _, reception := range Receptions {
				if pvz.ID != reception.PvzID ||
					reception.DateTime.Before(startDate) ||
					reception.DateTime.After(endDate) {
					continue
				}
				receptionsCount++
			}
			if receptionsCount == 0 {
				continue
			}
		}
		results = append(results, pvz)
	}

	offset := (page - 1) * limit
	if offset > len(results) {
		offset = len(results)
	}
	results = results[offset:]
	if len(results) > limit {
		results = results[:limit]
	}

	return results, nil
}

// Удаление ПВЗ
func (s *storage) DeletePVZ(pvzID string) error {
	for _, pvz := range PVZs {
		if pvzID == pvz.ID {
			return nil
		}
	}
	return errors.New("нет ПВЗ для удаления")
}

// Создание приемки
func (s *storage) CreateReception(reception model.Reception) error {
	return nil
}

// Поиск последней приемки
func (s *storage) FindLastReception(pvzID string) (model.Reception, error) {
	switch pvzID {
	case PVZ.ID:
		return model.Reception{}, nil
	case PvzInProgress.ID:
		return ReceptionInProgress, nil
	case PvzClosed.ID:
		return ReceptionClosed, nil
	case PvzWithProducts.ID:
		return ReceptionWithProducts, nil
	default:
		return model.Reception{}, errors.New("ПВЗ не найден")
	}
}

// Список приемок
func (s *storage) ListReceptions(pvzIDs []string, startDate, endDate time.Time) (
	map[string][]model.Reception, error) {
	results := make(map[string][]model.Reception, len(pvzIDs))
	for _, pvzID := range pvzIDs {
		for _, reception := range Receptions {
			if pvzID != reception.PvzID ||
				reception.DateTime.Before(startDate) ||
				reception.DateTime.After(endDate) {
				continue
			}
			results[pvzID] = append(results[pvzID], reception)
		}
	}
	return results, nil
}

// Закрытие приемки
func (s *storage) CloseReception(receptionID string) error {
	if receptionID != ReceptionInProgress.ID &&
		receptionID != ReceptionWithProducts.ID {
		return errors.New("нет активной приемки")
	}
	return nil
}

// Удаление приемки
func (s *storage) DeleteReception(receptionID string) error {
	for _, reception := range Receptions {
		if receptionID == reception.ID {
			return nil
		}
	}
	return errors.New("нет приемки для удаления")
}

// Создание товара
func (s *storage) CreateProduct(product model.Product) error {
	return nil
}

// Поиск последнего товара
func (s *storage) FindLastProduct(receptionID string) (model.Product, error) {
	if receptionID != ReceptionInProgress.ID &&
		receptionID != ReceptionWithProducts.ID {
		return model.Product{}, errors.New("нет активной приемки")
	}
	if receptionID != ReceptionWithProducts.ID {
		return model.Product{}, errors.New("нет товаров для удаления")
	}
	return Product, nil
}

// Список товаров
func (s *storage) ListProducts(receptionsIDs []string) (
	map[string][]model.Product, error) {
	results := make(map[string][]model.Product, len(receptionsIDs))
	for _, receptionID := range receptionsIDs {
		for _, product := range Products {
			if receptionID == product.ReceptionID {
				results[receptionID] = append(results[receptionID], product)
			}
		}
	}

	return results, nil
}

// Удаление товара
func (s *storage) DeleteProduct(productID string) error {
	for _, product := range Products {
		if productID == product.ID {
			return nil
		}
	}
	return errors.New("нет товара для удаления")
}
