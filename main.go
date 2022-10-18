package main

import (
	"fmt"

	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
)

func main() {
	config := secretcache.CacheConfig{
		// Uncomment the defaults below to make the example work again
		// MaxCacheSize: secretcache.DefaultMaxCacheSize,
		// VersionStage: secretcache.DefaultVersionStage,
		CacheItemTTL: 30 * 1000 * 1000 * 1000,
	}
	secrets, err := secretcache.New(
		func(c *secretcache.Cache) { c.CacheConfig = config },
	)
	if err != nil {
		panic(fmt.Errorf("error creating secret cache: %w", err))
	}
	secrets.GetSecretString("foo")
}
