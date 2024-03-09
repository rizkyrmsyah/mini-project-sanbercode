package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rizkyrmsyah/mini-project-sanbercode/controllers"
	"github.com/rizkyrmsyah/mini-project-sanbercode/database"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	// env config
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("success load file environment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection failed")
		panic(err)
	} else {
		fmt.Println("DB Connection success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	// router gin
	router := gin.Default()
	router.GET("/persons", controllers.GetAllPerson)
	router.POST("/persons", controllers.InsertPerson)
	router.PUT("/persons/:id", controllers.UpdatetPerson)
	router.DELETE("/persons/:id", controllers.DeletePerson)

	router.Run("localhost:8080")
}
