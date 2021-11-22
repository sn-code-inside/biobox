package newick

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"text/scanner"
)

type Scanner struct {
	s *bufio.Scanner
}
type Node struct {
	Id                 int
	Child, Sib, Parent *Node
	Label              string
	Length             float64
	HasLength          bool
}

var nodeId = 1

func (s *Scanner) Scan() bool {
	return s.s.Scan()
}
func (s *Scanner) Tree() *Node {
	var root *Node
	var tokens []string
	tree := s.Text()
	tree = strings.ReplaceAll(tree, "[", "/*")
	tree = strings.ReplaceAll(tree, "]", "*/")
	tree = strings.ReplaceAll(tree, "'", "\"")
	tree = strings.ReplaceAll(tree, "\"\"", "'")
	c1 := []rune(tree)
	var c2 []rune
	isNum := false
	for _, r := range c1 {
		if r == ':' {
			isNum = true
			c2 = append(c2, '"')
		}
		if isNum && (r == ',' || r == ';' || r == ' ' || r == ')') {
			isNum = false
			c2 = append(c2, '"')
		}
		c2 = append(c2, r)
	}
	tree = string(c2)
	var tsc scanner.Scanner
	tsc.Init(strings.NewReader(tree))
	for t := tsc.Scan(); t != scanner.EOF; t = tsc.Scan() {
		text := tsc.TokenText()
		if text[0] == '"' {
			var err error
			text, err = strconv.Unquote(text)
			if err != nil {
				log.Fatalf("couldn't unquote %q\n", text)
			}
		} else {
			text = strings.ReplaceAll(text, "_", " ")
		}
		tokens = append(tokens, text)
	}
	i := 0
	v := root
	for i < len(tokens) {
		t := tokens[i]
		if t == "(" {
			if v == nil {
				v = NewNode()
			}
			c := NewNode()
			v.Child = c
			c.Parent = v
			v = v.Child
		}
		if t == ")" {
			v = v.Parent
		}
		if t == "," {
			s := NewNode()
			s.Parent = v.Parent
			v.Sib = s
			v = v.Sib
		}
		if t[0] == ':' {
			l, err := strconv.ParseFloat(t[1:], 64)
			if err != nil {
				log.Fatalf("didn't understand %q\n", t[1:])
			}
			v.Length = l
			v.HasLength = true
		}
		if t == ";" {
			break
		}
		if strings.IndexAny(t[:1], ")(,:;") == -1 {
			v.Label = t
		}
		i++
	}
	root = v
	return root
}
func (s *Scanner) Text() string {
	return s.s.Text()
}

// String turns a tree into its Newick string.
func (n *Node) String() string {
	w := new(bytes.Buffer)
	writeTree(n, w)
	return w.String()
}

// NewScanner returns a scanner for scanning Newick-formatted
// phylogenies.
func NewScanner(r io.Reader) *Scanner {
	sc := new(Scanner)
	sc.s = bufio.NewScanner(r)
	sc.s.Split(scanTrees)
	return sc
}
func scanTrees(data []byte, atEOF bool) (advance int,
	token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, ';'); i >= 0 {
		return i + 1, data[0 : i+1], nil
	}
	if atEOF {
		err := fmt.Errorf("open tree: %q", string(data))
		return len(data), data, err
	}
	return 0, nil, nil
}

// NewNode returns a new node with a unique Id.
func NewNode() *Node {
	n := new(Node)
	n.Id = nodeId
	nodeId++
	return n
}
func writeTree(v *Node, w *bytes.Buffer) {
	if v == nil {
		return
	}
	if v.Parent != nil && v.Parent.Child.Id != v.Id {
		fmt.Fprint(w, ",")
	}
	if v.Child != nil {
		fmt.Fprint(w, "(")
	}
	writeTree(v.Child, w)
	printLabel(w, v)
	writeTree(v.Sib, w)
	if v.Parent != nil && v.Sib == nil {
		fmt.Fprint(w, ")")
	}
	if v.Parent == nil {
		fmt.Fprint(w, ";")
	}
}
func printLabel(w *bytes.Buffer, v *Node) {
	label := v.Label
	if strings.IndexAny(label, "(),") != -1 {
		label = strings.ReplaceAll(label, "'", "''")
		label = fmt.Sprintf("'%s'", label)
	} else {
		label = strings.ReplaceAll(label, " ", "_")
	}
	fmt.Fprintf(w, "%s", label)
	if v.HasLength {
		fmt.Fprintf(w, ":%.3g", v.Length)
	}
}