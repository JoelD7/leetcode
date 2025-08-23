package integer_to_roman

import (
	"strconv"
)

type Table struct {
	subtractiveDigits map[string]struct{}
	subtractiveSymbol map[string]string
	symbol            map[string]string
	multipleSymbol    map[string]string
	unitByPosition    map[int]string
}

func intToRoman(num int) string {
	subtractiveDigits := map[string]struct{}{
		"4": {},
		"9": {},
	}
	subtractiveSymbol := map[string]string{
		"4":   "IV",
		"9":   "IX",
		"40":  "XL",
		"90":  "XC",
		"400": "CD",
		"900": "CM",
	}
	symbol := map[string]string{
		"1":    "I",
		"4":    "IV",
		"5":    "V",
		"9":    "IX",
		"10":   "X",
		"40":   "XL",
		"50":   "L",
		"90":   "XC",
		"100":  "C",
		"400":  "CD",
		"500":  "D",
		"900":  "CM",
		"1000": "M",
	}
	multipleSymbol := map[string]string{
		"2":    "II",
		"3":    "III",
		"6":    "VI",
		"7":    "VII",
		"8":    "VIII",
		"20":   "XX",
		"30":   "XXX",
		"60":   "LX",
		"70":   "LXX",
		"80":   "LXXX",
		"200":  "CC",
		"300":  "CCC",
		"600":  "DC",
		"700":  "DCC",
		"800":  "DCCC",
		"2000": "MM",
		"3000": "MMM",
		"6000": "ML",
		"7000": "MCL",
		"8000": "MCC",
	}
	unitByPosition := map[int]string{
		0: "",
		1: "0",
		2: "00",
		3: "000",
	}

	numStr := strconv.Itoa(num)
	size := len(numStr)
	var digit, unit, digitUnit, roman string
	var ok bool
	var result string
	var pos int

	for ; pos < len(numStr); pos++ {
		size--

		digit = string(numStr[pos])
		unit, _ = unitByPosition[size]
		digitUnit = digit + unit

		if digit == "0" && size > 1 {
			continue
		}

		_, ok = subtractiveDigits[digit]
		if ok {
			roman, ok = subtractiveSymbol[digitUnit]
			if ok {
				result += roman
				continue
			}
		}

		roman, ok = symbol[digitUnit]
		if ok {
			result += roman
			continue
		}

		roman, ok = multipleSymbol[digitUnit]
		if ok {
			result += roman
			continue
		}
	}

	return result
}

func intToRomanRecursive(num int) string {
	table := &Table{
		subtractiveDigits: map[string]struct{}{
			"4": {},
			"9": {},
		},
		subtractiveSymbol: map[string]string{
			"4":   "IV",
			"9":   "IX",
			"40":  "XL",
			"90":  "XC",
			"400": "CD",
			"900": "CM",
		},
		symbol: map[string]string{
			"1":    "I",
			"4":    "IV",
			"5":    "V",
			"9":    "IX",
			"10":   "X",
			"40":   "XL",
			"50":   "L",
			"90":   "XC",
			"100":  "C",
			"400":  "CD",
			"500":  "D",
			"900":  "CM",
			"1000": "M",
		},
		multipleSymbol: map[string]string{
			"2":    "II",
			"3":    "III",
			"6":    "VI",
			"7":    "VII",
			"8":    "VIII",
			"20":   "XX",
			"30":   "XXX",
			"60":   "LX",
			"70":   "LXX",
			"80":   "LXXX",
			"200":  "CC",
			"300":  "CCC",
			"600":  "DC",
			"700":  "DCC",
			"800":  "DCCC",
			"2000": "MM",
			"3000": "MMM",
			"6000": "ML",
			"7000": "MCL",
			"8000": "MCC",
		},
		unitByPosition: map[int]string{
			0: "",
			1: "0",
			2: "00",
			3: "000",
		},
	}

	// 3749
	//MMMDCCXLIX
	numStr := strconv.Itoa(num)
	size := len(numStr)
	var digit, unit, digitUnit, roman string
	var ok bool

	digit = string(numStr[0])              //3
	unit, _ = table.unitByPosition[size-1] //000
	digitUnit = digit + unit               //700

	_, ok = table.subtractiveDigits[digit]
	if ok {
		roman, ok = table.subtractiveSymbol[digitUnit]
		if ok {
			return roman + intToRomanHelper(numStr, table)
		}
	}

	roman, ok = table.symbol[digitUnit]
	if ok {
		return roman + intToRomanHelper(numStr, table)
	}

	roman, ok = table.multipleSymbol[digitUnit]
	if ok {
		return roman + intToRomanHelper(numStr, table)
	}

	return ""
}

// Processes the remaining digits of a number. If the caller was processing 7823, this function will process 823.
func intToRomanHelper(num string, table *Table) string {
	if len(num) == 1 {
		//This has already been processed by an up the stack call
		return ""
	}

	num = num[1:]
	size := len(num)

	roman := ""
	digit := string(num[0])
	unit, _ := table.unitByPosition[size-1]
	digitUnit := digit + unit

	if digit == "0" && size > 1 {
		return intToRomanHelper(num, table)
	}

	_, ok := table.subtractiveDigits[digit]
	if ok {
		roman, ok = table.subtractiveSymbol[digitUnit]
		if ok {
			return roman + intToRomanHelper(num, table)
		}
	}

	roman, ok = table.symbol[digitUnit]
	if ok {
		return roman + intToRomanHelper(num, table)
	}

	roman, ok = table.multipleSymbol[digitUnit]
	if ok {
		return roman + intToRomanHelper(num, table)
	}

	return ""
}
