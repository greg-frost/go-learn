package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"sort"
	"strings"
)

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
}
