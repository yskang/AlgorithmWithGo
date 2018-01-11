package leetcode

import "log"

type Trie struct {
	val    string
	nexts  map[rune]*Trie
	isWord bool
}

/** Initialize your data structure here. */
func ConstructTire() Trie {
	return Trie{"", make(map[rune]*Trie), false}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if this.Search(word) {
		return
	}
	currentNode := this
	for charIndex, char := range word {
		if _, isExist := currentNode.nexts[char]; !isExist {
			currentNode.nexts[char] = &Trie{string(char), make(map[rune]*Trie), false}
		}
		if charIndex == len(word)-1 {
			currentNode.nexts[char].isWord = true
		}
		currentNode = currentNode.nexts[char]
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	currentNode := this
	for charIndex, char := range word {
		if childTrie, isExist := currentNode.nexts[char]; isExist {
			if charIndex == len(word)-1 && childTrie.isWord {
				return true
			}
			currentNode = childTrie
		} else {
			log.Println("no")
			break
		}
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	currentNode := this
	for _, char := range prefix {
		if childTrie, isExist := currentNode.nexts[char]; isExist {
			currentNode = childTrie
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
