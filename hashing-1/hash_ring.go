package hashing

import (
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
	"log"
	"sort"
	"sync"
	"slices"
)

var (
	ErrNoConnectedNodes = errors.New("no connected nodes available")
	ErrNodeExists = errors.New("node already exists")
	ErrNodeNotFound = errors.New("node not found")
	ErrInHashingKey = errors.New("error in hashing the key")
)

type ICacheNode interface {
	GetIdentifier() string
}

type hashRingConfig struct {
	HashFunction func() hash.Hash64
	EnableLogs bool
}

type HashRingConfigFn func(*hashRingConfig)

func SetHashFunction(f func() hash.Hash64) HashRingConfigFn {
	return func (config *hashRingConfig) {
		config.HashFunction = f
	}
}

func EnableVerboseLogs(enabled bool) HashRingConfigFn {
	return func (config *hashRingConfig) {
		config.EnableLogs = enabled
	}
}

type HashRing struct {
	mu sync.RWMutex
	config hashRingConfig
	nodes sync.Map
	sortedKeysOfNodes []uint64
}

func InitHashRing(opts ...HashRingConfigFn) *HashRing {
	config := &hashRingConfig{
		HashFunction: fnv.New64a,
		EnableLogs: false,
	}

	for _, opt := range opts {
		opt(config)
	}

	return &HashRing{
		config: *config,
		sortedKeysOfNodes: make([]uint64,0),
	}
}

func(h *HashRing) AddServer(node ICacheNode) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	hashValue, err := h.generateHash(node.GetIdentifier())
	if err != nil {
		return fmt.Errorf("%w : %s",ErrNodeExists, node.GetIdentifier())
	}

	h.nodes.Store(hashValue,node)
	h.sortedKeysOfNodes = append(h.sortedKeysOfNodes, hashValue)

	slices.Sort(h.sortedKeysOfNodes) //sorting hash keys for binary search

	if h.config.EnableLogs {
		log.Printf("[HashRing] Added node : %s (hash: %d)",node.GetIdentifier(),hashValue)
	}

	return nil

}

func (h *HashRing) GetServer(key string) (ICacheNode, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	hashValue, err := h.generateHash(key)
	if err != nil {
		return nil, fmt.Errorf("%w : %s",ErrInHashingKey, key)
	}

	//performs a binary search on sortedKeyOfNodes
	index,err := h.search(hashValue)
	if err != nil {
		return nil, err
	}

	nodeHash := h.sortedKeysOfNodes[index]
	if node, ok := h.nodes.Load(nodeHash); ok {
		if h.config.EnableLogs{
			log.Printf("[HashRing] Key '%s' (hash: %d) mapped to node (hash:%d)",key,hashValue,nodeHash)
		}
		return node.(ICacheNode), nil
	}

	return nil, fmt.Errorf("%w: no node found for key %s", ErrNodeNotFound, key)
}

func (h *HashRing) RemoveServer(node ICacheNode) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	hashValue, err := h.generateHash(node.GetIdentifier())
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInHashingKey, node.GetIdentifier())
	}

	if _, found := h.nodes.LoadAndDelete(hashValue); !found {
		return fmt.Errorf("%w: %s", ErrNodeNotFound, node.GetIdentifier())
	}

	index, err := h.search(hashValue)
	if err != nil {
		return err
	}

	h.sortedKeysOfNodes = append(h.sortedKeysOfNodes[:index], h.sortedKeysOfNodes[index+1:]...)

	if h.config.EnableLogs {
		log.Printf("[HashRing] Removed node: %s (hash: %d)",node.GetIdentifier(), hashValue)
	}

	return nil
}

func (h *HashRing) search(key uint64) (int, error) {
	if len(h.sortedKeysOfNodes) == 0 {
		return -1, ErrNoConnectedNodes
	}

	index := sort.Search(len(h.sortedKeysOfNodes),func(i int) bool {
		return h.sortedKeysOfNodes[i] >= key
	})

	if index == len(h.sortedKeysOfNodes){
		index = 0
	}

	return index, nil
}

func (h *HashRing) generateHash(key string) (uint64, error) {
	hash := h.config.HashFunction()
	if _, err := hash.Write([]byte(key)); err != nil {
		return 0,err
	}

	return hash.Sum64(), nil
}

