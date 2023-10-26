package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	r := gin.Default()

	// Cambia la cadena de conexión a la base de datos.
	dsn := "root:@tcp(127.0.0.1:3306)/usuarios?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// AutoMigrate solo crea la tabla si no existe, no modifica la estructura si ya existe.
	db.AutoMigrate(&User{})

	// Endpoint para obtener todos los usuarios
	r.GET("/users", func(c *gin.Context) {
		var users []User
		db.Find(&users)

		c.JSON(200, gin.H{
			"data": users,
		})
	})

	// Endpoint para obtener un usuario por su ID
	r.GET("/users/:id", func(c *gin.Context) {
		var user User

		if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
			c.JSON(400, gin.H{"error": "Record not found!"})
			return
		}

		c.JSON(200, gin.H{
			"data": user,
		})
	})

	// Endpoint para crear un nuevo usuario
	r.POST("/users", func(c *gin.Context) {
		var user User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		db.Create(&user)

		c.JSON(200, gin.H{
			"data": user,
		})
	})

	// Endpoint para actualizar un usuario por su ID
	r.PUT("/users/:id", func(c *gin.Context) {
		var user User

		if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
			c.JSON(400, gin.H{"error": "Record not found!"})
			return
		}

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		db.Save(&user)

		c.JSON(200, gin.H{
			"data": user,
		})
	})

	// Endpoint para eliminar un usuario por su ID
	r.DELETE("/users/:id", func(c *gin.Context) {
		var user User

		if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
			c.JSON(400, gin.H{"error": "Record not found!"})
			return
		}

		db.Delete(&user)

		c.JSON(200, gin.H{
			"message": "User deleted successfully!",
		})
	})

	r.Run(":8080")
}
