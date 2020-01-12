/*
 * 2848869
 * 8089098
 * 3861852
 */

package activity

import (
	"log"
)
//inspired by https://medium.com/hackernoon/build-a-go-cache-in-10-minutes-c908a8255568

const SIZE = 10 // size of cache

//Node contains am id, an activity and its neighbor nodes
type Node struct {
	ActivityId	string
	Val   Activity
	Left  *Node
	Right *Node
}

// double linked list
type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

// maps activity to node in Queue
type Hash map[string]*Node

type Cache struct {
	Queue Queue
	Hash  Hash
}

//create a new cache
func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

//create a new queue
func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}
	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}

//checkin and verify a new activity to checkin the cache
func (c *Cache) Check(id string, activity Activity) Activity {
	node := &Node{}
	isInCache, cacheNode := c.GetNode(id)

	if isInCache {
		node = c.Remove(cacheNode)
	} else {
		node = &Node{ActivityId: id, Val: activity}
	}

	c.AddNode(node)
	c.Hash[id] = node
	return node.Val
}

//get an activity saved in cache by its id
func (c *Cache) GetActivity(id string) (activityInCache bool, activityFromCache Activity) {
	var hasNode = false
	var activity Activity
	if val, ok := c.Hash[id]; ok {
		activity = val.Val
		hasNode = true
	}
	return hasNode, activity
}

//get a node from cache by its id
func (c *Cache) GetNode(id string) (activityInCache bool, nodeFromCache *Node) {
	var hasNode = false
	var node *Node
	if val, ok := c.Hash[id]; ok {
		node = val
		hasNode = true
	}
	return hasNode, node
}

//remove a node by its id
func (c *Cache) RemoveById(id string) {
	hasNode, node := c.GetNode(id)
	if hasNode {
		//reassign neighbors
		left := node.Left
		right := node.Right
		left.Right = right
		right.Left = left
		//shorten queue length by one
		c.Queue.Length -= 1
		//remove node
		delete(c.Hash, node.ActivityId)
	} else {
		log.Println("Cannot retrieve node from cache")
	}
}

//remove node from cache
func (c *Cache) Remove(n *Node) *Node {
	//reassign neighbors
	left := n.Left
	right := n.Right
	left.Right = right
	right.Left = left
	//shorten queue length by one
	c.Queue.Length -= 1
	//remove node
	delete(c.Hash, n.ActivityId)
	return n
}

//add a node to cache
func (c *Cache) AddNode(n *Node) {
	tmp := c.Queue.Head.Right
	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = tmp
	tmp.Left = n

	c.Queue.Length++
	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}
}