package internal

// snippet-comment:[These are tags for the AWS doc team's sample catalog. Do not remove.]
// snippet-sourceauthor:[Doug-AWS]
// snippet-sourcedescription:[Runs a Lambda function.]
// snippet-keyword:[AWS Lambda]
// snippet-keyword:[Invoke function]
// snippet-keyword:[Go]
// snippet-sourcesyntax:[go]
// snippet-service:[lambda]
// snippet-keyword:[Code Sample]
// snippet-sourcetype:[full-example]
// snippet-sourcedate:[2018-03-16]
/*
 Copyright 2010-2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 This file is licensed under the Apache License, Version 2.0 (the "License").
 You may not use this file except in compliance with the License. A copy of the
 License is located at
 http://aws.amazon.com/apache2.0/
 This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
 OF ANY KIND, either express or implied. See the License for the specific
 language governing permissions and limitations under the License.
*/

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
