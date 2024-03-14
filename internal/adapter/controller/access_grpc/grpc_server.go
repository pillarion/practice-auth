package access

import (
	"github.com/pillarion/practice-auth/internal/core/port/service/access"
	desc "github.com/pillarion/practice-auth/pkg/access_v1"
)

type server struct {
	desc.UnimplementedAccessV1Server
	accessService access.Service
}

// NewServer returns a new server.
func NewServer(as access.Service) *server {
	return &server{
		accessService: as,
	}
}
