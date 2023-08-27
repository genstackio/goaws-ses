package goaws_ses

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
)

//goland:noinspection GoUnusedExportedFunction
func ExtractSesNotificationBodyFromSqsBody(raw []byte) ([]byte, error) {
	var d events.SNSEntity
	err := json.Unmarshal(raw, &d)
	if nil != err {
		return []byte{}, err
	}
	return []byte(d.Message), nil
}

//goland:noinspection GoUnusedExportedFunction
func ProcessSesEvent(event []byte, onDelivery onDeliveryFn, onComplain onComplaintFn, onBounce onBounceFn) (SesEventProcessedResponse, error) {
	var data SesEventBasic
	err := json.Unmarshal(event, &data)
	if err != nil {
		return SesEventProcessedResponse{}, err
	}
	switch data.NotificationType {
	case "Delivery":
		return processSesDeliveryEvent(event, onDelivery)
	case "Complaint":
		return processSesComplaintEvent(event, onComplain)
	case "Bounce":
		return processSesBounceEvent(event, onBounce)
	default:
		return SesEventProcessedResponse{}, errors.New("unknown ses notification type " + data.NotificationType)
	}
}

func processSesDeliveryEvent(event []byte, callback onDeliveryFn) (SesEventProcessedResponse, error) {
	var data SesDeliveryEvent
	err := json.Unmarshal(event, &data)
	if err != nil {
		return SesEventProcessedResponse{}, err
	}
	err = callback(data)
	if err != nil {
		return SesEventProcessedResponse{}, err
	}
	return SesEventProcessedResponse{Status: "DELIVERED", Notification: data.Mail.MessageId}, nil
}

func processSesComplaintEvent(event []byte, callback onComplaintFn) (SesEventProcessedResponse, error) {
	var data SesComplaintEvent
	err := json.Unmarshal(event, &data)
	if err != nil {
		return SesEventProcessedResponse{}, err
	}
	err = callback(data)
	if err != nil {
		return SesEventProcessedResponse{}, err
	}
	return SesEventProcessedResponse{Status: "COMPLAINED", Notification: data.Mail.MessageId}, nil
}

func processSesBounceEvent(event []byte, callback onBounceFn) (SesEventProcessedResponse, error) {
	var data SesBounceEvent
	err := json.Unmarshal(event, &data)
	if err != nil {
		return SesEventProcessedResponse{}, err
	}
	err = callback(data)
	if err != nil {
		return SesEventProcessedResponse{}, err
	}
	return SesEventProcessedResponse{Status: "BOUNCED", Notification: data.Mail.MessageId}, nil
}
