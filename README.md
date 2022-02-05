## Data structures for autocompletion

Data structures optimized for searching words by tokens

Install
```sh
go get github.com/BON4/vocabTree
```

Useage
```go
tree := NewTree(&RadixNode{})

tree.Add("romane")
tree.Add("romanus")
tree.Add("romulus")
tree.Add("rubens")
tree.Add("ruber")
tree.Add("rubicon")
tree.Add("rubicundus")

tree.Fprint(os.Stdout) //Print tree to stdout

tree.FindMatch("ru", func(s string) {
	fmt.Printf("Match: %s\n", s)
}) /* Output:
    Match: rubens
    Match: ruber
    Match: rubicon
    Match: rubicundus
   */
```

Implemented data structures
* Tire (prefix tree)
* Radix tree

To implement
* DFA (deterministic finite automaton)