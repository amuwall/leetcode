package leetcode705

import (
	"testing"
)

func TestMyHashMap(t *testing.T) {
	myHashSet := Constructor()
	myHashSet.Add(1)
	myHashSet.Add(2)
	if myHashSet.Contains(1) != true {
		t.Errorf("Contains(1) got %v, want %v", myHashSet.Contains(1), true)
		return
	}
	if myHashSet.Contains(3) != false {
		t.Errorf("Contains(3) got %v, want %v", myHashSet.Contains(3), false)
		return
	}
	myHashSet.Add(2)
	if myHashSet.Contains(2) != true {
		t.Errorf("Contains(2) got %v, want %v", myHashSet.Contains(2), true)
		return
	}
	myHashSet.Remove(2)
	if myHashSet.Contains(2) != false {
		t.Errorf("Contains(2) got %v, want %v", myHashSet.Contains(2), false)
		return
	}
}
