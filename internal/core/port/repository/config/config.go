package config

import (
	ecfg "github.com/pillarion/practice-auth/internal/core/entity/config"
)

// Config defines the config interface.
type Config interface {
	GetConfig() (*ecfg.Config, error)
}
