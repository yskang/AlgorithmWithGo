package tests

import (
	"testing"
	"AlgorithmWithGo/cj2017_01"
	"AlgorithmWithGo/cj2017_02"
	"AlgorithmWithGo/cj2017_03"
	"AlgorithmWithGo/cj2017_04"
	"AlgorithmWithGo/cj2017_05"
)

func TestCJ2017_1(t *testing.T) {
	cj2017_01.CJ_2017_01()
}

func TestCJ2017_2(t *testing.T) {
	cj2017_02.CJ_2017_02()
}

func TestCJ2017_3(t *testing.T) {
	cj2017_03.CJ_2017_03()
	//cj2017_03.IsValid(1000, 1)
}

func TestCJ2017_3_Goroutine(t *testing.T) {
	cj2017_03.CJ_2017_03_with_goroutine()
}

func TestCJ2017_4(t *testing.T) {
	cj2017_04.CJ_2017_04()
}

func TestCJ2017_5(t *testing.T) {
	cj2017_05.CJ_2017_05()
}