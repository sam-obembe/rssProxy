package rssCache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var myCache = cache.New(24*time.Hour, 10*time.Minute)

func GetCache() *cache.Cache {
	return myCache
}
