package lru

import "testing"

var lru *Cache

func TestAdd(t *testing.T)  {
	lru = Init()
	lru.Add("key1","1234")
	if len(lru.cache) != 1{
		t.Error("cache add fail")
	}
}

func TestGet(t *testing.T)  {
	if v,ok := lru.Get("key1"); !ok ||v.(string) != "1234"{
		t.Error("cache get fail")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}
}