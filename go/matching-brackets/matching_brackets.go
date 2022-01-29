package brackets

var openBracketToCloseBracket = map[rune]rune{
	'[': ']',
	'{': '}',
	'(': ')',
}

func Bracket(input string) bool {
	stack := Stack{}
	for _, char := range input {
		if isOpenBracket(char) {
			stack.Push(char)
		} else if isCloseBracket(char) {
			if stack.Len() > 0 && isMatching(stack.Peek(), char) {
				stack.Pop()
			} else {
				return false
			}
		} else {
			continue
		}
	}
	return stack.Len() == 0
}

func isOpenBracket(r rune) bool {
	_, ok := openBracketToCloseBracket[r]
	return ok
}

func isCloseBracket(r rune) bool {
	for _, v := range openBracketToCloseBracket {
		if r == v {
			return true
		}
	}
	return false
}

func isMatching(open rune, close rune) bool {
	return openBracketToCloseBracket[open] == close
}

type Stack struct {
	slice []rune
}

func (s *Stack) Len() int {
	return len(s.slice)
}
func (s *Stack) Push(r rune) {
	s.slice = append(s.slice, r)
}
func (s *Stack) Peek() rune {
	return s.slice[len(s.slice)-1]
}
func (s *Stack) Pop() rune {
	result := s.Peek()
	s.slice = s.slice[:len(s.slice)-1]
	return result
}
