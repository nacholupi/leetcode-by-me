package array

import (
	"testing"
)

// ["RandomizedSet","insert","remove","insert","getRandom","remove","insert","getRandom"]
// input: [[],[1],[2],[2],[],[1],[2],[]]
// output: [null,true,false,true,1,true,false,2]
func TestRandomizedSet(t *testing.T) {
	obj := Constructor()
	if !obj.Insert(1) {
		t.Errorf("Insert(1) = false; want true")
	}
	if obj.Remove(2) {
		t.Errorf("Remove(2) = true; want false")
	}
	if !obj.Insert(2) {
		t.Errorf("Insert(2) = false; want true")
	}
	rand := obj.GetRandom()
	if rand != 1 && rand != 2 {
		t.Errorf("GetRandom() = %d; want 1 or 2", rand)
	}
	if !obj.Remove(1) {
		t.Errorf("Remove(1) = false; want true")
	}
	if obj.Insert(2) {
		t.Errorf("Insert(2) = true; want false")
	}
	rand = obj.GetRandom()
	if rand != 2 {
		t.Errorf("GetRandom() = %d; want 2", rand)
	}
}
