package journal

import "time"

// Journal defines the journal model
type Journal struct {
	ID        int64
	Action    string
	CreatedAt time.Time
}
