package trie

import (
    "testing"
    "github.com/davecgh/go-spew/spew"
)

func TestTrie(t *testing.T)  {
    trie := NewTrie()
    trie.Add("abc")
    trie.Add("abcada")
    trie.Add("and")
    spew.Dump(trie)
    spew.Dump(trie.Regexp())
}

