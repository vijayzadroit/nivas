package main

import (
	"fmt"
	"log"
	service "nivasBackendMain/Helper/MinIo"
	"nivasBackendMain/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	err := godotenv.Load()

	if err != nil {
		log.Fatal("❌Error loading .env file")
	}

	if err := service.InitMinioClient(); err != nil {
		log.Fatalf("❌ MinIO initialization failed: %v", err)
	}

	r.SetTrustedProxies(nil)

	// ✅ CORS configuration to allow only one origin
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:3000"}, // Change to your allowed origin
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true // allow all origins dynamically
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// BrandRegistration.BrandRegistrationRoutes(r)
	routes.MainRoutes(r)

	fmt.Println("✅Server is Running at Port:" + os.Getenv("PORT"))
	r.Run("0.0.0.0:" + os.Getenv("PORT"))

}
