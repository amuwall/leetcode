package leetcode20

type Stack struct {
	data []byte
}

func (stack *Stack) Push(b byte) {
	stack.data = append(stack.data, b)
}

func (stack *Stack) Pop() byte {
	size := stack.Size()
	if size == 0 {
		return 0
	}

	b := stack.data[size-1]
	stack.data = stack.data[:size-1]
	return b
}

func (stack *Stack) Size() int {
	return len(stack.data)
}

func isValid(s string) bool {
	stack := &Stack{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '{', '[':
			stack.Push(s[i])
		case ')':
			if stack.Size() == 0 || stack.Pop() != '(' {
				return false
			}
		case '}':
			if stack.Size() == 0 || stack.Pop() != '{' {
				return false
			}
		case ']':
			if stack.Size() == 0 || stack.Pop() != '[' {
				return false
			}
		default:
			return false
		}
	}

	return stack.Size() == 0
}
