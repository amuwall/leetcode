package leetcode380

import "math/rand"

type RandomizedSet struct {
	nums     []int
	numIndex map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums:     []int{},
		numIndex: map[int]int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.numIndex[val]; ok {
		return false
	}

	this.nums = append(this.nums, val)
	this.numIndex[val] = len(this.nums) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.numIndex[val]; !ok {
		return false
	}

	index := this.numIndex[val]
	this.nums[index] = this.nums[len(this.nums)-1]
	this.numIndex[this.nums[index]] = index

	this.nums = this.nums[:len(this.nums)-1]
	delete(this.numIndex, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
