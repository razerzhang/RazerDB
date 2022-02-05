package cluster

import (
	"RazerCache/consistenthash"
	"github.com/toolkits/pkg/str"
)

type ClusterNode struct {
	Addrs []string `json:"addrs"`
}

//hashring of the service node
type ClusterRing struct {
	Cring *consistenthash.ConsistentHashRing
}

func (ring *ClusterRing)Init()  {
	ring.Cring = consistenthash.NewConsistentHashRing(Razer.replicates,str.KeysOfMap(Razer.Section.Cluster))
}

func (ring *ClusterRing)LoadBalance(key string)string  {
	cluster,err := ring.Cring.GetNode(key)
	if err != nil{
		err.Error()
	}
	return cluster
}
