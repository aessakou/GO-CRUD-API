package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type TODO struct {
	ID        uint   `gorm:"primaryKey"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load env file")
	}

	PORT := os.Getenv("PORT")

	dbhost := os.Getenv("DBHOST")
	dbuser := os.Getenv("DBUSER")
	dbuserpw := os.Getenv("DBUSERPW")
	dbname := os.Getenv("DBNAME")
	dbport := os.Getenv("DBPORT")

	dsn := "host=" + dbhost + " user=" + dbuser + " password=" + dbuserpw + " dbname=" + dbname + " port=" + dbport + " sslmode=disable"
	db, erro := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if erro != nil {
		log.Fatal("Failed to connect to database!")
	}
	fmt.Println("The PostgreSQL database connected successfully.")
	db.AutoMigrate(&TODO{})
	fmt.Println("TODO has been migrated to the PostgreSQL database.")
	DB = db

	app := fiber.New()

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get database object:", err)
	}

	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Fatal("Failed to close database connection:", err)
		}
		fmt.Println("The database connection was closed successfully.")
	}()

	// app.Use(logger.New())

	// Get todos
	app.Get("/api/todos", get_todos)

	// Create a TODO
	app.Post("/api/todos", create_todo)

	// Update a TODO
	app.Patch("/api/todos/:id", update_todo)

	// Delete a TODO
	app.Delete("/api/todos/:id", delete_todo)

	app.Listen(":" + PORT)

}

func get_todos(c *fiber.Ctx) error {
	var todos []TODO
	if err := DB.Order("id").Find(&todos).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}
	return c.Status(200).JSON(todos)
}

func create_todo(c *fiber.Ctx) error {
	todo := &TODO{}

	if err := c.BodyParser(todo); err != nil {
		return err
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "The Body field required!"})
	}

	if err := DB.Create(&todo).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}

	return c.Status(201).JSON(todo)

}

func update_todo(c *fiber.Ctx) error {
	id := c.Params("id")

	err := DB.Model(&TODO{}).Where("id = ?", id).Update("Completed", "true").Error
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{"msg": "TODO has been successfully updated"})
}

func delete_todo(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := DB.Where("id = ?", id).Delete(&TODO{}).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err})
	}
	return c.Status(200).JSON(fiber.Map{"msg": "TODO has been successfully deleted!"})
}
