package tests

import (
	"testing"
	"fmt"
	"AlgorithmWithGo/myLibs"
)

func TestBigAdd(t *testing.T) {
	simple := myLibs.BigAdd("100000", "1")
	carry := myLibs.BigAdd("9999", "2")
	nn := myLibs.BigAdd("-1", "-1")
	np := myLibs.BigAdd("-1", "10")
	pn := myLibs.BigAdd("10", "-1")
	zero := myLibs.BigAdd("0000", "0000")

	if zero != "0" {
		t.Error("zero error!", zero)
	}

	if simple != "100001" {
		t.Fail()
	}

	if carry != "10001" {
		t.Fail()
	}

	if nn != "-2" {
		t.Error("negative negative", nn)
	}

	if np != "9" {
		t.Error("negative positive", np)
	}

	if pn != "9" {
		t.Error("positive negative", np)
	}
}

func TestGetBigInt(t *testing.T) {
	posVal, pos := myLibs.GetBigInt("123456789")
	negVal, neg := myLibs.GetBigInt("-123456789")
	zero, pos := myLibs.GetBigInt("000000")
	zeroNum, pos := myLibs.GetBigInt("000100")

	if zeroNum != "100" {
		t.Error("zeroNum error", zeroNum)
	}

	if zero != "0" {
		t.Error("zero error", zero)
	}

	if posVal != "123456789" {
		t.Error("value error : ", posVal)
	}

	if pos != true {
		t.Error("sign error", pos)
	}

	if negVal != "123456789" {
		t.Error("value error : ", negVal)
	}

	if neg != false {
		t.Error("sign error", neg)
	}
}

func TestCompareAandB(t *testing.T) {
	equal := myLibs.CompareAandB("12345", "12345")
	bigA := myLibs.CompareAandB("12346", "12345")
	bigB := myLibs.CompareAandB("12345", "12346")

	if equal != "=" {
		t.Error("equal failure")
	}

	if bigA != ">" {
		t.Error("big A failure")
	}

	if bigB != "<" {
		t.Error("big A failure")
	}
}

func TestBigSub(t *testing.T) {
	simple := myLibs.BigSub("2", "1")
	zero := myLibs.BigSub("123", "123")
	carry := myLibs.BigSub("10000000", "9")
	negative := myLibs.BigSub("1", "100000")
	nn := myLibs.BigSub("-123456", "-654321")
	np := myLibs.BigSub("-123456", "123456")
	pn := myLibs.BigSub("123456", "-123456")

	if simple != "1" {
		t.Error("Error!", simple)
	}

	if zero != "0" {
		t.Error("Error!", zero)
	}

	if carry != "9999991" {
		t.Error("Error!", carry)
	}

	if negative != "-99999" {
		t.Error("Negative!", negative)
	}

	if nn != "530865" {
		t.Error("Negative - Negative", nn)
	}

	if np != "-246912" {
		t.Error("Negative - positive", np)
	}

	if pn != "246912" {
		t.Error("Positive - Negative", np)
	}
}

func TestBigMul(t *testing.T) {
	simple := myLibs.BigMul("12345", "54321")
	zero := myLibs.BigMul("12345", "0")
	minus := myLibs.BigMul("-12345", "54321")

	if minus != "-670592745" {
		t.Error("minus error!", minus)
	}

	if simple != "670592745" {
		t.Error("simple error!", simple)
	}

	if zero != "0" {
		t.Error("zero error!", zero)
	}
}

func TestBigDiv(t *testing.T) {
	simple, remainder := myLibs.BigDiv("1", "1")
	large, largeRe := myLibs.BigDiv("987654321", "63748")
	minus, minusRe := myLibs.BigDiv("987654321", "-63748")

	if simple != "1" || remainder != "0" {
		t.Error("simple Error!", simple, remainder)
	}

	if large != "15493" || largeRe != "6557" {
		t.Error("large Error!", large, largeRe)
	}

	if minus != "-15493" || minusRe != "6557" {
		t.Error("minus Error!", minus, minusRe)
	}
}

func TestMyStack_Clear(t *testing.T) {
	s := " "
	fmt.Println(s[0])
}