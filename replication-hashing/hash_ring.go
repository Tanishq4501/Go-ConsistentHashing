package replicationhashing

import (
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
	"log"
	"slices"
	"sort"
	"sync"
)

var (
	ErrNoConnectedNodes = errors.New("no connected nodes available")
	ErrNodeExists       = errors.New("node already exists")
	ErrNodeNotFound     = errors.New("node not found")
	ErrInHashingKey     = errors.New("error in hashing the key")
)

type ICacheNode interface {
	GetIdentifier() string
}

type hashRingConfig struct {
	VirtualNodes int
	HashFunction func() hash.Hash64
	EnableLogs   bool
}

type HashRingConfigFn func(*hashRingConfig)

func SetVirtualNodes(count int) HashRingConfigFn {
	return func(cfg *hashRingConfig) {
		cfg.VirtualNodes = count
	}
}

func SetHashFunction(f func() hash.Hash64) HashRingConfigFn {
	return func(config *hashRingConfig) {
		config.HashFunction = f
	}
}

func EnableVerboseLogs(enabled bool) HashRingConfigFn {
	return func(config *hashRingConfig) {
		config.EnableLogs = enabled
	}
}

type HashRing struct {
	mu                sync.RWMutex
	config            hashRingConfig
	hostMap           sync.Map // nodeId -> timeAdded
	vNodeMap          sync.Map // hash -> node
	sortedKeysOfNodes []uint64 // sorted hash values (includes virtual nodes)
}

func InitHashRing(opts ...HashRingConfigFn) *HashRing {
	config := &hashRingConfig{
		HashFunction: fnv.New64a,
		VirtualNodes: 3,
	}

	for _, opt := range opts {
		opt(config)
	}

	return &HashRing{
		config:            *config,
		sortedKeysOfNodes: make([]uint64, 0),
	}
}

func (h *HashRing) AddServer(node ICacheNode) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	nodeId := node.GetIdentifier()
	if _, exists := h.hostMap.Load(nodeId); exists {
		return fmt.Errorf("%w : %s", ErrNodeExists, nodeId)
	}

	virtualKeys := make([]uint64, 0, h.config.VirtualNodes)
	for i := 0; i < h.config.VirtualNodes; i++ {
		vNodeId := fmt.Sprintf("%s_%d", nodeId, i)
		hash, err := h.generateHash(vNodeId)
		if err != nil {
			return fmt.Errorf("%w for virtual node %s", ErrInHashingKey, vNodeId)
		}
		h.vNodeMap.Store(hash, node)
		virtualKeys = append(virtualKeys, hash)

		if h.config.EnableLogs {
			log.Printf("[HashRing] Added virtual node %s -> hash %d", vNodeId, hash)
		}
	}

	h.hostMap.Store(nodeId, struct{}{})
	h.sortedKeysOfNodes = append(h.sortedKeysOfNodes, virtualKeys...)
	slices.Sort(h.sortedKeysOfNodes)

	if h.config.EnableLogs {
		log.Printf("[HashRing] Node %s added with %d virtual nodes", nodeId, h.config.VirtualNodes)
	}

	return nil

}

func (h *HashRing) RemoveServer(node ICacheNode) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	nodeId := node.GetIdentifier()
	if _, exists := h.hostMap.Load(nodeId); exists {
		return fmt.Errorf("%w : %s", ErrNodeExists, nodeId)
	}

	for i := 0; i < h.config.VirtualNodes; i++ {
		vNodeId := fmt.Sprintf("%s_%d", nodeId, i)
		hash, err := h.generateHash(vNodeId)
		if err != nil {
			return fmt.Errorf("%w for virtual node %s", ErrInHashingKey, vNodeId)
		}
		h.vNodeMap.Delete(hash)

		index := slices.Index(h.sortedKeysOfNodes, hash)
		if index >= 0 {
			h.sortedKeysOfNodes = append(h.sortedKeysOfNodes[:index], h.sortedKeysOfNodes[index+1:]...)
		}

		if h.config.EnableLogs {
			log.Printf("[HashRing] Removed node: %s (hash: %d)", vNodeId, hash)
		}

	}

	h.hostMap.Delete(nodeId)
	return nil
}

func (h *HashRing) GetServer(key string) (ICacheNode, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	hashValue, err := h.generateHash(key)
	if err != nil {
		return nil, fmt.Errorf("%w : %s", ErrInHashingKey, key)
	}

	//performs a binary search on sortedKeyOfNodes
	index, err := h.search(hashValue)
	if err != nil {
		return nil, err
	}

	nodeHash := h.sortedKeysOfNodes[index]
	if node, ok := h.vNodeMap.Load(nodeHash); ok {
		if h.config.EnableLogs {
			log.Printf("[HashRing] Key '%s' (hash: %d) mapped to node (hash:%d)", key, hashValue, nodeHash)
		}
		return node.(ICacheNode), nil
	}

	return nil, fmt.Errorf("%w: no node found for key %s", ErrNodeNotFound, key)
}

func (h *HashRing) search(key uint64) (int, error) {
	if len(h.sortedKeysOfNodes) == 0 {
		return -1, ErrNoConnectedNodes
	}

	index := sort.Search(len(h.sortedKeysOfNodes), func(i int) bool {
		return h.sortedKeysOfNodes[i] >= key
	})

	if index == len(h.sortedKeysOfNodes) {
		index = 0
	}

	return index, nil
}

func (h *HashRing) generateHash(key string) (uint64, error) {
	hash := h.config.HashFunction()
	if _, err := hash.Write([]byte(key)); err != nil {
		return 0, err
	}

	return hash.Sum64(), nil
}
