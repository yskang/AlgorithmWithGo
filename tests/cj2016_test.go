package tests

import (
	"testing"
	"AlgorithmWithGo/cj2016_01"
	"AlgorithmWithGo/cj2016_02"
	"AlgorithmWithGo/cj2016_03"
	"AlgorithmWithGo/cj2016_04"
)

func TestCJ2016_1(t *testing.T) {
	cj2016_01.CJ_2016_01()
}

func TestCJ2016_2(t *testing.T) {
	cj2016_02.CJ_2016_02()
}

func TestCJ2016_3(t *testing.T) {
	cj2016_03.CJ_2016_03()
	//cj2016_03.IsValid(1000, 1)
}

func TestCJ2016_4(t *testing.T) {
	cj2016_04.CJ_2016_04()
}