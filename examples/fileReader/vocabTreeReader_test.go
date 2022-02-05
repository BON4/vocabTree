package fileReader

import (
	"fmt"
	"os"
	"testing"
	"vocabTree/trees"
)

func TestConstructTreeFromFile(t *testing.T) {
	f1, err := os.Open("vocabTreeReader.go")
	if err != nil {
		t.Fatal(err)
	}
	f2, err := os.Open("vocabTreeReader_test.go")
	if err != nil {
		t.Fatal(err)
	}
	defer f1.Close()
	defer f2.Close()

	tree, err := ConstructTreeFromFile(&trees.RadixNode{}, f1, Golang)
	if err != nil {
		t.Fatal(err)
	}

	if err := AppendFromFile(tree, f2, Golang); err != nil {
		t.Fatal(err)
	}

	tree.Fprint(os.Stdout)

	tree.FindMatch("L", func(s string) {
		fmt.Println(s)
	})

	fmt.Println(tree.Height())
}
