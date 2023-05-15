package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := sql.Open("mysql", connStr)
	log.Printf("connStr: %s", connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connect DB")
	defer db.Close()

	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		lang := r.URL.Query().Get("lang")
		log.Printf("lang: %s", lang)
		sqlStr := "SELECT text FROM greetings WHERE lang = ?;"

		rows, err := db.Query(sqlStr, lang)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var hello string
		for rows.Next() {
			if err := rows.Scan(&hello); err != nil {
				log.Fatal(err)
			}
		}
		// FIXME: 日本語がブラウザで文字化けする表示される。アプリ側の文字コード設定の問題かもしれない。
		json.NewEncoder(w).Encode(hello)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
