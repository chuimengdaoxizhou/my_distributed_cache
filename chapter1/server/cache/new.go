package cache

import (
	"log"
)

func New(typ string) Cache {
	var c Cache
	if typ == "inmemory" {
		c = newInMemoryCache()
	}
	if c == nil {
		panic("unknown cache type" + typ)
	}
	log.Println(typ, "reday to serve")
	return c
}
