package trees

import (
	"testing"
)

func TestTireNode_Add(t *testing.T) {
	n := NewTire()

	n.Add("max")
	n.Add("min")
	n.Add("mean")
	n.Add("marge")
}

func TestTireNode_Insert(t *testing.T) {
	n := NewTire()

	n.Add("max")
	n.Add("min")
	n.Add("mean")
	n.Add("marge")
}
