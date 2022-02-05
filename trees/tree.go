package trees

import (
	"fmt"
	"io"
)

type Tree struct {
	Node
}

func NewTree(treeType Node) Tree {
	switch treeType.(type) {
	case *RadixNode:
		return Tree{NewRadix()}
	case *TireNode:
		return Tree{NewTire()}
	default:
		panic("Tree of this type is not implemented")
	}
}

func (tree *Tree) TraverseCond(cond Condition, viewer Viewer) {
	traverseCond(tree, cond, viewer)
}

func (tree *Tree) FindMatch(str string, viewer Viewer) {
	findMatch(tree, str, viewer)
}

func (tree *Tree) Fprint(w io.Writer) {
	fmt.Fprintf(w, "#\n")
	fprint(tree, w, false, "")
}

func (tree *Tree) Height() int {
	return height(tree, 0)
}
