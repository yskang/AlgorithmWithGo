package leetcode

import (
	"strconv"


	"fmt"
	"math"
)

func FindMaximumXOR(nums []int) int {
	return findMaximumXOR(nums)
}

func findMaximumXOR(nums []int) int {
	fmt.Println("Start for", nums)

	// find maximum number
	maxNum := 0
	for _, num := range nums {
		if maxNum < num {
			maxNum = num
		}
	}

	l := 0
	for i := 0 ; i < 32 ; i++ {
		maxNum = maxNum >> 1
		if maxNum == 0 {
			l = i
			break
		}
	}

	fmt.Println("max order ", l)

	maximums := make([]int, 0)
	for _, num := range nums {
		if num >= int(math.Pow(2, float64(l))) {
			maximums = append(maximums, num)
		}
	}
	fmt.Println("maximums", maximums)

	maxXORs := make([]int, 0)

	for _, maxN := range maximums {
		fmt.Println("find for ", maxN, strconv.FormatInt(int64(maxN), 2))
		// find position of 0
		maxNumBin := strconv.FormatInt(int64(maxN), 2)
		comNumStr := ""
		for _, c := range maxNumBin {
			if c == '0' {
				comNumStr += "1"
			} else {
				comNumStr += "0"
			}
		}

		reverseFilter, _ := strconv.ParseInt(comNumStr, 2, 64)

		fmt.Println("empty filter", reverseFilter, comNumStr)

		// find maximum number of position 0 to 1
		minNum := maxN
		maximumCommon := int32(0)
		maxNums := make([]int32, 0)
		for _, num := range nums {
			temp := int32(num) & int32(reverseFilter)
			if temp > maximumCommon {
				maximumCommon = temp
				maxNums = []int32{int32(num)}
			} else if temp == maximumCommon {
				maxNums = append(maxNums, int32(num))
			}

			if minNum > num {
				minNum = num
			}
		}
		fmt.Println("maximum common", maximumCommon, strconv.FormatInt(int64(maximumCommon), 2))
		fmt.Println("maxNums", maxNums)

		if maximumCommon == 0 {
			maxXORs = append(maxXORs, maxN ^ minNum)
			fmt.Println("append", maxN ^ minNum, maxN, "xor", minNum)
		} else {
			mNum := int32(math.MaxInt32)
			for _, m := range maxNums {
				temp := maximumCommon ^ int32(m)
				if temp < mNum {
					mNum = temp
				}
			}
			mNum = mNum | maximumCommon
			maxXORs = append(maxXORs, maxN ^ int(mNum))
			fmt.Println("append", maxN ^ int(mNum), maxN, "xor", mNum)
		}
	}

	fmt.Println("XORs", maxXORs)
	maxXOR := 0

	for _, xor := range maxXORs {
		if xor > maxXOR {
			maxXOR = xor
		}
	}
	return maxXOR
}

func FindMaximumXOR2(nums []int) int {
	max := 0
	for _, num := range nums {
		for _, num2 := range nums {
			if num ^ num2 > max {
				fmt.Println("new max", num ^ num2, num, "xor", num2)
				max = num ^ num2
			}
		}
	}
	return max
}