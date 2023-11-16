package main

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var client42 = OauthClient{
	Url:	"https://api.intra.42.fr/oauth/token",
	Uid:	os.Getenv("BACKEND_42_UID"),
	Secret:	os.Getenv("BACKEND_42_SECRET"),
}

var clientTwitch = OauthClient{
	Url:	"https://id.twitch.tv/oauth2/token",
	Uid:	os.Getenv("BACKEND_TWITCH_UID"),
	Secret:	os.Getenv("BACKEND_TWITCH_SECRET"),
}

var db *sql.DB

func dbTextHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	queryTitle := req.URL.Query().Get("title")
	if queryTitle == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rows, err := squirrel.
		Select("value").
		From("texts").
		Where(squirrel.Eq{"title": queryTitle}).
		Limit(1).
		RunWith(db).Query()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	if !rows.Next() {
		w.WriteHeader(http.StatusNotFound)
	}

	text := ""
	err = rows.Scan(&text)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(text))
	return
}

func initDbTexts() error {
	textsDir, err := os.ReadDir("./data/texts")
	if err != nil {
		return err
	}
	for _, file := range textsDir {
		textData, err := os.ReadFile("./data/texts/" + file.Name())
		if err != nil {
			return err
		}

		fileName := strings.ReplaceAll(file.Name(), filepath.Ext(file.Name()), "")

		_, err = squirrel.
			Insert("texts").
			Columns("title", "value").
			Values(fileName, string(textData)).
			RunWith(db).Exec()
		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE constraint failed") {
				continue
			}
			return err
		}
	}

	return nil
}

func dbImgHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	queryTitle := req.URL.Query().Get("title")
	if queryTitle == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rows, err := squirrel.
		Select("value").
		From("imgs").
		Where(squirrel.Eq{"title": queryTitle}).
		Limit(1).
		RunWith(db).Query()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	if !rows.Next() {
		w.WriteHeader(http.StatusNotFound)
	}

	text := ""
	err = rows.Scan(&text)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(text))
	return
}

func initDbImgs() error {
	imgsDir, err := os.ReadDir("./data/imgs")
	if err != nil {
		return err
	}
	for _, file := range imgsDir {
		imgData, err := os.ReadFile("./data/imgs/" + file.Name())
		if err != nil {
			return err
		}

		fileName := strings.ReplaceAll(file.Name(), filepath.Ext(file.Name()), "")

		_, err = squirrel.
			Insert("imgs").
			Columns("title", "value").
			Values(fileName, string(imgData)).
			RunWith(db).Exec()
		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE constraint failed") {
				continue
			}
			return err
		}
	}

	return nil
}

func initDb(dsn string) error {
	var err error

	db, err = sql.Open("sqlite3", dsn)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS texts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL UNIQUE,
			value TEXT NOT NULL
		);
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS imgs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL UNIQUE,
		value TEXT NOT NULL
	);
	`)
	if err != nil {
		return err
	}

	err = initDbTexts()
	if err != nil {
		return err
	}

	err = initDbImgs()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := initDb("file:./data/db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	client42.Token.AccessToken = ""
	clientTwitch.Token.AccessToken = ""
	http.HandleFunc("/me", meHandler42)
	http.HandleFunc("/twitchapi", twitchHandler)
	http.HandleFunc("/db/text", dbTextHandler)
	http.HandleFunc("/db/img", dbImgHandler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
