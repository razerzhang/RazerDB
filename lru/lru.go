package lru

import "container/list"

type Cache struct {


	ll *list.List
	cache map[interface{}]*list.Element
}

type Key interface {}

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
func (c *Cache)Add()  {
	if c.cache == nil{
		c.cache = make(map[interface{}]*list.Element)
		c.ll = list.New()
	}
}