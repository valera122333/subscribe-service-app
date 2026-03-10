package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"subscriptions-service/internal/config"
	"subscriptions-service/internal/handler"
	"subscriptions-service/internal/repository"
	"subscriptions-service/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {

	cfg := config.Load()

	var sqlDB *sql.DB
	var err error

	// ждём postgres
	for i := 0; i < 10; i++ {

		sqlDB, err = sql.Open("postgres", cfg.DB)
		if err == nil {
			err = sqlDB.Ping()
		}

		if err == nil {
			break
		}

		log.Println("waiting for postgres...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("cannot connect to postgres")
	}

	log.Println("connected to postgres")

	// запускаем миграции
	goose.SetDialect("postgres")

	if err := goose.Up(sqlDB, "migrations"); err != nil {
		log.Fatal(err)
	}

	log.Println("migrations applied")

	// создаём пул pgx
	db, err := pgxpool.New(context.Background(), cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.New(db)
	service := service.New(repo)
	handler := handler.New(service)

	http.HandleFunc("/subscriptions", handler.Create)
	http.HandleFunc("/subscriptions/list", handler.List)
	http.HandleFunc("/subscriptions/get", handler.Get)
	http.HandleFunc("/subscriptions/delete", handler.Delete)
	http.HandleFunc("/subscriptions/total", handler.Total)
	log.Println("server started")

	http.ListenAndServe(":"+cfg.Port, nil)
}
