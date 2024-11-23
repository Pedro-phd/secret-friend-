package main

import (
	"log"
	"os"
	"secret-friend/domain/user"
	"secret-friend/infra"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DSN")

	engine := gin.Default()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	infra.SetupInfra(db)

	userController := user.NewUserController(engine, db)
	userController.SetupUser()

	engine.Run(":8080")

}
