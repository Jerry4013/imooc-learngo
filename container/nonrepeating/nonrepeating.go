package main

import "fmt"

func myLengthOfNonRepeatingSubStr(s string) int {
	dict := make(map[rune]int)
	start, max := 0, 0
	for i, v := range []rune(s) {
		if dict[v] > start {
			start = dict[v]
		}
		if i-start+1 > max {
			max = i - start + 1
		}
		dict[v] = i + 1
	}
	return max
}

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(myLengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(myLengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(myLengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(myLengthOfNonRepeatingSubStr(""))
	fmt.Println(myLengthOfNonRepeatingSubStr("b"))
	fmt.Println(myLengthOfNonRepeatingSubStr("abcdef"))
	fmt.Println(myLengthOfNonRepeatingSubStr("这里是慕课网"))
	fmt.Println(myLengthOfNonRepeatingSubStr("一二三二一"))
	fmt.Println(myLengthOfNonRepeatingSubStr("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
