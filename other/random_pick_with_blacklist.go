package other

import "math/rand"

type Solution struct {
	indexMap map[int]int
	sz       int
}

func Constructor1(n int, blacklist []int) Solution {
	indexMap := make(map[int]int)
	sz := n - len(blacklist)
	for i := 0; i < len(blacklist); i++ {
		indexMap[blacklist[i]] = 666
	}

	last := n - 1
	for i := 0; i < len(blacklist); i++ {
		if blacklist[i] >= sz {
			continue
		}
		for last >= 0 {
			if _, ok := indexMap[last]; !ok {
				break
			}
			last--
		}
		indexMap[blacklist[i]] = last
		last--
	}
	return Solution{indexMap: indexMap, sz: sz}
}

func (this *Solution) Pick() int {
	randomIdx := rand.Intn(this.sz)
	if _, ok := this.indexMap[randomIdx]; ok {
		return this.indexMap[randomIdx]
	}
	return randomIdx
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
