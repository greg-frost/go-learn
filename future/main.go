package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Интерфейс "в будущем"
type Future interface {
	Result() (string, error)
}

// Структура "в будущем"
type InnerFuture struct {
	once sync.Once
	wg   sync.WaitGroup

	res   string
	err   error
	chRes <-chan string
	chErr <-chan error
}

// Результат "в будущем"
func (p *InnerFuture) Result() (string, error) {
	p.once.Do(func() {
		p.wg.Add(1)
		defer p.wg.Done()

		p.res = <-p.chRes
		p.err = <-p.chErr
	})

	p.wg.Wait()

	return p.res, p.err
}

// Вызов функции "в будущем" (промис)
func Promise(ctx context.Context, delay time.Duration) Future {
	chRes := make(chan string)
	chErr := make(chan error)

	go func() {
		select {
		case <-time.After(delay):
			chRes <- "OK"
			chErr <- nil
		case <-ctx.Done():
			chRes <- ""
			chErr <- ctx.Err()
		}
	}()

	return &InnerFuture{chRes: chRes, chErr: chErr}
}

func main() {
	fmt.Println(" \n[ В БУДУЩЕМ ]\n ")

	// Настройка
	ctx := context.Background()
	future := Promise(ctx, 3*time.Second)

	// Работа
	time.Sleep(time.Second)
	fmt.Println("Ожидание...")
	time.Sleep(time.Second)
	res, err := future.Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
