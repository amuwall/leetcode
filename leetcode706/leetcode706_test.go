package leetcode706

import (
	"testing"
)

func TestMyHashMap(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyHashMap{}
			this.Remove(tt.args.key)
		})
	}

	myHashMap := Constructor()
	myHashMap.Put(1, 1)
	myHashMap.Put(2, 2)
	if myHashMap.Get(1) != 1 {
		t.Errorf("Get(1) got %v, want %v", myHashMap.Get(1), 1)
		return
	}
	if myHashMap.Get(3) != -1 {
		t.Errorf("Get(3) got %v, want %v", myHashMap.Get(3), -1)
		return
	}
	myHashMap.Put(2, 1)
	if myHashMap.Get(2) != 1 {
		t.Errorf("Get(2) got %v, want %v", myHashMap.Get(2), 1)
		return
	}
	myHashMap.Remove(2)
	if myHashMap.Get(2) != -1 {
		t.Errorf("Get(2) got %v, want %v", myHashMap.Get(2), -1)
		return
	}
}
