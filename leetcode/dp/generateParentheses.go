package dp

func GenerateParenthesis(n int) []string {
    var res []string
    var stack []byte
    backtrack(n, 0, 0, &res, stack)
    return res
}

func backtrack(n int, openN int, closeN int, res *[]string, stack []byte) {
    if openN == n && closeN == n {
        *res = append(*res, string(stack))
        return
    }
    if openN < n {
        stack = append(stack, '(')
        backtrack(n, openN + 1, closeN, res, stack)
        stack = stack[:len(stack) - 1]
    }
    if closeN < openN {
        stack = append(stack, ')')
        backtrack(n, openN, closeN + 1, res, stack)
        stack = stack[:len(stack) - 1]
    }
}