package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/0x000def42/microshards-go-config/app/admin"
	"github.com/0x000def42/microshards-go-config/factories"
	"github.com/0x000def42/microshards-go-config/repositories"
	"github.com/0x000def42/microshards-go-config/request_handlers"
	"github.com/joho/godotenv"
	"github.com/qiangxue/go-env"
)

type Config struct {
	DbSqliteName string `env:"DB_SQLITE"`
}

func main() {

	godotenv.Load()

	var cfg Config
	loader := env.New("APP_", log.Printf)
	if err := loader.Load(&cfg); err != nil {
		panic(err)
	}

	sqliteClient := factories.NewSqliteClient(cfg.DbSqliteName)

	userRepository := repositories.NewUserRepositorySqlite(sqliteClient)

	adminUserService := admin.NewUserService(userRepository)

	adminModule := admin.NewModule(adminUserService)

	requestHandlersHttp := []request_handlers.RequestHandlerHttp{}
	requestHandlersHttp = append(requestHandlersHttp, request_handlers.NewRequestHandlerHttpAdmin(adminModule))

	httpServer := factories.NewHttpServer(requestHandlersHttp)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	httpServer.Shutdown(ctx)

}
