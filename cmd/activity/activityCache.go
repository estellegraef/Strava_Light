/*
 * 2848869
 * 8089098
 * 3861852
 */

package activity

import (
	"fmt"
)

const SIZE = 5 // size of cache

//Node contains a activity and neighbor nodes
type Node struct {
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
type Hash map[Activity]*Node

type Cache struct {
	Queue Queue
	Hash  Hash
}

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}
	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}

func (c *Cache) Check(activity Activity) Activity {
	node := &Node{}
	if val, ok := c.Hash[activity]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Val: activity}
	}

	c.Add(node)
	c.Hash[activity] = node
	return node.Val
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
	delete(c.Hash, n.Val)
	return n
}

func (c *Cache) Add(n *Node) {
	fmt.Printf("add: %v\n", n.Val)
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

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%v}", node.Val)
		if i < q.Length-1 {
			fmt.Printf(" <--> ")
		}
		node = node.Right
	}
	fmt.Println("]")
}