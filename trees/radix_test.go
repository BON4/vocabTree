package trees

import (
	"fmt"
	"testing"
)

func TestRadixNode_Add(t *testing.T) {
	root := NewRadix()
	root.Add("romane")
	root.Add("romanus")
	root.Add("romulus")
	root.Add("rubens")
	root.Add("ruber")
	root.Add("rubicon")
	root.Add("rubicundus")
	fmt.Println("Done")
}
