package trees

import (
	"fmt"
	"io"
)

type Condition func(Node) bool
type Viewer func(s string)

type Node interface {
	Add(string)
	Val() string
	Complete() bool
	ChildMap(func(Node))
	Len() int
}

func traverseCond(root Node, cond Condition, viewer Viewer) {
	if cond(root) && len(root.Val()) > 0 {
		viewer(root.Val())
	}

	root.ChildMap(func(n Node) {
		traverseCond(n, cond, viewer)
	})
}

//findMatch - apply viewer to matching complete words
func findMatch(root Node, str string, viewer Viewer) {
	findMatchHelper(root, str, "", viewer)
}

func findMatchHelper(root Node, str string, accum string, viewer Viewer) {
	if root.Complete() {
		viewer(accum)
	}

	root.ChildMap(func(n Node) {
		delim := FindDelimeter(str, n.Val())
		if delim != 0 || len(str) == 0 {
			findMatchHelper(n, str[delim:], accum+n.Val(), viewer)
		}
	})
}

func height(v Node, c int) int {
	heights := make([]int, v.Len())
	i := 0
	v.ChildMap(func(n Node) {
		heights[i] = height(n, c+1)
		i++
	})

	if v.Len() == 0 {
		return c
	}
	return max(heights)
}

func fprint(tree Node, w io.Writer, root bool, padding string) {
	if tree == nil {
		return
	}

	index := 0
	tree.ChildMap(func(n Node) {
		fmt.Fprintf(w, "%s%s\n", padding+getPadding(root, getBoxType(index, tree.Len())), n.Val())
		fprint(n, w, false, padding+getPadding(root, getBoxTypeExternal(index, tree.Len())))
		index++
	})
}

type boxType int

const (
	Regular boxType = iota
	Last
	AfterLast
	Between
)

func (boxType boxType) String() string {
	switch boxType {
	case Regular:
		return "\u251c" // ├
	case Last:
		return "\u2514" // └
	case AfterLast:
		return " "
	case Between:
		return "\u2502" // │
	default:
		panic("invalid box type")
	}
}

func getBoxType(index int, len int) boxType {
	if index+1 == len {
		return Last
	} else if index+1 > len {
		return AfterLast
	}
	return Regular
}

func getBoxTypeExternal(index int, len int) boxType {
	if index+1 == len {
		return AfterLast
	}
	return Between
}

func getPadding(root bool, boxType boxType) string {
	if root {
		return ""
	}

	return boxType.String() + " "
}
