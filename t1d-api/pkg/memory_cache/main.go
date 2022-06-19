package memory_cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

// Memory Cache Singleton - Used to store key / values into container memory
func GetInstance() *cache.Cache {
	c := cache.New(5*time.Minute, 10*time.Minute)
	c.Set("api", "t1d", cache.DefaultExpiration)

	return c
}
