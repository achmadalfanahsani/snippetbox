package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"snippetbox.achmadalfanahsani.com/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Snippets *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
    
	flag.Parse()

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	app := &application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
		Snippets: &models.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}


func openDB(dsn string) (*sql.DB, error) {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return nil, err
		}

		if err = db.Ping(); err != nil {
			return nil, err
		}

		return db, nil
}