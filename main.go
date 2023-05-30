package main

import (
	"cte.se/ya_individual_go_backend/data"
	"cte.se/ya_individual_go_backend/handlers"
)

// testar lite
var config Config

func main() {

	readConfig(&config)

	data.InitDatabase(config.Database.File,
		config.Database.Server,
		config.Database.Database,
		config.Database.Username,
		config.Database.Password,
		config.Database.Port)

	router := handlers.SetupRouter()
	router.Run(":8080")
}
