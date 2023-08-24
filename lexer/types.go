package lexer

import (
	"fmt"
	"strings"
)

type ChildrenNodes []*Node

type CustomElement struct {
	Name string
	Attributes map[string]string
	Children ChildrenNodes
	JS []JSStatement
	CSS map[string]string
}

type HTMLElement struct {
	Name string
	Attributes map[string]string
	Children ChildrenNodes
	Dependencies []string
}

type JSStatement struct {

}

type Node interface {
	RenderHTML() string
	RenderJS() string
	RenderCSS() string
}

func (e *CustomElement) RenderHTML() string {
	return e.renderHTMLStart() + e.Children.renderChildrenHTML() + e.renderHTMLEnd()
}

func (e *CustomElement) renderHTMLStart() string {
	s := "<c-" + strings.ToLower(e.Name)
	for k, v := range e.Attributes {
		s += fmt.Sprint(" \"%s\"=\"%s\"", k, v)
	}
	s += ">"
	return s
}

func (e *CustomElement) renderHTMLEnd() string {
	return fmt.Sprintf("</c-%s>", strings.ToLower(e.Name))
}

func (e *CustomElement) RenderCSS() string {
	s := ""
	for k, v := range e.CSS {
		s += fmt.Sprintf("c-%s %s{%s}", strings.ToLower(e.Name), k, v)
	}
	return s + e.Children.renderChildrenJS()
}

func (e *CustomElement) RenderJS() string {
	s := fmt.Sprintf("" + 
		"class %s extends HTMLElement{" +
		"};" +
		"customElements.define(\"c-%s\", %s);",
		e.Name, strings.ToLower(e.Name), e.Name)
	return s + e.Children.renderChildrenJS()
}

func (e *HTMLElement) RenderHTML() string {
	return e.renderHTMLStart() + e.Children.renderChildrenHTML() + e.renderHTMLEnd()
}

func (e *HTMLElement) renderHTMLStart() string {
	s := "<" + e.Name
	for k, v := range e.Attributes {
		s += fmt.Sprintf(" \"%s\"=\"%s\"", k, v)
	}
	s += ">"
	return s
}

func (e *HTMLElement) renderHTMLEnd() string {
	return fmt.Sprintf("</%s>", e.Name)
}

func (e *HTMLElement) RenderJS() string {
	return "" + e.Children.renderChildrenJS()
}

func (e *HTMLElement) RenderCSS() string {
	return "" + e.Children.renderChildrenCSS()
}

func (e *ChildrenNodes) renderChildrenHTML() string {
	slc := []*Node(*e)
	s := ""
	for i := 0; i < len(slc); i++ {
		t := slc[i]
		s += (*t).RenderHTML()
	}
	return s
}

func (e *ChildrenNodes) renderChildrenCSS() string {
	slc := []*Node(*e)
	s := ""
	for i := 0; i < len(slc); i++ {
		t := slc[i]
		s += (*t).RenderCSS()
	}
	return s
}

func (e *ChildrenNodes) renderChildrenJS() string {
	slc := []*Node(*e)
	s := ""
	for i := 0; i < len(slc); i++ {
		t := slc[i]
		s += (*t).RenderJS()
	}
	return s
}
