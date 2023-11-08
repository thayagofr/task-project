package out

import (
	"context"
)

var _ NotificationSender = &NotificationSenderMock{}

type NotificationSenderMock struct {
	MockedSendTo func(ctx context.Context, message string, to string) error
}

func NewNotificationSenderMock() *NotificationSenderMock {
	return &NotificationSenderMock{}
}

func (mock *NotificationSenderMock) SendTo(ctx context.Context, message string, to string) error {
	if mock.MockedSendTo != nil {
		return mock.MockedSendTo(ctx, message, to)
	}
	return nil
}
