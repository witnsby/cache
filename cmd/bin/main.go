package main

import (
	pkgCache "cache/pkg/cache"
	"fmt"
)

func main() {
	cache := pkgCache.New()
	cache.Set("userId", 42)
	userId, ok := cache.Get("userId")
	if ok != nil {
		fmt.Println(ok)
	} else {
		fmt.Printf("%v\n", userId)
	}

	cache.Delete("userId")
	userId1, ok := cache.Get("userId")
	if ok != nil {
		fmt.Println(ok)
	} else {
		fmt.Printf("%v\n", userId1)
	}
	cache.Set("userId2", 1111)
	fmt.Printf("%v\n", cache)
}
