package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Структура "моя ошибка"
type MyError struct {
	When time.Time
	What string
}

// Интерфейс "моя ошибка"
type MyErrorI interface {
	as()
}

// Пустой метод для интерфейса
func (e *MyError) as() {}

// Вывод моей ошибки
func (e *MyError) Error() string {
	return fmt.Sprintf(
		"В %v\nПроизошло: %s",
		e.When.Format("2006/01/02 03:04:05"), e.What,
	)
}

// Генерация ошибки
func run() error {
	return &MyError{
		time.Now(),
		"Фиаско, братан!",
	}
}

// Генерация своей ошибки
func myRun() error {
	var e error
	e = errors.New("НИЧИВО-НИ-ВЫШЛА!") // Можно так, ...
	e = fmt.Errorf("НИЧИВО-НИ-ВЫШЛА!") // ... а можно и этак
	return e
}

// Мультиобработка ошибок
func multiErrors(val1, val2 int) (_ int, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("в multiErrors: %w", err)
		}
	}()

	val3, err := func(v int) (int, error) {
		if v <= 0 {
			return 0, errors.New("нет значения val1")
		}
		return v * 2, nil
	}(val1)
	if err != nil {
		return 0, err
	}

	val4, err := func(v int) (int, error) {
		if v <= 0 {
			return 0, errors.New("нет значения val2")
		}
		return v * 3, nil
	}(val2)
	if err != nil {
		return 0, err
	}

	return func(v1, v2 int) (int, error) {
		if v1 <= 0 || v2 <= 0 {
			return 0, errors.New("нет значений val3 и val4")
		}
		return v1 + v2, nil
	}(val3, val4)
}

// Проверка наличия файла
func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("в fileChecker: %w", err)
	}
	f.Close()
	return nil
}

// Структура "срез ошибок"
type SliceError []error

// Вывод ошибок в строку
func (errs SliceError) Error() string {
	var out string
	separator := ", "
	for _, err := range errs {
		out += err.Error() + separator
	}
	return strings.TrimRight(out, separator)
}

// Проверка строки
func stringChecker(input string) error {
	var (
		err      SliceError
		spaces   int
		hasDigit bool
	)
	if len([]rune(input)) >= 20 {
		err = SliceError{errors.New("строка слишком длинная")}
	}
	for _, ch := range input {
		if ch == ' ' {
			spaces++
		} else if ch >= '0' && ch <= '9' {
			hasDigit = true
		}
	}
	if hasDigit {
		err = append(err, errors.New("найдены числа"))
	}
	if spaces != 2 {
		err = append(err, errors.New("нет двух пробелов"))
	}
	if len(err) == 0 {
		return nil
	}
	return err
}

// Обработчик паники
func saveFromPanic() {
	err := recover()
	fmt.Println("Обычная функция:", err)
	panic(err)
}

// Продолжение выполнения с паникой
func continueWithPanic(i int) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("ПАНИКА:", p)
		}
	}()
	fmt.Println(30 / i)
}

func main() {
	fmt.Println(" \n[ ОШИБКИ ]\n ")

	// Ошибки
	fmt.Println("Кастомные ошибки:")
	fmt.Println()
	if err := run(); err != nil {
		switch e := err.(type) {
		case *MyError:
			fmt.Println(e)
		default:
			fmt.Println("Неизвестная ошибка:", e)
		}
	}
	if myErr := myRun(); myErr != nil {
		fmt.Println(myErr)
	}
	fmt.Println()

	// Мультиобработка
	fmt.Println("Мультиобработка:")
	multi, err := multiErrors(2, 0)
	if err != nil {
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			err = wrappedErr
		}
	}
	fmt.Println(multi, err)
	fmt.Println()

	// IS и AS
	err = fileChecker("dont_exist.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("ОШИБКА IS: файла не существует")
		}
	}
	err = run()
	var MyErr MyErrorI
	if errors.As(err, &MyErr) {
		fmt.Println("ОШИБКА AS: совпадение типов ошибок")
	}
	fmt.Println()

	// Проверка введенной строки
	fmt.Printf("Введите строку (2 пробела, <20 символов, без чисел): ")
	reader := bufio.NewReader(os.Stdin)
	ret, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	if err = stringChecker(ret); err != nil {
		fmt.Println("Ошибки:", err)
	} else {
		fmt.Println("Все ОК!")
	}
	fmt.Println()

	// Паника
	fmt.Println("Обработка паники:")
	for _, val := range []int{1, 2, 0, 6} {
		continueWithPanic(val)
	}
	fmt.Println()
	fmt.Println("Вызов паники:")
	defer func() {
		err := recover()
		fmt.Println("Анонимная функция:", err)
	}()
	defer saveFromPanic()
	panic("Паника!")
}
