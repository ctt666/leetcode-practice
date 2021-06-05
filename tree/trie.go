package tree

type Trie struct {
	IsWord   bool
	Children [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, v := range word {
		v -= 'a'
		if node.Children[v] == nil {
			node.Children[v] = &Trie{}
		}
		node = node.Children[v]
	}
	node.IsWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	//node := this
	//for _, v := range word {
	//	if node.Children[v - 'a'] == nil {
	//		return false
	//	}
	//	node = node.Children[v - 'a']
	//}
	//return node.IsWord
	node := this.SearchPrefix(word)
	return node != nil && node.IsWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	//node := this
	//for _, v := range prefix {
	//	if node.Children[v - 'a'] == nil {
	//		return false
	//	}
	//	node = node.Children[v - 'a']
	//}
	//return true
	return this.SearchPrefix(prefix) != nil
}

func (this *Trie) SearchPrefix(prefix string) *Trie {
	node := this
	for _, v := range prefix {
		if node.Children[v-'a'] == nil {
			return nil
		}
		node = node.Children[v-'a']
	}
	return node
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
