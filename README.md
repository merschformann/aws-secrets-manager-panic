# aws-secrets-manager-panic

Minimal example of AWS secrets manager go package panicking when not using the defaults.

## Steps to reproduce

```bash
$ go run main.go
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x10 pc=0x72bb7d]

goroutine 1 [running]:
github.com/aws/aws-secretsmanager-caching-go/secretcache.(*secretCacheItem).getSecretValue(0x0, {0x7b976e?, 0x3?})
        github.com/aws/aws-secretsmanager-caching-go@v1.1.0/secretcache/cacheItem.go:168 +0x9d
github.com/aws/aws-secretsmanager-caching-go/secretcache.(*Cache).GetSecretStringWithStage(0xc00018fef0?, {0x7b76e4?, 0x43cec5?}, {0x7b976e, 0xa})
        github.com/aws/aws-secretsmanager-caching-go@v1.1.0/secretcache/cache.go:94 +0x3b
github.com/aws/aws-secretsmanager-caching-go/secretcache.(*Cache).GetSecretString(...)
        github.com/aws/aws-secretsmanager-caching-go@v1.1.0/secretcache/cache.go:86
main.main()
        aws-secrets-manager-panic/main.go:22 +0xdd
exit status 2
```

## Expected behaviour

Expected to receive an error from either `secretcache.New` for an incomplete configuration or `cache.GetSecretString` for an invalid secret ID.
