package lexer

type Node struct {
	Name string
	Attributes map[string]string
	ChildNodes []*Node
}
