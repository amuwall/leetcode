package leetcode705

type MyHashSet struct {
	nodes [1000001]bool
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func (this *MyHashSet) Add(key int) {
	this.nodes[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.nodes[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
	return this.nodes[key]
}
