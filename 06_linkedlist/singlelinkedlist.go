package _6_linkedlist

import (
	"fmt"
	"go/format"
)

/*
单链表基本操作
*/

// 链表结点
type ListNode struct {
	next *ListNode
	value interface{}
}

// 链表
type LinkedList struct {
	head *ListNode
	length uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

// 获取下一个结点
func (this *ListNode) GetNext() *ListNode {
	return this.next
}

func (this *ListNode) GetValue() interface{} {
	return this.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

// 在某个结点后面插入结点
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}
	newNode := NewListNode(v)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	this.length++
	return true
}

// 在某个结点前面抽入结点
func (this *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if nil == p || p == this.head{
		return false
	}
	//  获取当前结点的前一个结点
	cur := this.head.next
	pre := this.head

	if cur == nil {
		return false
	}
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	newNode := NewListNode(v)
	newNode.next = pre.next
	pre.next = newNode
	this.length++
	
	return true
}

// 在链表头部插入结点
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

// 在链表尾部插入结点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for nil != cur.next {
		cur = cur.next
	}

	return this.InsertAfter(cur, v)
}

// 通过索引查找结点
func (this *LinkedList) FindByIndex (index uint) *ListNode {
	if index >= this.length {
		return nil
	}
	cur := this.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

// 删除传入的结点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if p == nil {
		return false
	}
	cur := this.head.next
	pre := this.head

	if cur == nil {
		return false
	}
	
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	pre.next = p.next
	p = nil
	this.length--
	return true
}

// 打印链表
func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}
