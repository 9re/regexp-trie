package trie

type HashTree struct {
    nodes map[rune]*HashTree
}

// initializer 
func (h *HashTree) Init() *HashTree {
    h.nodes = make(map[rune]*HashTree)
    return h
}

func (h *HashTree) Get(key rune) *HashTree {
    return h.nodes[key]
}

func (h *HashTree) Put(key rune, value *HashTree) {
    h.nodes[key] = value
}

func (h *HashTree) Size() int {
    return len(h.nodes)
}

func NewHashTree() *HashTree {
    return new(HashTree).Init()
}

