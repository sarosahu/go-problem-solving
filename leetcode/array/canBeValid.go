package array

/*
 * 2116. Check if a Parentheses String Can Be Valid

 * A parentheses string is a non-empty string consisting only of '(' and ')'. It is valid if any of the following conditions is true:

It is ().
It can be written as AB (A concatenated with B), where A and B are valid parentheses strings.
It can be written as (A), where A is a valid parentheses string.
You are given a parentheses string s and a string locked, both of length n. locked is a binary string consisting only of '0's and '1's. For each index i of locked,

If locked[i] is '1', you cannot change s[i].
But if locked[i] is '0', you can change s[i] to either '(' or ')'.
Return true if you can make s a valid parentheses string. Otherwise, return false.

Example 1:
Input: s = "))()))", locked = "010100"
Output: true
Explanation: locked[1] == '1' and locked[3] == '1', so we cannot change s[1] or s[3].
We change s[0] and s[4] to '(' while leaving s[2] and s[5] unchanged to make s valid.

Example 2:

Input: s = "()()", locked = "0000"
Output: true
Explanation: We do not need to make any changes because s is already valid.

Example 3:

Input: s = ")", locked = "0"
Output: false
Explanation: locked permits us to change s[0]. 
Changing s[0] to either '(' or ')' will not make s valid.

Example 4:

Input: s = "(((())(((())", locked = "111111010111"
Output: true
Explanation: locked permits us to change s[6] and s[8]. 
We change s[6] and s[8] to ')' to make s valid.

Constraints:

n == s.length == locked.length
1 <= n <= 10^5
s[i] is either '(' or ')'.
locked[i] is either '0' or '1'.
 */

// This one is easy to understand
/*
 * A useful trick (when doing any parentheses validation) is to greedily check balance left-to-right, and then right-to-left.

Left-to-right check ensures that we do not have orphan ')' parentheses.
Right-to-left checks for orphan '(' parentheses.
We go left-to-right:

Count wild (not locked) characters.
Track the balance bal for locked parentheses.
If the balance goes negative, we check if we have enough wild characters to compensate.
In the end, check that we have enough wild characters to cover positive balance (open parentheses).
This approach alone, however, will fail for ["))((", "0011"] test case. That is why we also need to do the same going right-to-left.
 */
func CanBeValid(s string, locked string) bool {
    n := len(s)
    if n % 2 == 1 {
        return false
    }
    if isValid(s) {
        return true
    }
    bal, flippable := 0, 0
    // Iterate left to right
    for i := 0; i < n; i++ {
        if locked[i] == '0' {
            flippable++
        } else {
            if s[i] == '(' {
                bal++
            } else {
                bal--
            }
        }
        // More closing brackets than we can balance
        if bal + flippable < 0 {
            return false
        }
    }
    if bal > flippable {
        return false
    }

    bal, flippable = 0, 0
    // Iterate right to left
    for i := n - 1; i >= 0; i-- {
        if locked[i] == '0' {
            flippable++
        } else {
            if s[i] == ')' {
                bal++
            } else {
                bal--
            }
        }
        // More opening brackets than we can balance
        if bal + flippable < 0 {
            return false
        }
    }
    if bal > flippable {
        return false
    }
    return true
}

func isValid(s string) bool {
	bal := 0
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			bal++
		} else {
			bal--
		}
		if bal < 0 {
			return false
		}
	}
	return bal == 0
}

