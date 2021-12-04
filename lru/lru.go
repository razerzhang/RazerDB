package lru

import (
	"container/list"

)

type Cache struct {
	Maxlength int

	ll *list.List
	cache map[interface{}]*list.Element

	OnEvicted func(key Key,value interface{})
}

type Key interface {}

//the store value
type entry struct {
	key Key
	value interface{}
}

func Init()*Cache  {
	return &Cache{
		ll: list.New(),
		cache:  make(map[interface{}]*list.Element),
	}
}

//Add add a value to the Cache
func (c *Cache)Add(key Key,value interface{})  {
	if c.cache == nil{
		c.cache = make(map[interface{}]*list.Element)
		c.ll = list.New()
	}
	//if key had been exist in the cache,move to front
	if vl,exist := c.cache[key];exist{
		c.ll.MoveToFront(vl)
		vl.Value.(*entry).value =value
		return
	}
	el := c.ll.PushFront(&entry{key,value})
	c.cache[key] = el
	//todo  remove the oldest data from list
	if c.Maxlength > 0 && c.ll.Len() > c.Maxlength{
		c.DelOldest()
	}
}

func (c *Cache)Del(key Key)  {
	//remove value in cache
	if c.cache == nil{
		return
	}

	//todo remove data from list and cach
	if ele,hit := c.cache[key];hit{
		c.delElement(ele)
	}
}

//delete oldest item
func (c *Cache)DelOldest()  {
	if c.cache == nil{
		return
	}
	ele := c.ll.Back()
	if ele != nil{
		c.delElement(ele)
	}
}

func (c *Cache)delElement(e *list.Element)  {
	//delete list
	c.ll.Remove(e)
	//delete  cach
	kv := e.Value.(*entry)
	delete(c.cache,e.Value.(*entry).key)
	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}
}