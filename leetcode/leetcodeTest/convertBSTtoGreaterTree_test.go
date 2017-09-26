package leetcodeTest

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

func TestConvert(t *testing.T) {
	root := myLibs.MakeTreeNode("5,2,13")

	if !myLibs.CompareTreeNode(leetcode.ConvertBST(root), myLibs.MakeTreeNode("18, 20, 13")) {
		t.Error("one depth error!", root)
	}

	root = myLibs.MakeTreeNode("5,3,7,1,4,6,8")

	if !myLibs.CompareTreeNode(leetcode.ConvertBST(root), myLibs.MakeTreeNode("26,33,15,34,30,21,8")) {
		t.Error("one depth error!", root)
	}
}