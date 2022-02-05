package fileReader

import (
	"bufio"
	"errors"
	"io"
	"regexp"
	"strings"
	"vocabTree/trees"
)

type Language int

const (
	Golang Language = iota
)

var m = map[string]struct{}{}

func discardRepeat(s []string) []string {
	for i, w := range s {
		if _, ok := m[w]; ok {
			s[i] = ""
		} else {
			m[w] = struct{}{}
		}
	}
	return s
}

func appendFromGolangFile(root trees.Tree, scanner *bufio.Scanner) error {
	var validGolangWord = regexp.MustCompile(`[^\n()*, {}_.:=;(")><\-+(0-9)\]\[\t&^\\\/!]*`)

	for scanner.Scan() {
		s := scanner.Text()
		if !strings.HasPrefix(s, "//") {
			words := discardRepeat(validGolangWord.FindAllString(s, -1))
			for _, word := range words {
				if len(word) > 1 {
					root.Add(word)
				}
			}
		}
	}
	return nil
}

func ConstructTreeFromFile(nodeType trees.Node, r io.Reader, l Language) (trees.Tree, error) {
	var root = trees.NewTree(nodeType)

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	switch l {
	case Golang:
		return root, appendFromGolangFile(root, s)
	default:
		return root, errors.New("provided language not implemented")
	}
}

func AppendFromFile(root trees.Tree, r io.Reader, l Language) error {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	switch l {
	case Golang:
		return appendFromGolangFile(root, s)
	default:
		return errors.New("provided language not implemented")
	}
}
