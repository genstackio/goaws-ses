package goaws_ses

import (
	"time"
)

type SesEventProcessedResponse struct {
	Status       string `json:"status"`
	Notification string `json:"notification"`
}

type SesEventBasic struct {
	NotificationType string `json:"notificationType"`
}

type SesBounceEvent struct {
	NotificationType string      `json:"notificationType"`
	Bounce           BounceEvent `json:"Bounce"`
	Mail             MailEvent   `json:"mail"`
}
type BounceEvent struct {
	BounceType        string              `json:"bounceType"`
	ReportingMTA      string              `json:"reportingMTA"`
	BouncedRecipients []BouncedRecipients `json:"bouncedRecipients"`
	BounceSubType     string              `json:"bounceSubType"`
	Timestamp         time.Time           `json:"timestamp"`
	FeedbackId        string              `json:"feedbackId"`
	RemoteMtaIp       string              `json:"remoteMtaIp"`
}
type BouncedRecipients struct {
	EmailAddress   string `json:"emailAddress"`
	Status         string `json:"status"`
	Action         string `json:"action"`
	DiagnosticCode string `json:"diagnosticCode"`
}

type SesComplaintEvent struct {
	NotificationType string         `json:"notificationType"`
	Complaint        ComplaintEvent `json:"complaint"`
	Mail             MailEvent      `json:"mail"`
}
type ComplaintEvent struct {
	UserAgent             string                 `json:"userAgent"`
	ComplainedRecipients  []ComplainedRecipients `json:"complainedRecipients"`
	ComplaintFeedbackType string                 `json:"complaintFeedbackType"`
	ArrivalDate           string                 `json:"arrivalDate"`
	Timestamp             time.Time              `json:"timestamp"`
	FeedbackId            string                 `json:"feedbackId"`
}
type ComplainedRecipients struct {
	EmailAddress string `json:"emailAddress"`
}

type SesDeliveryEvent struct {
	NotificationType string        `json:"notificationType"`
	Delivery         DeliveryEvent `json:"delivery"`
	Mail             MailEvent     `json:"mail"`
}
type DeliveryEvent struct {
	Timestamp            time.Time `json:"timestamp"`
	Recipients           []string  `json:"recipients"`
	ProcessingTimeMillis int       `json:"processingTimeMillis"`
	ReportingMTA         string    `json:"reportingMTA"`
	SmtpResponse         string    `json:"smtpResponse"`
	RemoteMtaIp          string    `json:"remoteMtaIp"`
}

type MailEvent struct {
	Timestamp        time.Time `json:"timestamp"`
	MessageId        string    `json:"messageId"`
	Source           string    `json:"source"`
	SourceArn        string    `json:"sourceArn"`
	SourceIp         string    `json:"sourceIp"`
	SendingAccountId string    `json:"sendingAccountId"`
	CallerIdentity   string    `json:"callerIdentity"`
	Destination      []string  `json:"destination"`
}

type onDeliveryFn func(event SesDeliveryEvent) error
type onComplaintFn func(event SesComplaintEvent) error
type onBounceFn func(event SesBounceEvent) error
