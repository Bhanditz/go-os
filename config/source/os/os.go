package os

/*
	OS source uses the Micro config-srv
*/

import (
	"github.com/pydio/go-os/config"
)

func NewSource(opts ...config.SourceOption) config.Source {
	return config.NewSource(opts...)
}
