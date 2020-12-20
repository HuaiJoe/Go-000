package main

import (
	"fmt"
	"log"
	"net/http"
	hp "week04/api/http"
	config "week04/configs"
	"week04/internal/infrastructure/mysql"
	"week04/internal/service"
	"week04/pkg/database"
)

func main() {
	engine, err := database.NewDBEngine(
		config.MYSQL_HOST,
		config.MYSQL_PORT,
		config.MYSQL_USER,
		config.MYSQL_PASSWORD,
	)
	if err != nil {
		fmt.Print("error")
		log.Fatal("create connections failed.", err)
	}
	defer engine.Close()

	repo := mysql.NewUserRepo(engine)
	userCaseService := service.NewUserCase(repo)
	mux := http.NewServeMux()

	hp.MakeUserHandlers(mux, userCaseService)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
