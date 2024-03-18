package leetcode303

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	array := NumArray{
		sums: make([]int, len(nums)),
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			array.sums[i] = nums[i]
		} else {
			array.sums[i] = nums[i] + array.sums[i-1]
		}
	}

	return array
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.sums[right]
	}
	return this.sums[right] - this.sums[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
