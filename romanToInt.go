package main

// Not 100% the greatest solution, beats around 60% at 6ms.
// Could probably be trimmed down a little bit, but ... moving on.
var numerals map[rune]int = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	sum := 0

	c_i := 1
	var c_p rune

	for _, c := range s {
		if c == c_p {
			c_i++
		} else {
			prev_val := numerals[c_p]
			curr_val := numerals[c]

			if curr_val > prev_val {
				sum -= prev_val * c_i
			} else {
				sum += prev_val * c_i
			}

			c_i = 1
		}

		c_p = c
	}

	if c_i > 0 {
		sum += numerals[c_p] * c_i
	}

	return sum
}
