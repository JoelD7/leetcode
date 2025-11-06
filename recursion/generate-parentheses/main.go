package generate_parentheses

func generateParenthesis(n int) []string {
	res := make([]string, 0)

	res = backtrack(0, 0, n, "", res)

	return res
}

func backtrack(open, closed, n int, s string, res []string) []string {
	if open == closed && open+closed == n*2 {
		res = append(res, s)
		return res
	}

	if open < n {
		res = backtrack(open+1, closed, n, s+"(", res)
	}

	if closed < open {
		res = backtrack(open, closed+1, n, s+")", res)
	}

	return res
}
