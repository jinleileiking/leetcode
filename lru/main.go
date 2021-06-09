package main

import (
	"container/list"

	"github.com/davecgh/go-spew/spew"
)

type LRU struct {
	Hash map[string]string
	List *list.List
	Len  int
}

func (lru *LRU) Get(key string) string {
	if v, ok := lru.Hash[key]; ok {
		// put key to back

		for e := lru.List.Front(); e != nil; e = e.Next() {
			if e.Value == key {
				lru.List.MoveToBack(e)
				break
			}
		}
		return v
	}

	return ""
}

func (lru *LRU) New(length int) {
	m := make(map[string]string)
	(*lru).Hash = m
	(*lru).List = list.New()
	(*lru).Len = length
}

func (lru *LRU) Put(key, val string) {

	// check if delete

	if len(lru.Hash) == lru.Len {

		// find head
		h := lru.List.Front()

		// delete map
		delete(lru.Hash, h.Value.(string))

		// delete list
		lru.List.Remove(h)
	}

	lru.Hash[key] = val
	lru.List.PushBack(key)

	return
}

func main() {

	l := LRU{}

	l.New(10)

	l.Put("k1", "v1")
	l.Put("k2", "v2")
	l.Put("k3", "v3")
	l.Put("k4", "v4")
	l.Put("k5", "v5")
	l.Put("k6", "v6")
	l.Put("k7", "v7")
	l.Put("k8", "v8")
	l.Put("k9", "v9")
	l.Put("k10", "v10")
	l.Put("k11", "v11")
	l.Get("k5")
	l.Put("k12", "v12")

	spew.Dump(l)
}
