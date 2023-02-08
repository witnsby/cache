# cache

The module could be imported via command ``go get github.com/witnsby/cache``

the module implemented 3 methods:
- `Set(key, value)` - allows to add `value` to cache via key `key`
- `Get(key string)` - retrieve value by `value`
- `Delete(key)` - delete key 

```go
# example of implemetation.
cache := cache.New()

cache.Set("userId", 42)

userId, err := cache.Get("userId")
if err != nil {
fmt.Println(err)
} else {
fmt.Printf("%v\n", userId)
}

cache.Delete("userId")
```
