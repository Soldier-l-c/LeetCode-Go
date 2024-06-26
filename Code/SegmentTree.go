package Code

type SegmentTree struct {
	tree   []int
	length int
}

func (this *SegmentTree) Init(nums []int) {

	this.length = len(nums)
	if this.length == 0 {
		return
	}
	this.tree = make([]int, 4*this.length)
	this.length--
	for i, v := range nums {
		this.Update(i, v)
	}
}

// index从0开始，0-length
func (this *SegmentTree) Update(index, val int) {
	this.update(0, 0, this.length, index, val)
}

func (this *SegmentTree) SumRange(left, right int) int {
	return this.get(0, 0, this.length, left, right)
}

func (this *SegmentTree) update(cur_sum_index, l, r, index, val int) {
	if l == r {
		this.tree[cur_sum_index] = val
		return
	}

	mid := l + ((r - l) >> 1)
	if index <= mid {
		this.update(cur_sum_index*2+1, l, mid, index, val)
	} else {
		this.update(cur_sum_index*2+2, mid+1, r, index, val)
	}

	this.tree[cur_sum_index] = this.tree[cur_sum_index*2+1] + this.tree[cur_sum_index*2+2]
}

func (this *SegmentTree) get(cur_sum_index, l, r, L, R int) int {
	if L <= l && r <= R {
		return this.tree[cur_sum_index]
	}

	res, mid := 0, l+((r-l)>>1)

	if L <= mid {
		res += this.get(cur_sum_index*2+1, l, mid, L, R)
	}

	if R > mid {
		res += this.get(cur_sum_index*2+2, mid+1, r, L, R)
	}

	return res
}

/*
class NumArray
{
public:
    NumArray(std::vector<int>& nums)
    {
        if (nums.empty())
            return;

        length_ = nums.size() - 1;

        CreateSegmentTree(nums);
    }

    void update(int index, int val)
    {
        update(0, 0, length_, index, val);
    }

    int sumRange(int left, int right)
    {
        return get(0, 0, length_, left, right);
    }

private:

    void update(int cur_sum_index, int l, int r, int index, int val)
    {
        if (l == r)
        {
            segment_tree_[cur_sum_index] = val;
            return;
        }

        auto mid = (l + r) >> 1;

        if (index <= mid)
        {
            update(cur_sum_index * 2 + 1, l, mid, index, val);
        }
        else
        {
            update(cur_sum_index * 2 + 2, mid+1, r, index, val);
        }
        segment_tree_[cur_sum_index] = segment_tree_[cur_sum_index * 2 + 1] + segment_tree_[cur_sum_index * 2 + 2];
    }

    int get(int cur_sum_index, int l, int r, int L, int R)
    {
        if (L <= l && r <= R)return segment_tree_[cur_sum_index];

        auto mid = (l + r) >> 1;

        auto res{ 0 };
        if (L <= mid)
        {
            res += get(cur_sum_index * 2 + 1, l, mid, L, R);
        }

        if (R > mid)
        {
            res += get(cur_sum_index * 2 + 2, mid + 1, r, L, R);
        }

        return res;
    }

    void CreateSegmentTree(const std::vector<int>& nums)
    {
        auto n = nums.size();
        if (n == 0)return;
        segment_tree_.resize(4 * n);

        for (auto i =0; i<nums.size(); ++i)
        {
            update(0, 0, nums.size() - 1, i, nums[i]);
        }
    }

private:
    std::vector<int>segment_tree_;
    int length_{ 0 };
};
*/
