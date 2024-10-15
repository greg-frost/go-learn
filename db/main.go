package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

// Дескриптор БД
var db *sql.DB

// Подготовленный запрос
var stmt *sql.Stmt

// Структура "альбом"
type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

// Получение списка альбомов по артисту
func albumsByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var a Album
		if err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, a)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

// Получение списка альбомов по артисту (с контекстом и таймаутом)
func albumsByArtistContext(ctx context.Context, timeout time.Duration, name string) ([]Album, error) {
	var albums []Album

	queryCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	rows, err := db.QueryContext(queryCtx, "SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtistContext %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var a Album
		if err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtistContext %q: %v", name, err)
		}
		albums = append(albums, a)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtistContext %q: %v", name, err)
	}
	return albums, nil
}

// Получение альбома по ID
func albumByID(id int64) (Album, error) {
	var a Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
		if err == sql.ErrNoRows {
			return a, fmt.Errorf("albumById %d: нет такого альбома", id)
		}
		return a, fmt.Errorf("albumById %d: %v", id, err)
	}
	return a, nil
}

// Получение альбома по ID (подготовленный запрос)
func albumByIDPrepared(id int64) (Album, error) {
	if stmt == nil {
		var err error
		stmt, err = db.Prepare("SELECT * FROM album WHERE id = ?")
		if err != nil {
			log.Fatal(err)
		}
		//defer stmt.Close()
	}

	var a Album

	if err := stmt.QueryRow(id).Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
		if err == sql.ErrNoRows {
			return a, fmt.Errorf("albumByIDPrepared %d: нет такого альбома", id)
		}
		return a, fmt.Errorf("albumByIDPrepared %d: %v", id, err)
	}
	return a, nil
}

func main() {
	fmt.Println(" \n[ БАЗА ДАННЫХ ]\n ")

	/* Подключение */

	// БД, адрес, логин и пароль
	dbname := "learn"
	addr := "127.0.0.1:3306"
	username := os.Getenv("DB_USER")
	if username == "" {
		username = "root"
	}
	password := os.Getenv("DB_PASS")

	sep := "   "
	var err error

	// Вариант 1
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		username, password, addr, dbname)
	db, err = sql.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}

	// Конфигурация
	cfg := mysql.Config{
		User:                 username,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 addr,
		DBName:               dbname,
		AllowNativePasswords: true,
	}

	// Вариант 2
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// Вариант 3
	connector, err := mysql.NewConnector(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	db = sql.OpenDB(connector)

	defer db.Close()

	// Пинг
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Успешное подключение")
	fmt.Println()

	// Удаление старой таблицы
	_, err = db.Exec("DROP TABLE IF EXISTS album")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Старая таблица удалена")

	// Создание новой таблицы
	_, err = db.Exec(`
		CREATE TABLE album (
			id INT AUTO_INCREMENT NOT NULL,
			title VARCHAR(128) NOT NULL,
			artist VARCHAR(255) NOT NULL,
			price DECIMAL(5,2) NOT NULL,
			PRIMARY KEY (id))
		`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Новая таблица создана")

	// Вставка записей
	insert, err := db.Exec(`
		INSERT INTO album
			(title, artist, price)
		VALUES
			('Blue Train', 'John Coltrane', 56.99),
			('Giant Steps', 'John Coltrane', 63.99),
			('Jeru', 'Gerry Mulligan', 17.99),
			('Jeru (remastered)', 'Gerry Mulligan', 19.99),
			('Sarah Vaughan', 'Sarah Vaughan', 34.98)
		`)
	if err != nil {
		log.Fatal(err)
	}
	inserted, _ := insert.RowsAffected()
	fmt.Printf("Добавлены записи (%d)\n", inserted)
	fmt.Println()

	// Поиск альбомов по артисту
	fmt.Println("Поиск по артисту:")
	fmt.Println()

	artists := []string{"John Coltrane", "Jack Cocktail"}
	for _, artist := range artists {
		albums, err := albumsByArtist(artist)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Альбомы артиста %q:\n", artist)
		if len(albums) == 0 {
			fmt.Println(sep + "(не найдено)")
		}
		for _, a := range albums {
			fmt.Printf("%s%d - %s ($%0.2f)\n", sep, a.ID, a.Title, a.Price)
		}
	}
	fmt.Println()

	// Поиск альбомов по ID
	fmt.Println("Поиск по ID:")
	fmt.Println()

	ids := []int64{3, 6}
	for _, id := range ids {
		a, err := albumByID(id)
		fmt.Printf("Альбом c ID = %d:\n", id)
		if err != nil {
			fmt.Println(sep + "(не найдено)")
		} else {
			fmt.Printf("%s%s, %s ($%0.2f)\n", sep, a.Title, a.Artist, a.Price)
		}
	}
	fmt.Println()

	// Подготовленные запросы
	fmt.Println("Сравнение запросов:")
	fmt.Println()

	times := 1000

	start := time.Now()
	for i := 0; i < times; i++ {
		albumByID(int64(i%5 + 1))
	}
	fmt.Println("Обычные:", time.Now().Sub(start))

	start = time.Now()
	for i := 0; i < times; i++ {
		albumByIDPrepared(int64(i%5 + 1))
	}
	fmt.Println("Подготовленные:", time.Now().Sub(start))
	fmt.Println()

	// Убаление таблицы
	_, err = db.Exec("DROP TABLE album")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Таблица удалена")
}
