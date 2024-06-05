package main

import (
	"fmt"
	"net"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"
)

func main() {
	fmt.Println(" \n[ СКАНИРОВАНИЕ ПОРТОВ ]\n ")

	// Фоновый запуск сервера (открытие порта)
	go http.ListenAndServe("localhost:8080", nil)
	time.Sleep(100 * time.Millisecond)

	// Проверка одного порта
	open := scanPort("tcp", "localhost", 8080)
	fmt.Println("Порт 8080 открыт:", open)
	fmt.Println()

	/* Сканироване портов 1-1024 */

	fmt.Println("Идет сканирование портов...")
	fmt.Println()

	// TCP
	tcp := scanPorts("tcp", "localhost", 1, 1024)
	tcp, tcpCount, tcpDots := preparePorts(tcp)
	fmt.Printf("[ tcp ] Открытые: %v%s Всего: %d\n", tcp, tcpDots, tcpCount)

	// UDP
	udp := scanPorts("udp", "localhost", 1, 1024)
	udp, udpCount, udpDots := preparePorts(udp)
	fmt.Printf("[ upd ] Открытые: %v%s Всего: %d\n", udp, udpDots, udpCount)
}

// Сканирование одного порта
func scanPort(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)

	conn, err := net.DialTimeout(protocol, address, 5*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}

// Сканирование нескольких портов
func scanPorts(protocol, hostname string, from, to int) []int {
	var res []int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := from; i <= to; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			if scanPort(protocol, hostname, port) {
				mu.Lock()
				defer mu.Unlock()
				res = append(res, port)
			}
		}(i)
	}
	wg.Wait()

	return res
}

// Подготовка списка портов
func preparePorts(ports []int) ([]int, int, string) {
	sort.Ints(ports)

	var dots string
	cap := 10
	count := len(ports)
	if count > cap {
		ports = ports[:cap]
		dots = "..."
	}

	return ports, count, dots
}
