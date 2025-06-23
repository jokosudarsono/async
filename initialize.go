package async

import (
	"sync"
)

var (
	once sync.Once

	core   *asyncCore
	config *asyncConfig
)

type asyncCore struct {
	copyContext []ContextKey
}

type AsyncOption func(*asyncCore)

func Init(options ...AsyncOption) {
	core = new(asyncCore)

	for _, opt := range options {
		opt(core)
	}
}

func SetCopyContext(keys ...ContextKey) AsyncOption {
	return func(ac *asyncCore) {
		ac.copyContext = keys
	}
}

type asyncConfig struct {
	copyContext []ContextKey
}

func GetConfig() *asyncConfig {
	once.Do(func() {
		config = new(asyncConfig)
		config.copyContext = []ContextKey{}

		if core != nil {
			config.copyContext = core.copyContext
		}
	})

	return config
}

func (c *asyncConfig) GetCopyContext() []ContextKey {
	return c.copyContext
}
