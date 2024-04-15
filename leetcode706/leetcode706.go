package leetcode706

type MyHashMap struct {
	nodes [1000001]int
}

func Constructor() MyHashMap {
	myHashMap := MyHashMap{}
	for i, _ := range myHashMap.nodes {
		myHashMap.nodes[i] = -1
	}

	return myHashMap
}

func (this *MyHashMap) Put(key int, value int) {
	this.nodes[key] = value
}

func (this *MyHashMap) Get(key int) int {
	return this.nodes[key]
}

func (this *MyHashMap) Remove(key int) {
	this.nodes[key] = -1
}
