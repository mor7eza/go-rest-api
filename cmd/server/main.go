package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/mor7eza/go-rest-api.git/internal/transport/http"
)

type App struct{}

func (app *App) Run() error {
	fmt.Println("Running App")

	handler := transportHTTP.NewHanlder()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup routes")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go Rest Api")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting REST Api")
		fmt.Println(err)
	}

}
