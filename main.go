package main

import "fmt"

var stackTmpl = `
package stack

type Node struct {
	value {{.Type}}
	next  *Node
}

type {{.Name}} struct {
	top *Node
}

func (s *{{.Name}}) IsEmpty() bool {
	return s.top == nil
}

func (s *{{.Name}}) Push(v {{.Type}}) {
    s.top = &Node{v, s.top}
}

func (s *{{.Name}}) Pop() ({{.Type}}, bool) {
	if s.IsEmpty() {
		return {{.DefaultValue}}, false
	}

	v := s.top.value
	s.top = s.top.next

	return v, true
}

func (s *{{.Name}}) Peek() ({{.Type}}, bool) {
	if s.IsEmpty() {
		return {{.DefaultValue}}, false
	}

	return s.top.value, true
}
`

type data struct {
	Name         string
	Type         string
	DefaultValue string
}

func main() {
	//d := data{Name: "IntStack", Type: "int", DefaultValue: "0"}
	//t := template.Must(template.New("stack").Parse(stackTmpl))
	//err := t.Execute(os.Stdout, d)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//s := &stack.IntStack{}
	//s.Push(5)
	//s.Push(7)
	//fmt.Println(s.Pop())

	m := make(map[int][]int)
	fmt.Println(m)
	for i := 0; i < 10; i++ {
		if _, ok := m[8]; !ok {
			m[8] = make([]int, 6)
			fmt.Println(m[8])
		}
		m[8] = append(m[8], i)
	}
	fmt.Println(m)
}
