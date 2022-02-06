package configuracion

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Configuracion func to get env value from key
func Configuracion(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file. ", err.Error())
	}
	return os.Getenv(key)
}
