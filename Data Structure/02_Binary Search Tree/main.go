// Node & BST Wrapper, String() - InOrderTraversal, Add(), Search(), Remove()

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type bst struct {
	root *node
	size int
}

type node struct {
	value int
	left  *node
	right *node
}

func (n node) String() string {
	return strconv.Itoa(n.value)
}

func (b *bst) insert(value int) *node {
	b.root = b.insertByNode(b.root, value)
	return b.root
}

func (b *bst) insertByNode(root *node, value int) *node {
	if root == nil {
		b.size++
		return &node{value: value}
	}
	if value < root.value {
		root.left = b.insertByNode(root.left, value)
	} else {
		root.right = b.insertByNode(root.right, value)
	}
	return root
}

func (b *bst) remove(value int) *node {
	return b.removeByNode(b.root, value)
}

func (b *bst) removeByNode(root *node, value int) *node {
	if root == nil {
		return root
	}

	if value < root.value {
		root.left = b.removeByNode(root.left, value)
	} else if value > root.value {
		root.right = b.removeByNode(root.right, value)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}
		root.value = b.minValue(root.right)

		root.right = b.removeByNode(root.right, root.value)
	}
	return root
}

func (b bst) minValue(root *node) int {
	min := root.value
	for root.left != nil {
		min = root.left.value
		root = root.left
	}
	return min
}

func (b *bst) search(value int) (*node, bool) {
	return b.searchByNode(b.root, value)
}

func (b bst) searchByNode(root *node, value int) (*node, bool) {
	if root == nil {
		return nil, false
	}

	if value == root.value {
		return root, true
	} else if value < root.value {
		return b.searchByNode(root.left, value)
	} else {
		return b.searchByNode(root.right, value)
	}
}

func (b bst) inOrderTraversal() {
	b.inOrderTraversalByNode(b.root)
	fmt.Println()
}  

func (b bst) inOrderTraversalByNode(root *node) {
	if root == nil {
		return
	}
	b.inOrderTraversalByNode(root.left)
	fmt.Printf("%d, ", root.value)
	b.inOrderTraversalByNode(root.right)
}

func (b bst) levelTraversal() {
	levels := new ([][]*node)
	b.levelTraversalByNode(levels, b.root, 0)

	sb := strings.Builder{}
	for _, level := range *levels {
		sb.WriteString(fmt.Sprintf("%s\n", level))
	}
	fmt.Print(sb.String())
}

func (b bst) levelTraversalByNode(q *[][]*node, root *node, level int) {
	if root == nil {
		return
	}

	if len(*q) <= level {
		*q = append(*q, []*node{root})
	} else {
		(*q)[level] = append((*q)[level], root)
	}
	b.levelTraversalByNode(q, root.left, level + 1)
	b.levelTraversalByNode(q, root.right, level + 1)
}

func (b bst) String() string {
	sb := strings.Builder{}
	b.inOrderTraversal2(&sb)
	return sb.String()
}

func (b bst) inOrderTraversal2(sb *strings.Builder) {
	b.inOrderTraversalByNode2(sb, b.root)
}

func (b bst) inOrderTraversalByNode2(sb *strings.Builder, root *node) {
	if root == nil {
		return
	}
	b.inOrderTraversalByNode2(sb, root.left)
	sb.WriteString(fmt.Sprintf("%s ", root))
	b.inOrderTraversalByNode2(sb, root.right)
}

func main() {
	n := &node{1, nil, nil}
	n.left = &node{0, nil, nil}
	n.right = &node{2, nil, nil}
	b := bst {
		root: n,
		size: 1,
	}
	b.search(2)
	b.inOrderTraversal()
	b.levelTraversal()

	b2 := bst{}
	b2.insert(0)
	b2.insert(1)
	b2.insert(2)
	b2.inOrderTraversal()
	b2.levelTraversal()

	fmt.Println(b2)
	b2.remove(2)
	b2.remove(1)
	b2.remove(2)
	fmt.Println(b2)

	var s []int
	fmt.Println(s[0])

}