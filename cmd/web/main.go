package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/erik-jenkins/snippetbox/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

// application contains all dependencies that need to be injected
// into the handlers
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.CookieStore
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
}

// logging
var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
var errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	secret := flag.String("secret", "secret", "The secret used to secure user sessions")
	dbHost := flag.String("dbhost", "mysql", "Hostname for MySQL server")
	dbPort := flag.String("dbport", "3306", "Port for MySQL server")
	dbUser := flag.String("dbuser", "snippetlord", "Username for MySQL user")
	dbPass := flag.String("dbpass", "password", "Password for MySQL user")
	dbRetries := flag.Int("dbretries", 10, "Number of retry attempts to connect to the DB")
	flag.Parse()

	// session management
	if *secret == "secret" {
		infoLog.Println("WARNING: using insecure session secret. Change this before going to production!")
	}
	session := sessions.NewCookieStore([]byte(*secret))

	// database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/snippetbox?parseTime=true", *dbUser, *dbPass, *dbHost, *dbPort)
	db, err := openDB(dsn, *dbRetries)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// template cache
	tc, err := newTemplateCache("./ui/html")
	if err != nil {
		errorLog.Fatal(err)
	}

	// application struct contains dependencies that are passed to the handlers
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: tc,
	}

	// tlsConfig contains the non-default TLS settings that we want to use
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s\n", *addr)
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	errorLog.Fatal(err)
}

// openDB opens a connection to the database and attempts to connect using
// exponential backoff. `dbRetries` is the number of times we should attempt
// to reconnect to the DB.
func openDB(dsn string, dbRetries int) (*sql.DB, error) {
	maxNumRetries := dbRetries
	waitPeriod := 1 * time.Second

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	for numTries := 1; numTries <= maxNumRetries; numTries++ {
		if err = db.Ping(); err == nil {
			return db, nil
		}

		infoLog.Printf("Failed to connect to the DB %d times. Waiting %v", numTries, waitPeriod)
		time.Sleep(waitPeriod)
		waitPeriod *= 2
	}

	return nil, err
}
