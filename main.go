package main

import (
	"fmt"
	"lru-cache/dt"
)

// we need hashMap for to check if referred key exists in cache or not
// and double linked list for to keep track of the freshness of values
// last referred one goes to the head of the double linked list
// when we reach the limit of cache size, remove the least recently used one ( showed by tail )

func main() {
	lru := dt.NewLRUCache(3)

	lru.AddEntry("Born Year", 1998)
	lru.AddEntry("Favorite Number", 52)

	fmt.Println(lru.CheckCache("Born Year"))

	fmt.Print(lru.GetCount())

}
