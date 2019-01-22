// Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.

// get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
// put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

// Follow up:
// Could you do both operations in O(1) time complexity?

// Example:

// LRUCache cache = new LRUCache( 2 /* capacity */ );

// cache.put(1, 1);
// cache.put(2, 2);
// cache.get(1);       // returns 1
// cache.put(3, 3);    // evicts key 2
// cache.get(2);       // returns -1 (not found)
// cache.put(4, 4);    // evicts key 1
// cache.get(1);       // returns -1 (not found)
// cache.get(3);       // returns 3
// cache.get(4);       // returns 4
// https://leetcode.com/problems/lru-cache/

// ["LRUCache","put","put","get","put","get","put","get","get","get"]
// [[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]

package main

import (
	"fmt"
)

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	fmt.Println("obj", obj)

	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	fmt.Println("obj", obj)

	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println("obj", obj)

	fmt.Println(obj.Get(3))
	fmt.Println("obj", obj)

	obj.Get(4)

	fmt.Println("obj", obj)
}

type node struct {
	prev *node
	next *node
	key  int
	val  int
}

type LRUCache struct {
	data     map[int]*node
	capacity int
	head     *node
	tail     *node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		data:     make(map[int]*node),
	}
}

func (this *LRUCache) remove(n *node) {
	if n == this.head {
		this.head = n.next
	}
	if n == this.tail {
		this.tail = n.prev
	}

	if n.next != nil {
		n.next.prev = n.prev
	}

	if n.prev != nil {
		n.prev.next = n.next
	}
}

func (this *LRUCache) add(n *node) {
	n.prev = nil
	n.next = this.head
	if n.next != nil {
		n.next.prev = n
	}
	this.head = n
	if this.tail == nil {
		this.tail = n
	}

}

func (this *LRUCache) Get(key int) int {
	n, ok := this.data[key]
	if !ok {
		return -1
	}
	this.remove(n)
	this.add(n)
	return n.val
}

func (this *LRUCache) Put(key int, value int) {
	n, ok := this.data[key]

	if !ok {
		n = &node{key: key, val: value}
		this.data[key] = n
	} else {
		n.val = value
		this.remove(n)
	}

	this.add(n)
	if len(this.data) > this.capacity {
		n = this.tail
		if n != nil {
			this.remove(n)
			delete(this.data, n.key)
		}
	}
}
