package main

import "fmt"

//func myLengthOfNonRepeatingSubStr(s string) int {
//	dict := make(map[rune]int)
//	start, max := 0, 0
//	for i, v := range []rune(s) {
//		if dict[v] > start {
//			start = dict[v]
//		}
//		if i-start+1 > max {
//			max = i - start + 1
//		}
//		dict[v] = i + 1
//	}
//	return max
//}

var lastOccurred = make([]int, 0xffff) // 65535

func lengthOfNonRepeatingSubStr(s string) int {

	for i := range lastOccurred {
		lastOccurred[i] = -1
	}
	// lastOccurred[0x65] = 1
	// lastOccurred[0x8BFE] = 6

	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI := lastOccurred[ch]; lastI != -1 && lastI >= start {
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
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))
	fmt.Println(lengthOfNonRepeatingSubStr("这里是慕课网"))
	fmt.Println(lengthOfNonRepeatingSubStr("一二三二一"))
	fmt.Println(lengthOfNonRepeatingSubStr("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
