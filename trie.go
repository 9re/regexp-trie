package trie

import (
    "fmt"
    "regexp"
    "sort"
    "strings"
    "github.com/davecgh/go-spew/spew"
)

type Trie struct {
    hash *HashTree
}

func (t *Trie) Init() *Trie  {
    t.hash = NewHashTree()
    return t
}

func (t *Trie) Add(s string) {
    h := t.hash
    buildFunc := func(c rune) rune {
        if h.Get(c) == nil {
            h.Put(c, NewHashTree())
        }
        h = h.Get(c)
        return c
    }
    strings.Map(buildFunc, s)
    h.Put(rune(0), NewHashTree())
}

func NewTrie() *Trie {
    return new(Trie).Init()
}

func build(h *HashTree) string {
    if h.Size() == 1 {
        if h.Get(rune(0)) != nil {
            fmt.Println("-0-")
            spew.Dump(h)
            return ""
        }
    }
    var keys []rune
    for k:= range h.nodes {
        keys = append(keys, k)
    }
    sort.Sort(sortRunes(keys))

    var alt []string
    var cc []rune
    q := false
    for _, k := range keys {
        v := h.Get(k)
        if v.Size() == 0 {
            fmt.Println("-1-")
            q = true
        } else {
            if recurse := build(h.Get(k)); recurse != "" {
                alt = append(alt, string(k) + recurse)
            } else {
                cc = append(cc, k)
            }
        }
    }

    cconly := len(alt) == 0
    if len(cc) > 0 {
        if len(cc) == 1 {
            alt = append(alt, string(cc[0]))
        } else {
            alt = append(alt, fmt.Sprintf("[%s]", string(cc)))
        }
    }

    var result string
    if len(alt) == 1 {
        result = alt[0]
    } else {
        result = fmt.Sprintf("(?:%s)", strings.Join(alt, "|"))
    }
    
    if q {
        if cconly {
            result = result + "?"
        } else {
            result = fmt.Sprintf("(?:%s)?", result)
        }
    }
    
    return result
}

func (t *Trie) Regexp() *regexp.Regexp {
    r, _ := regexp.Compile(build(t.hash))
    return r
}

