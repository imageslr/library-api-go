package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	// Instance -
	Instance = cache.New(5*time.Minute, 10*time.Minute)
	// DefaultExpiration -
	DefaultExpiration = cache.DefaultExpiration
)
