package internal

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"

	"encoding/json"
)

type Request struct {
	Email string
	Code  string
}

func SendEmail(email string, code string) error {
	sess := session.Must(session.NewSession())
	client := lambda.New(sess, &aws.Config{
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			""),
		Region: aws.String(endpoints.UsEast2RegionID),
	})
	payload, err := json.Marshal(Request{email, code})
	if err != nil {
		log.Println("Error marshalling MyGetItemsFunction request", err)
	}

	_, err = client.Invoke(&lambda.InvokeInput{FunctionName: aws.String("ses"), Payload: payload})
	if err != nil {
		log.Println("Error calling ses", err)
	}
	return err
}
