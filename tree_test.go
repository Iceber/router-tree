package tree

import (
	"fmt"
	"testing"
)

var paths = []string{
	"/hi",
	"/b/",
	"/ABC/",
	"/search/:query",
	"/cmd/:tool/",
	"/src/*filepath",
	"/x",
	"/x/y",
	"/y/",
	"/y/z",
	"/0/:id",
	"/0/:id/1",
	"/1/:id/",
	"/1/:id/2",
	"/aa",
	"/a/",
	"/doc",
	"/doc/go_faq.html",
	"/doc/go1.html",
	"/doc/go/away",
	"/no/a",
	"/no/b",
	"/Π",
	"/u/apfêl/",
	"/u/äpfêl/",
	"/u/öpfêl",
	"/v/Äpfêl/",
	"/v/Öpfêl",
	"/w/♬",  // 3 byte
	"/w/♭/", // 3 byte, last byte differs
	"/w/𠜎",  // 4 byte
	"/w/𠜏/", // 4 byte
}

var nodeTypeStr = map[nodeType]string{root: "root", static: "static"}

func printNodeWithChildren(prefix string, n *node) {
	var s string
	if n.handle == nil {
		s = ""
	} else {
		s = string(*n.handle)
	}
	fmt.Printf("%s path:%s, nodeType:%s,indices:%s, handleString: %s\n", prefix, n.path, nodeTypeStr[n.nType], string(n.indices), s)

	prefix += "    "
	for i := range n.children {
		printNodeWithChildren(prefix, n.children[i])
	}
}

var tree = NewTree()

var mapTree = make(map[string]Handle)

func TestMain(m *testing.M) {
	for _, path := range paths {
		handle := path
		tree.AddRoute(path, Handle(&handle))
		mapTree[path] = Handle(&path)
	}
	//	printNodeWithChildren("", tree.root)
	m.Run()
}

func BenchmarkTreeGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, path := range paths {
			tree.GetValue(path)
		}
	}
}

func BenchmarkMapGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, path := range paths {
			_, _ = mapTree[path]
		}
	}
}
