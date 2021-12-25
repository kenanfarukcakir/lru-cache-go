package dt // data structures

import (
	"errors" // can use a (double) linked list to implement a queue
)

type CacheKey string
type CacheValue int

type DoubleLinkedListNode struct {
	key  CacheKey
	val  CacheValue
	prev *DoubleLinkedListNode
	next *DoubleLinkedListNode
}

type DoubleLinkedList struct {
	head *DoubleLinkedListNode
	tail *DoubleLinkedListNode
}

type LRUCache struct {
	dll     *DoubleLinkedList
	hashMap map[CacheKey]*DoubleLinkedListNode
	len     int
	maxSize int
}

func NewDoubleLinkedList() *DoubleLinkedList {
	dLL := DoubleLinkedList{}
	return &dLL
}

func NewLRUCache(maxSize int) LRUCache {
	lru := LRUCache{}

	lru.dll = NewDoubleLinkedList()
	lru.hashMap = make(map[CacheKey]*DoubleLinkedListNode, 0)
	lru.len = 0
	lru.maxSize = maxSize

	return lru
}

func (lru *LRUCache) AddEntry(key CacheKey, val CacheValue) {
	node := DoubleLinkedListNode{key: key, val: val}
	if lru.dll.head == nil {

		lru.dll.head = &node
		lru.dll.tail = &node

		lru.hashMap[key] = &node

		lru.len = lru.len + 1
		return
	}

	if lru.len == lru.maxSize {
		lru.RemoveLru()
	}

	oldHead := lru.dll.head

	node.next = oldHead
	oldHead.prev = &node
	lru.dll.head = &node

	lru.hashMap[key] = &node

	lru.len++

}

func (lru LRUCache) GetCount() int {
	return lru.len
}

func (lru *LRUCache) RemoveLru() {
	oldTail := lru.dll.tail
	delete(lru.hashMap, oldTail.key)

	prev := oldTail.prev
	prev.next = nil

	lru.dll.tail = prev
	lru.len--

}

// move last accessed to head of dll
func (lru LRUCache) CheckCache(key CacheKey) (CacheValue, error) {

	if node, ok := lru.hashMap[key]; ok {
		if node == lru.dll.head { // already head
			return node.val, nil
		} else if node == lru.dll.tail { // move tail to head
			tailBefore := lru.dll.tail.prev
			lru.dll.tail = tailBefore

			node.next = lru.dll.head
			lru.dll.head.prev = node
			node.prev = nil

			lru.dll.head = node

			return node.val, nil
		} else { // node in middle
			prev, next := node.prev, node.next

			prev.next = next
			next.prev = prev

			oldHead := lru.dll.head

			oldHead.prev = node
			node.next = oldHead
			node.prev = nil

			lru.dll.head = node
		}
		return node.val, nil
	} else {
		return 0, errors.New("not found")
	}

}
