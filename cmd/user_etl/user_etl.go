package main

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"forbole_code_test/provider"
	"forbole_code_test/provider/source"
	"forbole_code_test/repository"
	"forbole_code_test/service"

	_ "github.com/lib/pq"
)

var (
	dbHost     = os.Getenv("DB_HOST")
	dbPort     = os.Getenv("DB_PORT")
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName     = os.Getenv("DB_NAME")
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dsn)
	defer db.Close()

	if err != nil {
		panic(err)
	}

	userStore := repository.NewPostgresUserStore(db)
	userProvider := provider.NewUserProvider(source.NewRandomDataAPIUser())
	userService := service.NewUserService(userStore, userProvider)

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGSTOP)

	for {
		select {
		case <-done:
			fmt.Println("exiting...")
			return
		default:
			interval := time.Duration(250+rand.Intn(500)) * time.Millisecond
			fmt.Printf("fetching and creating user after %v\n", interval)

			time.Sleep(interval)

			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			go func() {
				defer cancel()

				err := userService.FetchAndCreateUser(ctx)
				if err != nil {
					fmt.Printf("error: %v\n", err)
				}
			}()
		}
	}
}
