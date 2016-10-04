package main

import (
	"testing"
	"fmt"
)

func TestBigAdd(t *testing.T) {
	simple := BigAdd("100000", "1")
	carry := BigAdd("9999", "2")
	nn := BigAdd("-1", "-1")
	np := BigAdd("-1", "10")
	pn := BigAdd("10", "-1")
	zero := BigAdd("0000", "0000")

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
	posVal, pos := GetBigInt("123456789")
	negVal, neg := GetBigInt("-123456789")
	zero, pos := GetBigInt("000000")
	zeroNum, pos := GetBigInt("000100")

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
	equal := CompareAandB("12345", "12345")
	bigA := CompareAandB("12346", "12345")
	bigB := CompareAandB("12345", "12346")

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
	simple := BigSub("2", "1")
	zero := BigSub("123", "123")
	carry := BigSub("10000000", "9")
	negative := BigSub("1", "100000")
	nn := BigSub("-123456", "-654321")
	np := BigSub("-123456", "123456")
	pn := BigSub("123456", "-123456")

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
	simple := BigMul("12345", "54321")
	zero := BigMul("12345", "0")
	minus := BigMul("-12345", "54321")

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
	simple, remainder := BigDiv("1", "1")
	large, largeRe := BigDiv("987654321", "63748")
	minus, minusRe := BigDiv("987654321", "-63748")

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