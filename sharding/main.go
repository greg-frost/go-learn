package main

import (
	"crypto/sha1"
	"fmt"
	"sync"
)

// Структура "сегмент"
type Shard struct {
	sync.RWMutex
	m map[string]interface{}
}

// Набор сегментов
type Shards []*Shard

// Конструктор набора сегментов
func NewShards(n int) Shards {
	shards := make(Shards, n)
	for i := 0; i < n; i++ {
		shard := make(map[string]interface{})
		shards[i] = &Shard{m: shard}
	}

	return shards
}

// Получение индекса
func (s Shards) getIndex(key string) int {
	checksum := sha1.Sum([]byte(key))
	hash := int(checksum[16])

	return hash % len(s)
}

// Получение сегмента
func (s Shards) getShard(key string) *Shard {
	index := s.getIndex(key)
	return s[index]
}

// Получение значения
func (s Shards) Get(key string) interface{} {
	shard := s.getShard(key)
	shard.RLock()
	defer shard.RUnlock()

	return shard.m[key]
}

// Установка значения
func (s Shards) Set(key string, value interface{}) {
	shard := s.getShard(key)
	shard.Lock()
	defer shard.Unlock()

	shard.m[key] = value
}

// Удаление значения
func (s Shards) Delete(key string) {
	shard := s.getShard(key)
	shard.Lock()
	defer shard.Unlock()

	delete(shard.m, key)
}

// Проверка значения
func (s Shards) Contains(key string) bool {
	shard := s.getShard(key)
	shard.RLock()
	defer shard.RUnlock()

	_, ok := shard.m[key]
	return ok
}

// Список всех ключей
func (s Shards) Keys() []string {
	keys := make([]string, 0)
	var m sync.Mutex
	var wg sync.WaitGroup

	wg.Add(len(s))
	for _, shard := range s {
		go func(sh *Shard) {
			sh.RLock()

			for key := range sh.m {
				m.Lock()
				keys = append(keys, key)
				m.Unlock()
			}

			sh.RUnlock()
			wg.Done()
		}(shard)
	}
	wg.Wait()

	return keys
}

func main() {
	fmt.Println(" \n[ СЕГМЕНТИРОВАНИЕ ]\n ")

	// Настройка
	shards := NewShards(5)
	shards.Set("alpha", 1)
	shards.Set("beta", 2)
	shards.Set("gamma", 3)
	shards.Set("wrong", 100)
	shards.Delete("wrong")

	// Работа
	keys := shards.Keys()
	for _, key := range keys {
		fmt.Printf("%s = %d\n", key, shards.Get(key))
	}
	fmt.Println()
	fmt.Println("Is alpha?", shards.Contains("alpha"))
	fmt.Println("Is omega?", shards.Contains("omega"))
}
