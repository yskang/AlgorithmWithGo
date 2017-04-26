package leetcode

type NumArrayII struct {
	numArray []int
}

func NumArrayConstructor(nums []int) NumArrayII {
	return NumArrayII{nums}
}

func (this *NumArrayII) Update(i int, val int)  {
	this.numArray[i] = val
}

func (this *NumArrayII) SumRangeII(i int, j int) int {
	sum := 0
	for k := i ; k <= j ; k++ {
		sum += this.numArray[k]
	}
	return sum
}

