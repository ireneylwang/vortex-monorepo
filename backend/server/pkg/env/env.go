package env

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func Load(path string) {
	fmt.Println("Loading .env file from", path)
	_ = godotenv.Load(path)
	gin.SetMode(os.Getenv("GIN_MODE"))
}

func OpenApiHost() string {
	return os.Getenv("OPEN_API_HOST")
}
