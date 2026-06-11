package stack

func IsValidParentheses(s string) bool {
    stack := Stack[rune]{}
    // Map to easily check matching pairs
    pairs := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    for _, char := range s {
        if openBracket, exists := pairs[char]; exists {
            topElement, _ := stack.Peek()
            if stack.IsEmpty() || topElement != openBracket {
                return false
            }
            stack.Pop()
        } else {
            stack.Push(char)
        }
    }
    return stack.IsEmpty()
}