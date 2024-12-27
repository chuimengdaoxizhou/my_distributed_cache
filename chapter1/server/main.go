package main

import (
	"server/cache"
	"server/http"
)

func main() {
	c := cache.New("inmemory")
	http.New(c).Listen()
}
