package trees

import (
	"fmt"
	"os"
	"testing"
)

func TestNode_TraverseCond(t *testing.T) {
	tree := NewTree(&RadixNode{})

	tree.Add("max")
	tree.Add("min")
	tree.Add("mean")
	tree.Add("marge")

	tree.TraverseCond(func(n Node) bool {
		return true
	}, func(s string) {
		fmt.Print(s)
	})
}

func TestRadixNode_FindMatch(t *testing.T) {
	tree := NewTree(&RadixNode{})

	tree.Add("romane")
	tree.Add("romanus")
	tree.Add("romulus")
	tree.Add("rubens")
	tree.Add("ruber")
	tree.Add("rubicon")
	tree.Add("rubicundus")

	tree.FindMatch("ru", func(s string) {
		fmt.Println(s)
	})
}

func TestTireNode_FindMatch(t *testing.T) {
	tree := NewTree(&TireNode{})

	tree.Add("romane")
	tree.Add("romanus")
	tree.Add("romulus")
	tree.Add("rubens")
	tree.Add("ruber")
	tree.Add("rubicon")
	tree.Add("rubicundus")

	tree.FindMatch("ru", func(s string) {
		fmt.Println(s)
	})
}

func TestRadixNode_Fprint(t *testing.T) {
	tree := NewTree(&RadixNode{})

	tree.Add("romane")
	tree.Add("romanus")
	tree.Add("romulus")
	tree.Add("rubens")
	tree.Add("ruber")
	tree.Add("rubicon")
	tree.Add("rubicundus")

	tree.Fprint(os.Stdout)
}

func TestTireNode_Fprint(t *testing.T) {
	tree := NewTree(&TireNode{})

	tree.Add("romane")
	tree.Add("romanus")
	tree.Add("romulus")
	tree.Add("rubens")
	tree.Add("ruber")
	tree.Add("rubicon")
	tree.Add("rubicundus")

	tree.Fprint(os.Stdout)
}
