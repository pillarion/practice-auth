package config

import (
	ecfg "github.com/pillarion/practice-auth/internal/core/entity/config"
)

// Config defines the config interface.
//
//go:generate minimock -o mock/ -s "_minimock.go"
type Config interface {
	Get() (*ecfg.Config, error)
}
