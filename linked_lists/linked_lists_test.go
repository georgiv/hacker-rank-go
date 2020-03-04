package linked_lists

import (
	"fmt"
	"strings"
	"testing"
)

func TestInsertNodeAtPosition(t *testing.T) {
	var tests = []struct {
		head     []int32
		data     int32
		position int32
		expected []int32
	}{
		{[]int32{1, 2, 3},
			4,
			2,
			[]int32{1, 2, 4, 3}},

		{[]int32{16, 13, 7},
			1,
			2,
			[]int32{16, 13, 1, 7}},
	}

	for _, test := range tests {
		head := newSinglyLinkedList(test.head...)
		expected := newSinglyLinkedList(test.expected...)
		if result := insertNodeAtPosition(head, test.data, test.position); !singlyLinkedListDeepEqual(result, expected) {
			t.Errorf("insertNodeAtPosition(%s, %d, %d) = %v, expected = %s",
				strings.Trim(fmt.Sprint(test.head), "[]"),
				test.data,
				test.position,
				result,
				strings.Trim(fmt.Sprint(test.expected), "[]"))
		}
	}
}

func TestSortedInsert(t *testing.T) {
	var tests = []struct {
		head     []int32
		data     int32
		expected []int32
	}{
		{[]int32{1, 3, 4, 10}, 5, []int32{1, 3, 4, 5, 10}},
	}

	for _, test := range tests {
		head := newDoublyLinkedList(test.head...)
		expected := newDoublyLinkedList(test.expected...)
		if result := sortedInsert(head, test.data); !doublyLinkedListDeepEqual(result, expected) {
			t.Errorf("sortedInsert(%s, %d) = %v, expected = %s",
				strings.Trim(fmt.Sprint(test.head), "[]"),
				test.data,
				result,
				strings.Trim(fmt.Sprint(test.expected), "[]"))
		}
	}
}

func TestReverse(t *testing.T) {
	var tests = []struct {
		head     []int32
		expected []int32
	}{
		{nil, nil},
		{[]int32{1}, []int32{1}},
		{[]int32{1, 2, 3, 4}, []int32{4, 3, 2, 1}},
	}

	for _, test := range tests {
		head := newDoublyLinkedList(test.head...)
		expected := newDoublyLinkedList(test.expected...)
		if result := reverse(head); !doublyLinkedListDeepEqual(result, expected) {
			t.Errorf("reverse(%s) = %s, expected = %s",
				strings.Trim(fmt.Sprint(test.head), "[]"),
				result,
				strings.Trim(fmt.Sprint(test.expected), "[]"))
		}
	}
}

func TestFindMergeNode(t *testing.T) {
	var tests = []struct {
		head1    []int32
		head2    []int32
		tail     []int32
		expected int32
	}{
		{[]int32{1}, []int32{1}, []int32{2, 3}, 2},
		{[]int32{1, 2}, []int32{1}, []int32{3}, 3},
	}

	for _, test := range tests {
		head1 := newSinglyLinkedList(test.head1...)
		head2 := newSinglyLinkedList(test.head2...)
		tail := newSinglyLinkedList(test.tail...)
		head1 = concatenateSinglyLinkedLists(head1, tail)
		head2 = concatenateSinglyLinkedLists(head2, tail)
		if result := findMergeNode(head1, head2); result != test.expected {
			t.Errorf("findMergeNode(%v, %v) = %d, expected = %d", head1, head2, result, test.expected)
		}
	}
}

func TestDetectCycle(t *testing.T) {
	var tests = []struct {
		head      []int32
		cycleData int32
		expected  bool
	}{
		{[]int32{1, 2, 3, 4, 5}, 3, true},
		{[]int32{1}, -1, false},
		{[]int32{1, 2, 3}, 2, true},
	}

	for _, test := range tests {
		head := newSinglyLinkedList(test.head...)
		head = insertCycleInSinglyLinkedList(head, test.cycleData)
		if result := detectCycle(head); result != test.expected {
			t.Errorf("linked list: %s, cycle at node with data: %d, result: %t, expected: %t",
				strings.Trim(fmt.Sprint(test.head), "[]"),
				test.cycleData,
				result,
				test.expected)
		}
	}
}

func singlyLinkedListDeepEqual(l1, l2 *SinglyLinkedListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}
	if l1 == nil || l2 == nil {
		return false
	}
	if l1.data != l2.data {
		return false
	}

	return singlyLinkedListDeepEqual(l1.next, l2.next)
}

func doublyLinkedListDeepEqual(l1, l2 *DoublyLinkedListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}
	if l1 == nil || l2 == nil {
		return false
	}
	if l1.data != l2.data {
		return false
	}

	return doublyLinkedListDeepEqual(l1.next, l2.next)
}

func newSinglyLinkedList(data ...int32) *SinglyLinkedListNode {
	if len(data) == 0 {
		return nil
	}

	var head *SinglyLinkedListNode
	var node *SinglyLinkedListNode

	for _, d := range data {
		if head == nil {
			head = &SinglyLinkedListNode{d, nil}
			node = head
		} else {
			node.next = &SinglyLinkedListNode{d, nil}
			node = node.next
		}
	}

	return head
}

func newDoublyLinkedList(data ...int32) *DoublyLinkedListNode {
	if len(data) == 0 {
		return nil
	}

	var head *DoublyLinkedListNode
	var node *DoublyLinkedListNode

	for _, d := range data {
		if head == nil {
			head = &DoublyLinkedListNode{d, nil, nil}
			node = head
		} else {
			node.next = &DoublyLinkedListNode{d, nil, nil}
			node.next.prev = node
			node = node.next
		}
	}

	return head
}

func concatenateSinglyLinkedLists(head *SinglyLinkedListNode, tail *SinglyLinkedListNode) *SinglyLinkedListNode {
	var node = head

	for {
		if node.next == nil {
			node.next = tail
			return head
		} else {
			node = node.next
		}
	}
}

func insertCycleInSinglyLinkedList(head *SinglyLinkedListNode, cycleData int32) *SinglyLinkedListNode {
	if cycleData < 0 {
		return head
	}

	var cycleNode *SinglyLinkedListNode
	node := head
	for {
		if node.data == cycleData {
			cycleNode = node
		}
		if node.next == nil {
			node.next = cycleNode
			return head
		}
		node = node.next
	}
}
