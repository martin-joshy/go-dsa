package dsa

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{value: value}
	}

	if value < root.value {
		root.left = insert(root.left, value)
	} else if value > root.value {
		root.right = insert(root.right, value)
	}

	return root
}

func search(root *Node, value int) bool {
	if root == nil {
		return false
	}

	if root.value == value {
		return true
	} else if value < root.value {
		return search(root.left, value)
	} else if value > root.value {
		return search(root.right, value)
	}
}

func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Println(root.value)
	inorder(root.right)
}

func dfs(root *Node) {
	stack := []*Node{root}

	for len(stack) > 0 {
		currNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Print(currNode.value)

		if currNode.right != nil {
			stack = append(stack, currNode.right)
		}

		if currNode.left != nil {
			stack = append(stack, currNode.left)
		}
	}
}

func bfs(root *Node) {
	queue := []*Node{root}

	for len(queue) > 0 {
		currNode := queue[0]
		queue := queue[1:]

		fmt.Print(currNode.value)

		if currNode.left != nil {
			queue = append(queue, currNode.left)
		}

		if currNode.right != nil {
			queue = append(queue, currNode.right)
		}
	}
}

func printLevels(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			currentNode := queue[0]
			queue = queue[1:]

			fmt.Print(currentNode.value, " ")

			if currentNode.left != nil {
				queue = append(queue, currentNode.left)
			}

			if currentNode.right != nil {
				queue = append(queue, currentNode.right)
			}
		}

		fmt.Println("")
	}
}
