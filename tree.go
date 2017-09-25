package main

import(
	"fmt"
)

type node struct {
	data int
	left *node
	right *node
}

func insert(head *node, data int){
	newnode := &node{data : data, left : nil, right : nil}
	temp := head
	for {
		if temp.data > data {
			if temp.left == nil {
				temp.left = newnode
				return
			}
			temp = temp.left
	
		}else if temp.data <= data {
			if temp.right == nil {
				temp.right = newnode
				return
			}
			temp = temp.right
		}
	}
}

func inorder(n *node){
	if n != nil {
		inorder(n.left)
		fmt.Printf("%d\n",n.data)
		inorder(n.right)
	}
}

func search(head *node, x int)bool{
	if head == nil {
		return false
	}
	if x == head.data {
		return true
	}else if x > head.data {
		return search(head.right, x)
	}else if x < head.data {
		return search(head.left, x)
	}
	return false
}

func deletenode(head *node, x int)*node {
	if head == nil {
		return head
	}
	if x < head.data {
		head.left = deletenode(head.left, x)
	}else if x > head.data {
		head.right = deletenode(head.right, x)
	}else{
		if head.left == nil {
			var temp *node = head.right
			return temp
		}else if head.right == nil {
			var temp *node = head.left
			return temp
		}
		var temp *node = minvalue(head.right)
		head.data = temp.data
		head.right = deletenode(head.right, temp.data)
	}
	return head
}

func minvalue(n *node)*node{
	var current *node
	if current.left != nil {
		current = current.left
	}
	return current
}

func main(){
	var head *node
	head = &node{data : 7, left : nil, right : nil}
	insert(head, 6)
	insert(head, 10)
	insert(head, 4)
	insert(head, 5)
	insert(head, 9)
	insert(head, 11)
	insert(head, 3)
	inorder(head)
	fmt.Println()
	fmt.Printf("%v\n", head.left)
	fmt.Printf("%v\n", head.right)
	fmt.Printf("%v\n", head.left.left)
	fmt.Printf("%v\n", head.left.right)
	fmt.Printf("%v\n", head.right.left)
	fmt.Printf("%v\n", head.right.right)
	fmt.Println()
	if search(head, 49){
		fmt.Printf("Present\n")
	}else{
		fmt.Printf("Not Present\n")
	}
	fmt.Println()
	deletenode(head, 11)
	inorder(head)
}


