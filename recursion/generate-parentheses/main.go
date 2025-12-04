package generate_parentheses

func generateParenthesis(n int) []string {
	res := make([]string, 0)

	return recurse(0, 0, n, "", res)
}

func recurse(open, closed, n int, s string, res []string) []string {
	if open == closed && open == n {
		res = append(res, s)
		return res
	}

	if open < n {
		res = recurse(open+1, closed, n, s+"(", res)
	}

	if closed < open {
		res = recurse(open, closed+1, n, s+")", res)
	}

	return res
}
