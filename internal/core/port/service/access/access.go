package access

import "context"

// Service defines the access service.
type Service interface {
	Check(ctx context.Context, endpoint string) error
}
