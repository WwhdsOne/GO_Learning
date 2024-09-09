package main

import "fmt"

func romanToInt(s string) int {
	ans := 0
	for index, c := range s {
		value := toRoman(c)
		if index < len(s)-1 && toRoman(rune(s[index+1])) > value {
			ans -= value
		} else {
			ans += value
		}
	}
	return ans
}

func toRoman(c rune) int {
	switch c {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}

func main() {
	str := "MCMXCIV"
	fmt.Println(romanToInt(str))
}
