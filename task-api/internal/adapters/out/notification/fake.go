package notification

import (
	"context"
	"log"
	"task-api/internal/ports/out"
)

var _ out.NotificationSender = &FakeNotificationSender{}

type FakeNotificationSender struct{}

func NewFakeNotificationSender() *FakeNotificationSender {
	return &FakeNotificationSender{}
}

func (f FakeNotificationSender) SendTo(ctx context.Context, message string, to string) error {
	log.Printf("sending the %s notification to %s \n", message, to)
	return nil
}
