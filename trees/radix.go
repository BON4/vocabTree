package trees

type RadixNode struct {
	Vocab     string
	Comp      bool
	ChildList []*RadixNode
}

func NewRadix() Node {
	return &RadixNode{
		Vocab:     "",
		Comp:      false,
		ChildList: []*RadixNode{},
	}
}

func (v *RadixNode) Val() string {
	return v.Vocab
}

func (v *RadixNode) Complete() bool {
	return v.Comp
}

func (v *RadixNode) Len() int {
	return len(v.ChildList)
}

func (v *RadixNode) ChildMap(f func(n Node)) {
	for i := 0; i < len(v.ChildList); i++ {
		f(v.ChildList[i])
	}
}

func (v *RadixNode) Add(word string) {
	foundNode := false

	for i, child := range v.ChildList {
		foundSubstringLen := FindDelimeter(child.Vocab, word)
		if foundSubstringLen > 0 {
			foundNode = true
			v.ChildList[i] = divideAndCreateNode(child, foundSubstringLen)
			if foundSubstringLen < len(word) {
				v.ChildList[i].Add(word[foundSubstringLen:])
			}
		}
	}

	if !foundNode {
		var newNode = &RadixNode{Vocab: word, ChildList: []*RadixNode{}, Comp: true}
		v.ChildList = append(v.ChildList, newNode)
	}
}

//divideAndCreateNode - divide RadixNode at specific place, and creates two separate nodes.
//If second RadixNode exists all childs will be asign to second RadixNode, otherwise childs remains
//("test" transforms to "te" -> "st", where delimEt == 2)
//("test" transforms to "test" -> [], where delimEt == len(Vocab))
func divideAndCreateNode(n *RadixNode, delimEt int) *RadixNode {
	if len(n.Vocab[delimEt:]) > 0 {
		newNode := &RadixNode{Vocab: n.Vocab[:delimEt], Comp: false}
		n.Vocab = n.Vocab[delimEt:]
		newNode.ChildList = []*RadixNode{n}
		return newNode
	}
	return n
}
