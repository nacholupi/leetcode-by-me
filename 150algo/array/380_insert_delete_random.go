package array

import (
	"math/rand"
)

type RandomizedSet struct {
	valueIdxMap map[int]int // value -> idx
	idxValueMap map[int]int // idx -> value
	r           *rand.Rand
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		valueIdxMap: make(map[int]int),
		idxValueMap: make(map[int]int),
		r:           rand.New(rand.NewSource(99)),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	mapLen := len(this.valueIdxMap)
	_, ok := this.valueIdxMap[val]
	if !ok {
		this.valueIdxMap[val] = mapLen
		this.idxValueMap[mapLen] = val
	}
	return !ok
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.valueIdxMap[val]

	if ok {
		lastIdx := len(this.valueIdxMap) - 1
		lastValue := this.idxValueMap[lastIdx]
		this.idxValueMap[idx] = lastValue
		this.valueIdxMap[lastValue] = idx
		delete(this.idxValueMap, lastIdx)
		delete(this.valueIdxMap, val)

	}
	return ok
}

func (this *RandomizedSet) GetRandom() int {
	if len(this.valueIdxMap) == 0 {
		return -1
	}
	key := this.r.Intn(len(this.valueIdxMap))
	return this.idxValueMap[key]
}
