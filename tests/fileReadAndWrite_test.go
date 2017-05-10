package tests

import (
	"testing"
	"os"
	"fmt"
	"bufio"
	"strings"
)

func TestReadInputFile(t *testing.T) {
	pathName := "/home/yskang/Downloads/p5_inputs/"
	fileName := "input001.txt"

	fullName := pathName + fileName

	inputs := readLine(fullName)

	inputData := strings.Split(inputs, "\n")

	for _, data := range inputData {
		fmt.Println(data)
	}
}

func readLine(path string) string {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	inputs := ""

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		inputs += scanner.Text() + "\n"
	}
	return  inputs
}

func TestWriteOutputFile(t *testing.T) {
	pathName := "/home/yskang/Downloads/p5_inputs/"
	fileName := "out001.txt"

	outFileName := pathName + fileName

	outFile, err := os.Create(outFileName)

	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()

	outFile.Write([]byte{255,120,60,33})
	outFile.WriteString("\n")
	for i := 0 ; i < 10 ; i++ {
		outFile.WriteString(fmt.Sprintf("%d\n", i))
	}
}