package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte)uint32

type Ring struct {
	hash  func(data []byte)uint32
	replicas int  //the replicas of virtual nodes
	vnodes []int  //hash ring
	hashmap map[int]string
}

//New build a consistenthashmap instance
func New(replicas int, fn Hash) *Ring {
	r := &Ring{
		replicas: replicas,
		hash:     fn,
		hashmap:  make(map[int]string),
	}
	if r.hash == nil {
		r.hash = crc32.ChecksumIEEE
	}
	return r
}

func (r *Ring)Add(keys ...string)  {
	for _, key := range keys {
		for i := 0; i < r.replicas; i++ {
			hash := int(r.hash([]byte(strconv.Itoa(i) + key)))
			r.vnodes = append(r.vnodes, hash)
			r.hashmap[hash] = key
		}
	}
	sort.Ints(r.vnodes)
}

func (r *Ring)Get(key string) string {
	hash := int(r.hash([]byte(key)))

	// Binary search for appropriate replica.
	idx := sort.Search(len(r.vnodes), func(i int) bool {
		return r.vnodes[i] >= hash
	})

	// Means we have cycled back to the first replica.
	if idx == len(r.vnodes){
		idx = 0
	}

	return r.hashmap[r.vnodes[idx]]
}