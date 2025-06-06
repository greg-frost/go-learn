package memory

import "sync"

// Структура "клиент"
type Client struct {
	m  map[string]interface{}
	mu sync.RWMutex
}

// Конструктор клиента
func NewClient() *Client {
	return &Client{
		m: make(map[string]interface{}),
	}
}

// Получение
func (c *Client) Retrieve(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, ok := c.m[key]
	return value, ok
}

// Сохранение
func (c *Client) Store(key string, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.m[key] = value
	return true
}

// Удаление
func (c *Client) Remove(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.m, key)
	return true
}

// Все ключи
func (c *Client) Keys() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var keys []string
	for key := range c.m {
		keys = append(keys, key)
	}

	return keys
}

// Все значения
func (c *Client) Values() []interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var values []interface{}
	for _, value := range c.m {
		values = append(values, value)
	}

	return values
}
