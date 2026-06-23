package array
/*
 * 678. Valid Parenthesis String (medium)

 Given a string s containing only three types of characters: '(', ')' and '*', return true if s is valid.

The following rules define a valid string:

Any left parenthesis '(' must have a corresponding right parenthesis ')'.
Any right parenthesis ')' must have a corresponding left parenthesis '('.
Left parenthesis '(' must go before the corresponding right parenthesis ')'.
'*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string "".
 

Example 1:
>>>>>>>>>
Input: s = "()"
Output: true

Example 2:
>>>>>>>>>
Input: s = "(*)"
Output: true

Example 3:
>>>>>>>>>
Input: s = "(*))"
Output: true

Constraints:

1 <= s.length <= 100
s[i] is '(', ')' or '*'.
 */


func CheckValidString(s string) bool {
    n := len(s)
    bal := 0
    for i := 0; i < n; i++ {
        if s[i] == '(' || s[i] == '*' {
            bal++
        } else {
            bal--
        }
        if bal < 0 {
            return false
        }
    }
    bal = 0
    for i := n - 1; i >= 0; i-- {
        if s[i] == ')' || s[i] == '*' {
            bal++
        } else {
            bal--
        }
        if bal < 0 {
            return false
        }
    }
    return true;
}