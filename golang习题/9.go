
// 题目
// 将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
// 示例：
// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
//
// 输入：l1 = [], l2 = [0]
// 输出：[0]

package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func TEST_DO_NOT_CHANGE(l1, l2 *ListNode) *ListNode {
    var head *ListNode
    // start下面可以改动



    // end上面可以改动
    return head
}

func main() {
    createList := func(vals []int) *ListNode {
        if len(vals) == 0 {
            return nil
        }
        head := &ListNode{Val: vals[0]}
        current := head
        for _, v := range vals[1:] {
            current.Next = &ListNode{Val: v}
            current = current.Next
        }
        return head
    }

    printList := func(node *ListNode) {
        for node != nil {
            fmt.Printf("%d -> ", node.Val)
            node = node.Next
        }
        fmt.Println("nil")
    }

    l1 := createList([]int{1, 2, 4})
    l2 := createList([]int{1, 3, 4})
    printList(TEST_DO_NOT_CHANGE(l1, l2))

    l1 = createList([]int{})
    l2 = createList([]int{})
    printList(TEST_DO_NOT_CHANGE(l1, l2))

    l1 = createList([]int{})
    l2 = createList([]int{0})
    printList(TEST_DO_NOT_CHANGE(l1, l2))
}
