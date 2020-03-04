// https://www.hackerrank.com/interview/interview-preparation-kit/linked-lists/challenges
package linked_lists

import "fmt"

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func (ll *SinglyLinkedListNode) String() string {
	if ll == nil {
		return "empty_linked_list"
	}

	result := fmt.Sprintf("%d", ll.data)

	for {
		ll = ll.next
		if ll == nil {
			break
		}
		result = fmt.Sprintf("%s %v", result, ll.data)
	}

	return result
}

type DoublyLinkedListNode struct {
	data int32
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

func (ll *DoublyLinkedListNode) String() string {
	if ll == nil {
		return "empty_linked_list"
	}

	result := fmt.Sprintf("%d", ll.data)

	for {
		ll = ll.next
		if ll == nil {
			break
		}
		result = fmt.Sprintf("%s %v", result, ll.data)
	}

	return result
}

// https://www.hackerrank.com/challenges/insert-a-node-at-a-specific-position-in-a-linked-list/problem
func insertNodeAtPosition(head *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	if head == nil {
		return &SinglyLinkedListNode{data, nil}
	}
	if position == 0 {
		return &SinglyLinkedListNode{data, head}
	}

	e := head

	var i int32 = 1
	for {
		if i == position {
			e.next = &SinglyLinkedListNode{data, e.next}
			break
		}

		e = e.next
		i++
	}

	return head
}

// https://www.hackerrank.com/challenges/insert-a-node-into-a-sorted-doubly-linked-list/problem
func sortedInsert(head *DoublyLinkedListNode, data int32) *DoublyLinkedListNode {
	if head == nil {
		return &DoublyLinkedListNode{data, nil, nil}
	}
	if head.data >= data {
		tail := head
		head := &DoublyLinkedListNode{data, tail, nil}
		head.next.prev = head
		return head
	}

	e := head

	for {
		if e.data >= data {
			node := &DoublyLinkedListNode{data, e, nil}
			node.prev = e.prev
			node.next.prev = node
			node.prev.next = node
			break
		}

		if e.next == nil {
			e.next = &DoublyLinkedListNode{data, nil, e}
			break
		}

		e = e.next
	}

	return head
}

// https://www.hackerrank.com/challenges/reverse-a-doubly-linked-list/problem
func reverse(head *DoublyLinkedListNode) *DoublyLinkedListNode {
	if head == nil || head.next == nil {
		return head
	}

	for {
		head.next, head.prev = head.prev, head.next

		if head.prev == nil {
			return head
		}

		head = head.prev
	}
}

// https://www.hackerrank.com/challenges/find-the-merge-point-of-two-joined-linked-lists/problem
func findMergeNode(head1 *SinglyLinkedListNode, head2 *SinglyLinkedListNode) int32 {
	count1 := countLinkedList(head1)
	count2 := countLinkedList(head2)

	var longest *SinglyLinkedListNode
	var shortest *SinglyLinkedListNode
	var diff int

	if count1 > count2 {
		diff = count1 - count2
		longest = head1
		shortest = head2
	} else {
		diff = count2 - count1
		longest = head2
		shortest = head1
	}

	for i := 0; i < diff; i++ {
		longest = longest.next
	}

	for {
		if longest == shortest {
			return longest.data
		}
		longest = longest.next
		shortest = shortest.next
	}
}

func countLinkedList(head *SinglyLinkedListNode) int {
	count := 0
	for {
		if head == nil {
			return count
		}
		head = head.next
		count++
	}
}

// https://www.hackerrank.com/challenges/ctci-linked-list-cycle/problem
func detectCycle(head *SinglyLinkedListNode) bool {
	if head == nil {
		return false
	}

	runner1 := head
	runner2 := head

	for {
		if runner1.next == nil || runner2.next == nil || runner2.next.next == nil {
			return false
		}
		runner1 = runner1.next
		runner2 = runner2.next.next

		if runner1 == runner2 {
			return true
		}
	}
}
