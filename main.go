package main

import (
	"final_project_golang/database"
	"final_project_golang/routes"
	"os"
)

// "final_project_golang/routes"

func main() {
	database.StartDB()

	r := routes.StartApp()
	port := os.Getenv("APP_PORT")
	r.Run(":" + port)
}
