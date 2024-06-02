package main

import (
	"errors"
	"fmt"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

// Запущенные сервисы
var services = map[int]bool{
	1: true,
	2: true,
	3: true,
}

func main() {
	fmt.Println(" \n[ CLI-ПРИЛОЖЕНИЕ ]\n ")

	// Инициализация
	app := cli.NewApp()
	app.Name = "CLI Go App"
	app.Usage = "Пример приложения для работы с флагами и командами"

	// Флаги
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Value: "Greg",
			Usage: "Имя для приветствия",
		},
		cli.BoolFlag{
			Name:  "english, e",
			Usage: "Язык приветствия",
		},
	}

	// Команды
	app.Commands = []cli.Command{
		{
			Name:      "up",
			ShortName: "u",
			Usage:     "Запуск сервиса",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "service, id",
					Usage: "ID сервиса",
					Value: 1,
				},
			},
			Action: func(c *cli.Context) error {
				id := c.Int("service")
				if services[id] {
					msg := fmt.Sprintf("сервис #%d уже запущен!", id)
					fmt.Println("Ошибка:", msg)
					return errors.New(msg)
				}

				fmt.Printf("Сервис #%d успешно запущен.\n", id)
				services[id] = true
				return nil
			},
		},
		{
			Name:      "down",
			ShortName: "d",
			Usage:     "Остановка сервиса",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "service, id",
					Usage: "ID сервиса",
					Value: 1,
				},
			},
			Action: func(c *cli.Context) error {
				id := c.Int("service")
				if !services[id] {
					msg := fmt.Sprintf("сервис #%d еще не запущен!", id)
					fmt.Println("Ошибка:", msg)
					return errors.New(msg)
				}

				fmt.Printf("Сервис #%d успешно остановлен.\n", id)
				services[id] = false
				return nil
			},
		},
	}

	// Действие по умолчанию
	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("name")
		english := c.GlobalBool("english")

		if english {
			fmt.Printf("Hello, %s!\n", name)
		} else {
			fmt.Printf("Привет, %s!\n", name)
		}

		return nil
	}

	// Запуск
	app.Run(os.Args)
}
