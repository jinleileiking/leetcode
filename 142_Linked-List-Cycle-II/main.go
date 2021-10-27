/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func detectCycle(head *ListNode) *ListNode {
    l := map[*LiastNode]bool{}
    for node := head; node != nil; node = node.Next {
        if l[node]  {
            return node
        }
        l[node] = true
    }
    return nil
}