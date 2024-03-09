package main

import (
	"database/sql"
	"fmt"
	"os"
	"practice/practice/database"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sahrulyo/sanbercode-batch54-mini-project-deployment/controllers"
	"github.com/sahrulyo/sanbercode-batch54-mini-project-deployment/database"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "Ulyasar10389#"
	dbname   = "user"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("DB Connection failed")
		panic(err)
	}

	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection failed")
		panic(err)
	}

	fmt.Println("DB Connection Success")

	database.DbMigrate(DB)
	defer DB.Close()

	// router gin
	router := gin.Default()
	router.GET("/persons", controllers.GetAllPerson)
	router.POST("/persons", controllers.InsertPerson)
	router.PUT("/persons/:id", controllers.UpdatetPerson)
	router.DELETE("/persons/:id", controllers.DeletePerson)

	router.Run(":" + os.Getenv("PORT"))
}
