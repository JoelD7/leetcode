package stack

/*
- Itera sobre el string y ve agregando cada caracter c a un stack
- Si c es un opening tag agregalo al stack
- Si c es closing tag, haz pop del stack y ese elemento debe corresopnder con el closing tag, si no lo hace, return false.
- Al final de la iteracion return true.
*/
func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}

	openingTag := map[string]struct{}{
		"(": {},
		"{": {},
		"[": {},
	}

	closingTag := map[string]struct{}{
		")": {},
		"}": {},
		"]": {},
	}

	closedToOpen := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	var char, stack, poped, openChar string
	var ok bool

	for _, c := range s {
		char = string(c)

		_, ok = openingTag[char]
		if ok {
			stack = char + stack
			continue
		}

		_, ok = closingTag[char]
		if ok && len(stack) == 0 {
			return false
		}

		if ok {
			poped = stack[0:1]
			stack = stack[1:]
		}

		openChar, _ = closedToOpen[char]
		if openChar != poped {
			return false
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
