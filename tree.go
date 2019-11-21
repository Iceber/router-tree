package tree

type nodeType uint8

const (
	static nodeType = iota
	root
)

type Handle *string // 实际为 httprouter.Handle

type Tree struct {
	root *node
}

func NewTree() *Tree {
	return &Tree{new(node)}
}

func (t *Tree) AddRoute(path string, handle Handle) {
	t.root.addRoute(path, handle)
}

func (t *Tree) GetValue(path string) Handle {
	return t.root.getValue(path)
}
