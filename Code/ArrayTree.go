package Code

type NumArray struct {
	nums []int
	t    TreeArry
}

type TreeArry struct {
	tree []int
	n    int
}

func lowBit(n int) int {
	return n & (-n)
}

func (this *TreeArry) Init(nums []int) {
	this.n = len(nums) + 1
	this.tree = make([]int, this.n)
	for i, v := range nums {
		this.Add(i+1, v)
	}
}

func (this *TreeArry) Add(index, val int) {
	for index < this.n {
		this.tree[index] += val
		index += lowBit(index)
	}
}

func (this *TreeArry) PrefixSum(index int) int {
	res := 0
	for index > 0 {
		res += this.tree[index]
		index -= lowBit(index)
	}
	return res
}

func Constructor(nums []int) NumArray {
	var a NumArray
	a.nums = nums
	a.t.Init(a.nums)
	return a
}

func (this *NumArray) Update(index int, val int) {
	this.t.Add(index+1, val-this.nums[index])
	this.nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.t.PrefixSum(right+1) - this.t.PrefixSum(left)
}
