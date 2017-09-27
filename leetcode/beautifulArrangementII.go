package leetcode

func ConstructArray(n int, k int) []int {
	return constructArray(n, k)
}

func constructArray(n int, k int) []int {
	arr := make([]int, 0)
	rest := make([]bool, n+1)
	arr = append(arr, 1)
	rest[1] = true
	for diff := k ; diff > 0 ; diff -= 1 {
		if len(arr) % 2 == 1 {
			arr = append(arr, arr[len(arr)-1] + diff)
			rest[arr[len(arr)-1]] = true
		} else {
			arr = append(arr, arr[len(arr)-1] - diff)
			rest[arr[len(arr)-1]] = true
		}
	}

	for len(arr) < n {
		for i := 1 ; i <= n ; i++ {
			if rest[i] == false {
				arr = append(arr, i)
				rest[i] = true
			}
		}
	}

	return arr
}

