package main

import "fmt"

type App struct{}

func (app *App) Run() error {
	fmt.Println("Running App")
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
