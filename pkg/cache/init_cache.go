package cache

import (
	"sync"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

var (
	authCache     *gocache.Cache
	authCacheOnce sync.Once
)

func GetAuthCache() *gocache.Cache {
	authCacheOnce.Do(func() {
		authCache = gocache.New(5*time.Minute, 10*time.Minute)
	})

	return authCache
}
