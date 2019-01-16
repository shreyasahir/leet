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

package main

import (
	"fmt"
)

func main() {
	obj := Constructor(2);
	param_1 := obj.Get(key);
	obj.Put(1,1);
	obj.Put(2,3)
	obj.get(1)
	obj.Put(3,4)
	fmt.Println("obj",obj)
}

type LRUCache struct {
	data     map[int]int
	frequency map[int]int
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity
	}
}

func (this *LRUCache) Get(key int) int {
	if this.data == nil {
		this.data = make(map[int]int)
	} else if val,ok := this.data[key];ok {
		this.frequency[key]++
		return this.data[key]
	} else{
		-1
	}

}

func (this *LRUCache) Put(key int, value int) {
	if this.Get(key) == -1{
		if len(this.data) == this.capacity {
			deleteKey = this.leastUsedKey()
			delete(this.data,deleteKey.Key)
		}
		this.data[key] = value
	}
}

func (this *LRUCache) leastUsedKey() {
	p := make(PairList, len(this.data))
	i := 0
	for k, v := range this.data {
	p[i] = Pair{k, v}
	}
	sort.Sort(p)
	return p[0]
}

type PairList []Pair
type Pair struct {
	Key string
	Value int
}
func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
