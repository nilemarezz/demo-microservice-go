package connection

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"github.com/uptrace/opentelemetry-go-extra/otelsqlx"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var MySqlDB *sqlx.DB

func ConnectMySQLDB() {
	dns := fmt.Sprintf("%v:%v@tcp(%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	count := 0
	for {
		connection, err := openDB(dns)
		if err != nil {
			log.Println("Database not ready...", dns)
			count++
		} else {
			log.Println("Database connected", dns)
			MySqlDB = connection
			break
		}

		if count > 10 {
			log.Panicln(err)
			os.Exit(1)
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}

func openDB(dns string) (*sqlx.DB, error) {
	db, err := otelsqlx.Open(os.Getenv("DB_DRIVER"), dns, otelsql.WithAttributes(semconv.DBSystemMySQL))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, err
}
