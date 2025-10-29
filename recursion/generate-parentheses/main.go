package generate_parentheses

func generateParenthesis(n int) []string {
	res := make([]string, 0)

	recurse(0, 0, n, &res, "")

	return res
}

func recurse(open, closed, n int, res *[]string, s string) {
	if open == closed && len(s) == n*2 {
		*res = append(*res, s)
		return
	}

	if open < n {
		recurse(open+1, closed, n, res, s+"(")
	}

	if closed < open {
		recurse(open, closed+1, n, res, s+")")
	}
}
