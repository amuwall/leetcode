package leetcode2452

type Node struct {
	NextNodes map[rune]*Node
}

func initTree(dictionary []string) (root *Node) {
	root = &Node{
		NextNodes: map[rune]*Node{},
	}

	for _, s := range dictionary {
		p := root
		for _, ch := range s {
			next, ok := p.NextNodes[ch]
			if !ok {
				p.NextNodes[ch] = &Node{
					NextNodes: map[rune]*Node{},
				}
				next = p.NextNodes[ch]
			}

			p = next
		}
	}

	return
}

func find(head *Node, word string, wordIndex, editCount int) bool {
	if editCount < 0 {
		return false
	}

	if wordIndex == len(word) {
		return true
	}

	for key, next := range head.NextNodes {
		nextEditCount := editCount
		if key != rune(word[wordIndex]) {
			nextEditCount = editCount - 1
		}
		if find(next, word, wordIndex+1, nextEditCount) {
			return true
		}
	}

	return false
}

func twoEditWords(queries []string, dictionary []string) []string {
	root := initTree(dictionary)

	result := []string{}
	for _, s := range queries {
		if find(root, s, 0, 2) {
			result = append(result, s)
		}
	}

	return result
}
