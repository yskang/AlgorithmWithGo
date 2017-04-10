package tests

import (
	"testing"
	"AlgorithmWithGo/leetcode"
	"AlgorithmWithGo/myLibs"
)

func TestNull(t *testing.T) {
	if nil != leetcode.ConvertBST(nil) {
		t.Error("nil error")
	}
}

func TestOneDepth(t *testing.T) {
	root := myLibs.MakeTreeNode("[5,2,13]")

	if !myLibs.CompareTreeNode(leetcode.ConvertBST(root), myLibs.MakeTreeNode("[18, 20, 13]")) {
		t.Error("one depth error!")
	}
}