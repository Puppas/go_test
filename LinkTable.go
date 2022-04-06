package main

import "fmt"

type LinkNode struct {
	val interface{}	//val可以为任何值
	next *LinkNode
}

//head指向空节点，易于实现链表操作
type LinkTable struct {
	head *LinkNode
	len int
}


func (list *LinkTable) init() {
	list.head = new(LinkNode)
	list.head.next = nil
	list.len = 0
}

func (list *LinkTable) empty() bool {
	return list.len == 0
}


func (list *LinkTable) size() int {
	return list.len
}


func (list *LinkTable) push_front(data interface{}) {
	new_node := &LinkNode{val: data}
	new_node.next = list.head.next
	list.head.next = new_node
	list.len++
}



//在pos前插入
func (list *LinkTable) insert(data interface{}, pos *LinkNode) {
	p := list.head
	new_node := &LinkNode{val: data}
	for p.next != pos {
		p = p.next
	}
	new_node.next = pos
	p.next = new_node
	list.len++
}

//链表删除指定元素
func (list *LinkTable) erase(pos *LinkNode) {
	p := list.head
	for p.next != pos {
		p = p.next
	}
	
	p.next = p.next.next
}

//打印链表
func (list *LinkTable) print() {
	if list.empty() {
		fmt.Println("empty")
	} else {
		p := list.head.next
		for p != nil {
			fmt.Print(p.val)
			fmt.Print(" ")
			p = p.next
		}
		fmt.Print("\n")
	}
}