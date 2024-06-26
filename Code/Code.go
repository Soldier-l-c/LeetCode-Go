package Code

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(node_list []*ListNode) (*ListNode, *ListNode) {
	var head *ListNode
	var tail *ListNode
	length := len(node_list)
	if 0 == length {
		return head, tail
	}
	start := length - 1

	head = node_list[start]
	tail = node_list[0]

	for start > 0 {
		node_list[start].Next = node_list[start-1]
		start--
	}
	tail.Next = nil
	return head, tail
}

func reverseListS(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func reverseKGroupS(head *ListNode, k int) *ListNode {
	if k == 1 || head == nil {
		return head
	}

	start := head
	end := head
	var last_tail, res_head *ListNode
	length := 0

	for end != nil {
		next := end.Next
		if (length+1)%k == 0 {
			end.Next = nil
			reverseListS(start)
			if res_head == nil {
				res_head = end
			}
			if last_tail != nil {
				last_tail.Next = end
			}
			last_tail = start
			last_tail.Next = next
			start = next
		}

		length++
		end = next
	}
	return res_head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	var res_head *ListNode
	var last_tail *ListNode
	rever_list := make([]*ListNode, k)

	length := 0

	for head != nil {
		next := head.Next
		rever_list[length%k] = head
		if (length+1)%k == 0 {
			cur_head, cur_tail := reverseList(rever_list)
			if nil == res_head || last_tail == nil {
				res_head = cur_head
				last_tail = cur_tail
				cur_tail.Next = next
			} else {
				last_tail.Next = cur_head
				last_tail = cur_tail
				last_tail.Next = next
			}
		}
		head = next
		length++
	}

	return res_head
}

type NumArray struct {
	st     SegmentTree
	at     TreeArry
	nums   []int
	use_at bool
}

func Constructor(nums []int) NumArray {
	var a NumArray
	a.use_at = true
	if !a.use_at {
		a.st.Init(nums)
	} else {
		a.at.Init(nums)
		a.nums = nums
	}

	return a
}

func (this *NumArray) Update(index int, val int) {
	if this.use_at {
		this.at.Add(index+1, val-this.nums[index])
		this.nums[index] = val
	} else {
		this.st.Update(index, val)
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if this.use_at {
		return this.at.PrefixSum(right+1) - this.at.PrefixSum(left)
	} else {
		return this.st.SumRange(left, right)
	}
}

func LeetCodeTest() {
	lfu_cache := Constructor1(10)
	lfu_cache.Put(10, 1)
}
