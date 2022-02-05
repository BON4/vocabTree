package trees

type TireNode struct {
	Vocab     string
	Comp      bool
	ChildList []*TireNode
}

func NewTire() Node {
	return &TireNode{
		Vocab:     "",
		Comp:      false,
		ChildList: []*TireNode{},
	}
}

func newTire(word string, comp bool, clist []*TireNode) *TireNode {
	if clist == nil {
		return &TireNode{
			Vocab:     word,
			Comp:      comp,
			ChildList: []*TireNode{},
		}
	}
	return &TireNode{
		Vocab:     word,
		Comp:      comp,
		ChildList: clist,
	}
}

func (v *TireNode) Val() string {
	return v.Vocab
}

func (v *TireNode) Complete() bool {
	return v.Comp
}

func (v *TireNode) Len() int {
	return len(v.ChildList)
}

func (v *TireNode) ChildMap(f func(n Node)) {
	for i := 0; i < len(v.ChildList); i++ {
		f(v.ChildList[i])
	}
}

//newBranch - creates new brach from word:
//word := "test" -> Branch["t" -> "e" -> "s" -> "t"]
func newBranch(word string) *TireNode {
	if len(word) == 1 {
		return newTire(word, true, nil)
	}
	return newTire(word[:1], false, []*TireNode{newBranch(word[1:])})
}

func (r *TireNode) Add(word string) {
	if len(word) > 0 {
		var foundNode = false

		for _, c := range r.ChildList {
			if c.Vocab == word[:1] {
				foundNode = true
				c.Add(word[1:])
			}
		}

		if !foundNode {
			var newNode = newBranch(word)
			r.ChildList = append(r.ChildList, newNode)
		}
	}
}
