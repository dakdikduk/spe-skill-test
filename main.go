package main

import (
	"fmt"
	"math"
)

func main() {
	// NarcissisticNumber
	fmt.Println(NarcissisticNumber(1634)) // true
	fmt.Println(NarcissisticNumber(153))  // true
	fmt.Println(NarcissisticNumber(111))  // false

	//ParityOutlier
	fmt.Println(ParityOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36})) // 11 true
	fmt.Println(ParityOutlier([]int{160, 3, 1719, 19, 11, 13, -21})) //160 true
	fmt.Println(ParityOutlier([]int{11, 13, 15, 19, 9, 13, -21}))    // 0 false

	//NeedleInTheHaysstack
	fmt.Println(FindNeedleInHaystack([]string{"red", "blue", "yellow", "black", "grey"}, "blue")) // 1

	//BlueOceanReverse
	first := []int{1, 2, 3, 4, 6, 10}
	second := []int{1}
	fmt.Println(BlueOceanReverse(first, second))
	first = []int{1, 5, 5, 5, 5, 3}
	second = []int{5}
	fmt.Println(BlueOceanReverse(first, second))

}

func NarcissisticNumber(num int) bool {
	oldNum := num
	total := 0
	length := 0
	var helper []int
	for num/10 > 0 {
		length++
		helper = append(helper, num%10)
		num = num / 10
	}
	length++
	helper = append(helper, num%10)

	for _, v := range helper {
		total = total + int(math.Pow(float64(v), float64(length)))
	}
	return total == oldNum
}

func ParityOutlier(nums []int) (int, bool) {
	evenCount, oddCount := 0, 0
	lastOdd, lastEven := 0, 0
	for _, v := range nums {
		if v%2 == 0 {
			evenCount++
			lastEven = v
		} else if v%2 == 1 {
			oddCount++
			lastOdd = v
		}
	}

	//assuming the sentence below always true
	// the array is either entirely contains whole of odd integers with 1 outlier even integer or whole of even integers with 1 outlier odd integer.
	if evenCount >= 2 && oddCount == 1 { // odd outlier
		return lastOdd, true
	} else if oddCount >= 2 && evenCount == 1 { // even outlier
		return lastEven, true
	}

	return 0, false
}

func FindNeedleInHaystack(haystack []string, needle string) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}

	return -1 //not found
}

func BlueOceanReverse(first []int, second []int) []int {
	var result []int

	secondMap := make(map[int]bool)
	for _, v := range second {
		secondMap[v] = true
	}
	for _, num := range first {
		if !secondMap[num] {
			result = append(result, num)
		}
	}

	return result
}
