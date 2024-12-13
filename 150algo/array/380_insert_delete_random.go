package array

import (
	"math/rand"
)

type RandomizedSet struct {
	valueIdxMap map[int]int // value -> idx
	idxValueMap map[int]int // idx -> value
	rd          *rand.Rand
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		valueIdxMap: make(map[int]int),
		idxValueMap: make(map[int]int),
		rd:          rand.New(rand.NewSource(99)),
	}
}

func (r *RandomizedSet) Insert(val int) bool {
	mapLen := len(r.valueIdxMap)
	_, ok := r.valueIdxMap[val]
	if !ok {
		r.valueIdxMap[val] = mapLen
		r.idxValueMap[mapLen] = val
	}
	return !ok
}

func (r *RandomizedSet) Remove(val int) bool {
	idx, ok := r.valueIdxMap[val]

	if ok {
		lastIdx := len(r.valueIdxMap) - 1
		lastValue := r.idxValueMap[lastIdx]
		r.idxValueMap[idx] = lastValue
		r.valueIdxMap[lastValue] = idx
		delete(r.idxValueMap, lastIdx)
		delete(r.valueIdxMap, val)

	}
	return ok
}

func (r *RandomizedSet) GetRandom() int {
	if len(r.valueIdxMap) == 0 {
		return -1
	}
	key := r.rd.Intn(len(r.valueIdxMap))
	return r.idxValueMap[key]
}
