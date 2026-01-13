package notifications

import (
	"log"
	"fmt"
	"context"
	
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"

	"github.com/pshebel/partiburo/mail/env"
)


var Client *sesv2.Client

func init() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(env.AwsRegion))
	if err != nil {
		log.Println(err)
		return
	}
	Client = sesv2.NewFromConfig(cfg)
}


func GetSesClient() *sesv2.Client {
	return Client
}

func GetContactListName(partyId int64) string {
	return fmt.Sprintf("partiburo_%s_%d", env.Env, partyId)
}
