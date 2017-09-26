package leetcodeTest

import (
	"testing"
	"AlgorithmWithGo/leetcode"
	"fmt"
)

func TestExample12(t *testing.T) {
	result := leetcode.NextGreaterElementIII(12)

	if result != 21 {
		t.Error("example 12 error!", result)
	}
	fmt.Println("test End")
}

func TestExample21(t *testing.T) {
	result := leetcode.NextGreaterElementIII(21)

	if result != -1 {
		t.Error("example 21 error!", result)
	}
	fmt.Println("test End")
}

func TestExampleAscendingOrder(t *testing.T) {
	result := leetcode.NextGreaterElementIII(12345)

	if result != 12354 {
		t.Error("example ascending order error!", result)
	}
	fmt.Println("test End")
}

func TestExampleDescendingOrder(t *testing.T) {
	result := leetcode.NextGreaterElementIII(54321)

	if result != -1 {
		t.Error("example descending order error!", result)
	}
	fmt.Println("test End")
}

func TestExampleTheOtherNumber(t *testing.T) {
	result := leetcode.NextGreaterElementIII(534976)

	if result != 536479 {
		t.Error("example the other number error!", result)
	}
	fmt.Println("test End")
}

func TestExampleTheOtherNumber2(t *testing.T) {
	result := leetcode.NextGreaterElementIII(534973)

	if result != 537349 {
		t.Error("example the other number error!", result)
	}
	fmt.Println("test End")
}

func TestBiggerThan32bit(t *testing.T) {
	result := leetcode.NextGreaterElementIII(1999999999)

	if result != -1 {
		t.Error("example big number error!", result)
	}
	fmt.Println("test End")
}