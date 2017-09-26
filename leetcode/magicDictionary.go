package leetcode

type MagicDictionary struct {
	words []string
}


/** Initialize your data structure here. */
func ConstructorMagicDic() MagicDictionary {
	return MagicDictionary{make([]string, 0)}
}


/** Build a dictionary through a list of words */
func (this *MagicDictionary) BuildDict(dict []string)  {
	this.words = dict
}


/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (this *MagicDictionary) Search(word string) bool {
	for _, w := range this.words {
		if len(w) == len(word) {
			if w != word {
				diffCount := 0
				for i := range w {
					if w[i] != word[i] {
						diffCount += 1
					}
				}
				if diffCount <= 1{
					return true
				}
			}
		}
	}
	return false
}


/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dict);
 * param_2 := obj.Search(word);
 */
