package main

import "fmt"

func main() {
	var result int32
	result = sherlockAndAnagrams("ifailuhkqq")
	fmt.Println(result)
	result = sherlockAndAnagrams("abcd")
	fmt.Println(result)
}

func sherlockAndAnagrams(s string) int32 {
	var anagrams int32 = 0
	target := []rune(s)
	long := len(target) // s string length
	for i := 1; i < long; i++ {
		for x := 0; x < long-i; x++ {
			m1 := make(map[rune]int)

			for _, v := range string(target[x : x+i]) { //target
				m1[v] += 1
			}
			for y := x + 1; y <= long-i; y++ {
				m2 := make(map[rune]int)
				for _, v := range string(target[y : y+i]) { //compare
					m2[v] += 1
				}
				var isAn bool = true
				for i, _ := range m1 { //compare each rune int
					if m1[i] != m2[i] {
						isAn = false
						break
					}
				}
				if isAn {
					anagrams++
				}
			}
		}
	}
	return anagrams
}
