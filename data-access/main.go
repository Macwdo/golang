package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Language struct {
	Id        int64
	Name      string
	Rating    float32
	CreatedAt string
}

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "languages",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	languageName := "Go"
	languages, err := languagesByName(languageName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Searching for languages by '%q' name\nFound: %v\n", languageName, languages)

	languages, err = allLanguages()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All languages:")
	for i, language := range languages {
		fmt.Printf("[%v][id=%v, name=%q, rating=%v]\n", i, language.Id, language.Name, language.Rating)
	}

}

func allLanguages() ([]Language, error) {
	var languages []Language

	rows, err := db.Query("SELECT * FROM languages")

	if err != nil {
		return nil, fmt.Errorf("allLanguages: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var language Language
		if err := rows.Scan(&language.Id, &language.Name, &language.Rating, &language.CreatedAt); err != nil {
			return nil, fmt.Errorf("allLanguages: %v", err)
		}
		languages = append(languages, language)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("allLanguages: %v", err)
	}

	return languages, nil

}

func languagesByName(name string) ([]Language, error) {
	var languages []Language

	rows, err := db.Query("SELECT * FROM languages WHERE name = ?", name)
	if err != nil {
		return nil, fmt.Errorf("languagesByName %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var language Language
		if err := rows.Scan(&language.Id, &language.Name, &language.Rating, &language.CreatedAt); err != nil {
			return nil, fmt.Errorf("languagesByName %q: %v", name, err)
		}
		languages = append(languages, language)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("languagesByName %q: %v", name, err)
	}
	return languages, nil
}
