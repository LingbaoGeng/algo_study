package other

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	nums     []int
	indexMap map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{nums: make([]int, 0), indexMap: make(map[int]int)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indexMap[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.indexMap[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.indexMap[val]; !ok {
		return false
	}
	index := this.indexMap[val]
	lastVal := this.nums[len(this.nums)-1]
	this.indexMap[lastVal] = index
	this.nums[index] = lastVal
	delete(this.indexMap, val)
	this.nums = this.nums[:len(this.nums)-1]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	fmt.Println(len(this.nums))
	randomIdx := rand.Intn(len(this.nums))
	return this.nums[randomIdx]
}
