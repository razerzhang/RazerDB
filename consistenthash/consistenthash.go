package consistenthash

type Hash func(data []byte)uint32

type Ring struct {
	hash  func(data []byte)uint32
	replicas int  //the replicas of virtual nodes
	vnodes []int
	hashmap map[int]string
}

//New build a consistenthashmap instance
func New(replicas int, fn Hash)  {

}