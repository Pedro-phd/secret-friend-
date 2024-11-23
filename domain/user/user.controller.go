package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	userService *Service
	engine      *gin.Engine
	db          *gorm.DB
}

func NewUserController(engine *gin.Engine, db *gorm.DB) *Controller { // construtor
	return &Controller{
		engine: engine,
		db:     db,
	}
}

func (c *Controller) SetupUser() {

	service := NewUserService(c.db)

	c.engine.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		res, err := service.GetUserById(id)

		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, res)
	})
}
