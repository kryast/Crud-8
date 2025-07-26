package main

import (
	"fmt"
	"log"

	"github.com/kryast/Crud-8.git/database"
	"github.com/kryast/Crud-8.git/models"
	"github.com/kryast/Crud-8.git/router"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to Connect DB", err)
	}

	db.AutoMigrate(&models.Employee{})

	r := router.SetupRouter(db)
	fmt.Println("Server Running on http://localhost:8080")
	r.Run(":8080")
}
