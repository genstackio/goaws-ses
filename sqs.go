package goaws_ses

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func GetQueueURL(c context.Context, api SQSDeleteMessageAPI, input *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
	return api.GetQueueUrl(c, input)
}

func RemoveMessage(c context.Context, api SQSDeleteMessageAPI, input *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
	return api.DeleteMessage(c, input)
}

//goland:noinspection GoUnusedExportedFunction
func ProcessSesEventSqsHandler(data []byte, receiptHandle string, queueName string, onDelivery onDeliveryFn, onComplain onComplaintFn, onBounce onBounceFn) (interface{}, error) {
	body, err := ExtractSesNotificationBodyFromSqsBody(data)

	if nil != err {
		return nil, err
	}

	result, err := ProcessSesEvent(body, onDelivery, onComplain, onBounce)

	if nil != err {
		return nil, err
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	client := sqs.NewFromConfig(cfg)

	qUInput := &sqs.GetQueueUrlInput{
		QueueName: &queueName,
	}
	result2, err2 := GetQueueURL(context.TODO(), client, qUInput)
	if err2 != nil {
		return nil, err2
	}

	queueURL := result2.QueueUrl

	dMInput := &sqs.DeleteMessageInput{
		QueueUrl:      queueURL,
		ReceiptHandle: &receiptHandle,
	}

	_, err = RemoveMessage(context.TODO(), client, dMInput)
	if err != nil {
		return nil, err
	}

	return result, nil
}
