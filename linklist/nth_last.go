// // /**
// //  * Definition for singly-linked list.
// //  * type ListNode struct {
// //  *     Val int
// //  *     Next *ListNode
// //  * }
// //  */
// //  func removeNthFromEnd(head *ListNode, n int) *ListNode {

// // }

// Given a linked list, remove the n-th node from the end of list and return its head.

// Example:

// Given linked list: 1->2->3->4->5, and n = 2.

// After removing the second node from the end, the linked list becomes 1->2->3->5.

// ListNode removeNthFromEnd(ListNode head, int n) {
// 	ListNode dummy = new ListNode(0);
// 	dummy.next = head;
// 	ListNode a = dummy, b = dummy;
// 	while (n-- > 0) {
// 		a = a.next;
// 	}
// 	while (a.next != null) {
// 		a = a.next;
// 		b = b.next;
// 	}
// 	b.next = b.next.next;
// 	return dummy.next;
// }

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dummy := &ListNode{Val: 0}
	dummy.Next = head
	a := dummy
	b := dummy
	for n > 0 {
		a = a.Next
		n--
	}
	for a.Next != nil {
		a = a.Next
		b = b.Next
	}
	b.Next = b.Next.Next
	return dummy.Next
}