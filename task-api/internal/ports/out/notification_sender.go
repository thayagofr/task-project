package out

import (
	"context"
)

type NotificationSender interface {
	SendTo(ctx context.Context, message string, to string) error
}
