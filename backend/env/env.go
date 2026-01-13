package env

import (
	"fmt"
	"os"
	"log"
	"strings"

	"github.com/pshebel/partiburo/backend/utils"
)


var (
	Port = os.Getenv("API_PORT")
	AllowedOriginsStr = os.Getenv("ALLOWED_ORIGINS")
	AllowedOrigins = []string{}
	DB = os.Getenv("DB_PATH")
	AwsRegion = os.Getenv("AWS_REGION")
	AwsAccountId = os.Getenv("AWS_ACCOUNT_ID")
	Env = os.Getenv("PARTIBURO_ENV")
	AwsSender = os.Getenv("AWS_SENDER")
)

func init() {

	fmt.Println(Port)
	fmt.Println(DB)
	if AllowedOriginsStr != "" {
		AllowedOrigins = strings.Split(AllowedOriginsStr, ",")
	}

	if !utils.IsValidEmail(AwsSender) {
		log.Println("invalid email")
	}
}
