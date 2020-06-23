/**
 *
 * @author liangjf
 * @create on 2020/6/23
 * @version 1.0
	https://coolcao.com/2020/04/30/SlidingWindowAlgorithm/
*/
package main

import "fmt"

//给定一个整数数组，计算长度为n的连续子数组的最大和
func maxSubSum(s []int, n int) int {
	if n <= 0 {
		return 0
	}

	if n >= len(s) {
		return len(s)
	}

	curWindowSum := 0
	for i := 0; i < n; i++ {
		curWindowSum += s[i]
	}

	sum := curWindowSum
	for i := n; i < len(s); i++ {
		curWindowSum += s[i] - s[i-n]
		if curWindowSum > sum {
			sum = curWindowSum
		}
	}
	return sum
}

func main() {
	n := 3
	s := []int{-3, 3, 1, -3, 2, 4, 7}

	fmt.Println(maxSubSum(s, n))
}
