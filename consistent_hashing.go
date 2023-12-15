package main

import (
	"fmt"
	"hash/fnv"
	"sort"
)

type ConsistentHash struct {
	hashFunction     HashFunction
	numberOfReplicas int
	circle           map[uint32]interface{}
	sortedKeys       []uint32
}

type HashFunction interface {
	hash(s string) uint32
}

type Hasher struct{}

func (h Hasher) hash(s string) uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(s))
	return hash.Sum32()
}

func NewConsistentHash(hashFunction HashFunction, numberOfReplicas int, nodes []interface{}) *ConsistentHash {
	ch := &ConsistentHash{
		hashFunction:     hashFunction,
		numberOfReplicas: numberOfReplicas,
		circle:           make(map[uint32]interface{}),
	}

	for _, node := range nodes {
		ch.add(node)
	}

	return ch
}

func (ch *ConsistentHash) add(node interface{}) {
	for i := 0; i < ch.numberOfReplicas; i++ {
		hash := ch.hashFunction.hash(fmt.Sprintf("%v%d", node, i))
		ch.circle[hash] = node
		ch.sortedKeys = append(ch.sortedKeys, hash)
	}
	sort.Slice(ch.sortedKeys, func(i, j int) bool {
		return ch.sortedKeys[i] < ch.sortedKeys[j]
	})
}

func (ch *ConsistentHash) remove(node interface{}) {
	for i := 0; i < ch.numberOfReplicas; i++ {
		hash := ch.hashFunction.hash(fmt.Sprintf("%v%d", node, i))
		delete(ch.circle, hash)
		ch.updateSortedKeys(hash)
	}
}

func (ch *ConsistentHash) updateSortedKeys(hash uint32) {
	for i, key := range ch.sortedKeys {
		if key == hash {
			ch.sortedKeys = append(ch.sortedKeys[:i], ch.sortedKeys[i+1:]...)
			break
		}
	}
}

func (ch *ConsistentHash) get(key interface{}) interface{} {
	if len(ch.circle) == 0 {
		return nil
	}

	hash := ch.hashFunction.hash(fmt.Sprintf("%v", key))
	if _, ok := ch.circle[hash]; !ok {
		pos := sort.Search(len(ch.sortedKeys), func(i int) bool {
			return ch.sortedKeys[i] >= hash
		})
		if pos == len(ch.sortedKeys) {
			hash = ch.sortedKeys[0]
		} else {
			hash = ch.sortedKeys[pos]
		}
	}
	return ch.circle[hash]
}

func main() {
	nodes := []interface{}{"Node1", "Node2", "Node3"}

	hasher := Hasher{}
	consistentHash := NewConsistentHash(hasher, 3, nodes)

	key := "Key1"
	node := consistentHash.get(key)
	fmt.Printf("Key %s is mapped to node %v\n", key, node)
}
