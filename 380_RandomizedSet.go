package leetcode

import "math/rand"

type RandomizedSet struct {
	nums       []int
	valToIndex map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums:       []int{},
		valToIndex: map[int]int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.valToIndex[val]; exists {
		return false
	}
	this.nums = append(this.nums, val)
	this.valToIndex[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, exists := this.valToIndex[val]
	if !exists {
		return false
	}
	last := this.nums[len(this.nums)-1]
	this.nums[index] = last
	this.valToIndex[last] = index
	this.nums = this.nums[:len(this.nums)-1]
	delete(this.valToIndex, val)
	return true

}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
