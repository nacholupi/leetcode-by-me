package trie

type Trie struct {
	children [26]*Trie
	isWord   bool
}

func Constructor() Trie {
	return Trie{}
}

func (r *Trie) Insert(word string) {
	node := r
	for _, c := range word {
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = &Trie{}
		}
		node = node.children[c-'a']
	}
	node.isWord = true
}

func (r *Trie) Search(word string) bool {
	node := r
	for _, c := range word {
		node = node.children[c-'a']
		if node == nil {
			return false
		}
	}
	return node.isWord
}

func (r *Trie) StartsWith(prefix string) bool {
	node := r
	for _, c := range prefix {
		node = node.children[c-'a']
		if node == nil {
			return false
		}
	}
	return true
}
