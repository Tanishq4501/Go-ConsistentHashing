package main

import (
	"fmt"
	"log"

	hashing1 "github.com/Tanishq4501/go-hash/hashing-1"
	redundanthashing "github.com/Tanishq4501/go-hash/redundant-hashing"
	replicationhashing "github.com/Tanishq4501/go-hash/replication-hashing"
)

// CacheNode implements ICacheNode interface for all three implementations
type CacheNode struct {
	ID string
}

func (n *CacheNode) GetIdentifier() string {
	return n.ID
}

func main() {
	fmt.Println("=== Consistent Hashing Demo ===\n")

	// Create test nodes
	nodes := []*CacheNode{
		{ID: "server-1"},
		{ID: "server-2"},
		{ID: "server-3"},
	}

	testKeys := []string{"user:1001", "user:1002", "user:1003", "cache:session:abc", "cache:session:xyz"}

	// Demo 1: Basic Hashing (hashing-1)
	fmt.Println("--- 1. Basic Consistent Hashing ---")
	demoBasicHashing(nodes, testKeys)

	// Demo 2: Replication Hashing with Virtual Nodes (replication-hashing)
	fmt.Println("\n--- 2. Consistent Hashing with Virtual Nodes ---")
	demoReplicationHashing(nodes, testKeys)

	// Demo 3: Redundant Hashing with Replication Factor (redundant-hashing)
	fmt.Println("\n--- 3. Consistent Hashing with Redundancy ---")
	demoRedundantHashing(nodes, testKeys)
}

func demoBasicHashing(nodes []*CacheNode, testKeys []string) {
	ring := hashing1.InitHashRing(hashing1.EnableVerboseLogs(false))

	// Add servers
	fmt.Println("Adding servers...")
	for _, node := range nodes {
		if err := ring.AddServer(node); err != nil {
			log.Printf("Error adding node %s: %v", node.ID, err)
		} else {
			fmt.Printf("✓ Added: %s\n", node.ID)
		}
	}

	// Map keys to servers
	fmt.Println("\nMapping keys to servers:")
	for _, key := range testKeys {
		node, err := ring.GetServer(key)
		if err != nil {
			log.Printf("Error getting server for %s: %v", key, err)
			continue
		}
		fmt.Printf("  %s -> %s\n", key, node.GetIdentifier())
	}

	// Remove a server and show remapping
	fmt.Println("\nRemoving server-2...")
	if err := ring.RemoveServer(nodes[1]); err != nil {
		log.Printf("Error removing node: %v", err)
	} else {
		fmt.Println("✓ Removed: server-2")
	}

	fmt.Println("\nKeys after removing server-2:")
	for _, key := range testKeys {
		node, err := ring.GetServer(key)
		if err != nil {
			log.Printf("Error getting server for %s: %v", key, err)
			continue
		}
		fmt.Printf("  %s -> %s\n", key, node.GetIdentifier())
	}
}

func demoReplicationHashing(nodes []*CacheNode, testKeys []string) {
	ring := replicationhashing.InitHashRing(
		replicationhashing.SetVirtualNodes(3),
		replicationhashing.EnableVerboseLogs(false),
	)

	// Add servers
	fmt.Println("Adding servers with 3 virtual nodes each...")
	for _, node := range nodes {
		if err := ring.AddServer(node); err != nil {
			log.Printf("Error adding node %s: %v", node.ID, err)
		} else {
			fmt.Printf("✓ Added: %s (with 3 virtual nodes)\n", node.ID)
		}
	}

	// Map keys to servers
	fmt.Println("\nMapping keys to servers:")
	for _, key := range testKeys {
		node, err := ring.GetServer(key)
		if err != nil {
			log.Printf("Error getting server for %s: %v", key, err)
			continue
		}
		fmt.Printf("  %s -> %s\n", key, node.GetIdentifier())
	}

	// Remove a server and show remapping
	fmt.Println("\nRemoving server-2...")
	if err := ring.RemoveServer(nodes[1]); err != nil {
		log.Printf("Error removing node: %v", err)
	} else {
		fmt.Println("✓ Removed: server-2")
	}

	fmt.Println("\nKeys after removing server-2:")
	for _, key := range testKeys {
		node, err := ring.GetServer(key)
		if err != nil {
			log.Printf("Error getting server for %s: %v", key, err)
			continue
		}
		fmt.Printf("  %s -> %s\n", key, node.GetIdentifier())
	}
}

func demoRedundantHashing(nodes []*CacheNode, testKeys []string) {
	ring := redundanthashing.InitHashRing(
		redundanthashing.SetVirtualNodes(3),
		redundanthashing.SetReplicationFactor(2),
		redundanthashing.EnableVerboseLogs(false),
	)

	// Add servers
	fmt.Println("Adding servers with 3 virtual nodes and replication factor 2...")
	for _, node := range nodes {
		if err := ring.AddNode(node); err != nil {
			log.Printf("Error adding node %s: %v", node.ID, err)
		} else {
			fmt.Printf("✓ Added: %s\n", node.ID)
		}
	}

	// Map keys to servers (with replication)
	fmt.Println("\nMapping keys to servers (with 2 replicas each):")
	for _, key := range testKeys {
		nodes, err := ring.GetNodesForKey(key)
		if err != nil {
			log.Printf("Error getting nodes for %s: %v", key, err)
			continue
		}
		fmt.Printf("  %s -> [", key)
		for i, node := range nodes {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(node.GetIdentifier())
		}
		fmt.Println("]")
	}

	// Also demo GetPrimaryNode
	fmt.Println("\nPrimary nodes only:")
	for _, key := range testKeys {
		node, err := ring.GetPrimaryNode(key)
		if err != nil {
			log.Printf("Error getting primary node for %s: %v", key, err)
			continue
		}
		fmt.Printf("  %s -> %s (primary)\n", key, node.GetIdentifier())
	}

	// Remove a server and show remapping
	fmt.Println("\nRemoving server-2...")
	if err := ring.RemoveNode(nodes[1]); err != nil {
		log.Printf("Error removing node: %v", err)
	} else {
		fmt.Println("✓ Removed: server-2")
	}

	fmt.Println("\nKeys after removing server-2 (replicas):")
	for _, key := range testKeys {
		nodes, err := ring.GetNodesForKey(key)
		if err != nil {
			log.Printf("Error getting nodes for %s: %v", key, err)
			continue
		}
		fmt.Printf("  %s -> [", key)
		for i, node := range nodes {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(node.GetIdentifier())
		}
		fmt.Println("]")
	}
}
