package Code

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

// index从1开始，1-n
func (this *TreeArry) Add(index, val int) {
	for index < this.n {
		this.tree[index] += val
		index += lowBit(index)
	}
}

// index从1开始，1-n
func (this *TreeArry) PrefixSum(index int) int {
	res := 0
	for index > 0 {
		res += this.tree[index]
		index -= lowBit(index)
	}
	return res
}
