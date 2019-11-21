package tree

type node struct {
	path     string
	nType    nodeType
	handle   Handle
	children []*node
	indices  []byte
}

func (*node) incrementChild(i int) int {
	return i
}

// 通过 path索引i 生成新的子节点
func (n *node) bifurcate(i int) {
	child := node{
		path:   n.path[i:],
		nType:  static,
		handle: n.handle,
	}
	n.indices = []byte{child.path[0]}
	n.children = []*node{&child}
	n.path = n.path[:i]
	n.handle = nil
}

// 匹配首字母来获取相应子节点
func (n *node) getChildNode(idxc byte) *node {
	for i := 0; i < len(n.indices); i++ {
		if idxc == n.indices[i] {
			i = n.incrementChild(i)
			return n.children[i]
		}
	}

	return nil
}

func (n *node) addRoute(path string, handle Handle) {
	fullPath := path

	if len(n.path) == 0 && len(n.indices) == 0 {
		n.insertChild(path, handle)
		n.nType = root
		return
	}

	for {
		i := longestCommonPrefix(path, n.path)

		// 部分匹配，需要取出匹配部分作为父节点
		if i < len(n.path) {
			n.bifurcate(i)
		}

		// path 长度大于匹配部分
		if i < len(path) {
			path = path[i:] // 未匹配部分

			idxc := path[0]
			childNode := n.getChildNode(idxc)
			if childNode != nil {
				n = childNode
				continue
			}

			// 创建新的子节点
			n.indices = append(n.indices, idxc)
			child := &node{}
			n.children = append(n.children, child)
			n.incrementChild(len(n.indices) - 1)

			n = child
			n.insertChild(path, handle)
			return
		}

		if n.handle != nil {
			panic("a handle is already registered for path '" + fullPath + "'")
		}
		n.handle = handle
		return
	}
}

func (n *node) insertChild(path string, handle Handle) {
	n.path = path
	n.handle = handle
}

func (n *node) getValue(path string) (handle Handle) {
	for {
		prefix := n.path

		if len(path) <= len(n.path) {
			if prefix == path {
				handle = n.handle
			}
			return
		}

		if path[:len(prefix)] == prefix {
			path = path[len(prefix):]

			idxc := path[0]
			for i, c := range n.indices {
				if c == idxc {
					n = n.children[i]
					prefix = n.path
					continue
				}
			}
			return
		}
	}
}
