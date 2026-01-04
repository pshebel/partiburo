package env

import (
	"fmt"
	"os"
	"strings"
)


var (
	Port = os.Getenv("API_PORT")
	AllowedOriginsStr = os.Getenv("ALLOWED_ORIGINS")
	AllowedOrigins = []string{}
	DB = os.Getenv("DB_PATH")
)

func init() {

	fmt.Println(Port)
	fmt.Println(DB)
	if AllowedOriginsStr != "" {
		AllowedOrigins = strings.Split(AllowedOriginsStr, ",")
	}
}
