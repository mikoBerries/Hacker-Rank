package main

import "fmt"

func main() {
	n := int32(5)
	q := make([][]int32, 3, 3)
	q[0] = []int32{1, 2, 100}
	q[1] = []int32{2, 5, 100}
	q[2] = []int32{3, 4, 100}
	max := arrayManipulation(n, q)
	fmt.Println(max)
}
func arrayManipulation(n int32, queries [][]int32) int64 {
	// Write your code here
	var max int64
	arr := make([]int64, n+1)
	for _, v := range queries { //marking
		arr[v[0]-1] += int64(v[2])
		arr[v[1]] -= int64(v[2])
	}
	for i := int32(1); i <= n-1; i++ { //sum arr
		arr[i] += arr[i-1]
		if max < arr[i] {
			max = arr[i]
		} //max check
	}
	return max
}
