package sqs

import (
	"encoding/json"
	"fmt"

	"github.com/xtianatilano/christian-golang-training-beginner/app/cmd/helpers"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Publisher struct {
	SQS      *sqs.SQS
	QueueUrl *string
}

func NewPublisher(region, endpoint string) (*Publisher, error) {
	cfg := aws.Config{
		Region: aws.String(region),
	}
	// if endpoint is not empty, we will use localstack
	if endpoint != "" {
		cfg.Endpoint = aws.String(endpoint)
	}

	sess := session.Must(session.NewSession(&cfg))
	var p Publisher
	sqsClient := sqs.New(sess)

	p.SQS = sqsClient
	queueUrl := helpers.MustHaveEnv("SQS_QUEUE_URL")
	p.QueueUrl = &queueUrl

	return &p, nil
}

func (p Publisher) Publish(msg interface{}) error {
	messageBody, _ := json.Marshal(msg)

	res, err := p.SQS.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(messageBody)),
		QueueUrl:    p.QueueUrl,
	})

	fmt.Println(res)

	return err
}
