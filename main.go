package main

import (
	"Portfolio_Nodes/app"
	"Portfolio_Nodes/bootstrap"
)

func main() {
	err := app.InitApp()
	if err != nil {
		panic(err)
	}

	//// Initialize Database
	_, err = app.InitDatabase()
	if err != nil {
		panic(err)
	}
	err = app.RunMigrations()
	if err != nil {
		panic(err)
	}

	// REST
	err = bootstrap.InitRest()
	if err != nil {
		panic(err)
	}
}
