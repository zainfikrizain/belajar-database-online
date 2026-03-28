package main

import (
	"Tugas-2/database"
	"Tugas-2/handler"
	"Tugas-2/repositories"
	"Tugas-2/services"
	"log"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Config struct {
	Port    string `mapstructure:"PORT"`
	DB_CONN string `mapstructure:"DB_CONN"`
}

func main() {
	// config
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:    viper.GetString("PORT"),
		DB_CONN: viper.GetString("DB_CONN"),
	}

	// DB
	db, err := database.InitDB(config.DB_CONN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// DI
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// Echo
	e := echo.New()

	// routes
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Server is running 🚀")
	})

	e.GET("/api/produk", productHandler.GetAll)
	e.POST("/api/produk", productHandler.Create)
	e.GET("/api/produk/:id", productHandler.GetByID)
	e.PUT("/api/produk/:id", productHandler.Update)
	e.DELETE("/api/produk/:id", productHandler.Delete)
	e.GET("/api/produk/category", productHandler.GetCategoryByProductName)
	// start
	e.Logger.Fatal(e.Start(":" + config.Port))
}
