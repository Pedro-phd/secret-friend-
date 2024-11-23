package infra

import (
	"fmt"
	"secret-friend/domain/user"

	"gorm.io/gorm"
)

func SetupInfra(db *gorm.DB) {
	db.AutoMigrate(&user.User{})

	Session := db.Session(&gorm.Session{PrepareStmt: true})

	if Session != nil {
		fmt.Println("Migration successful")
	}
}
