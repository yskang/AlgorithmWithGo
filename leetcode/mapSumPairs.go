package leetcode

import "strings"

type MapSum struct {
	pair map[string]int
}


/** Initialize your data structure here. */
func ConstructorMap() MapSum {
	return MapSum{make(map[string]int)}
}


func (this *MapSum) Insert(key string, val int)  {
	this.pair[key] = val
}


func (this *MapSum) Sum(prefix string) int {
	sum := 0
	for key, value := range this.pair {
		if strings.HasPrefix(key, prefix) {
			sum += value
		}
	}
	return sum
}


/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */