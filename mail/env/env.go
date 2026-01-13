package env

import (
	"os"
	"log"

	"github.com/pshebel/partiburo/mail/utils"
)


var (
	DB = os.Getenv("DB_PATH")
	AwsRegion = os.Getenv("AWS_REGION")
	AwsAccountId = os.Getenv("AWS_ACCOUNT_ID")
	Env = os.Getenv("PARTIBURO_ENV")
	AwsSender = os.Getenv("AWS_SENDER")
)

func init() {
	if !utils.IsValidEmail(AwsSender) {
		log.Println("invalid email")
	}
}
