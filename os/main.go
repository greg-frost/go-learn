package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"
)

// Проверка зависимости
func checkDependency(name string) error {
	if _, err := exec.LookPath(name); err != nil {
		msg := "Не удалось найти '%s' в PATH: %s"
		return fmt.Errorf(msg, name, err)
	}
	return nil
}

func main() {
	fmt.Println(" \n[ OS ]\n ")

	// ОС
	fmt.Print("Операционная система: ")
	OS := runtime.GOOS
	switch OS {
	case "windows":
		fmt.Println("WINDOWS")
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("LINUX")
	default:
		fmt.Println(strings.ToUpper(OS))
	}

	// Число ядер
	numCPU := runtime.NumCPU()
	fmt.Println("Число ядер процессора:", numCPU)
	fmt.Println()

	// Идентификатор процесса
	pid := os.Getpid()
	fmt.Println("ID процесса:", pid)

	// Разделители пути
	pathSep := string(os.PathSeparator)
	pathListSep := string(os.PathListSeparator)
	fmt.Println("Разделители пути:", pathSep, "и", pathListSep)

	// Текущий каталог
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Текущий каталог:", pwd)
	fmt.Println()

	// Хост
	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Имя хоста:", host)

	// IP-адреса
	addrs, err := net.LookupHost(host)
	if err != nil {
		log.Fatal(err)
	}
	sort.Strings(addrs)
	fmt.Println("IP-адреса:")
	for _, addr := range addrs {
		parts := strings.Split(addr, "%")
		fmt.Println(parts[0])
	}
	fmt.Println()

	// Зависимости
	fmt.Println("Зависимости:")
	err = checkDependency("ping")
	fmt.Print("ping - ")
	if err == nil {
		fmt.Println("доступно")
	} else {
		fmt.Println("недоступно")
	}
	err = checkDependency("pong")
	fmt.Print("pong - ")
	if err == nil {
		fmt.Println("доступно")
	} else {
		fmt.Println("недоступно")
	}
}
