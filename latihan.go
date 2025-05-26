package main

import (
	"fmt"
)

// Definisi node untuk singly linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// Fungsi untuk membalik linked list
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head

	for current != nil {
		next := current.Next   // simpan node berikutnya
		current.Next = prev    // balik arah pointer
		prev = current         // geser prev ke current
		current = next         // geser current ke next
	}
	return prev
}

// Fungsi bantu: buat linked list dari slice
func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, v := range nums[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
}

// Fungsi bantu: cetak linked list
func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Contoh penggunaan
	nums := []int{1, 2, 3, 4, 5}
	head := createLinkedList(nums)

	fmt.Print("Sebelum dibalik: ")
	printLinkedList(head)

	reversed := reverseList(head)

	fmt.Print("Setelah dibalik: ")
	printLinkedList(reversed)
}
