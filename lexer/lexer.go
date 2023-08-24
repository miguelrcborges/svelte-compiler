package lexer

import (
	"os"
	"bytes"

	"errors"
	"fmt"
)

func Parse(fn string) (Node, error) {
	b, err := os.ReadFile(fn)
	l := 1
	if err != nil {
		return nil, err
	}

	body := HTMLElement {
		Name: "body",
	}

	for i := 0; i < len(b); i++ {
		if r, s := isNewline(b, i); r {
			i += s
			l++
			continue
		} else if b[i] == ' ' || b[i] == '\t' {
			continue
		} else if b[i] != '<' {
			return nil, errors.New(fmt.Sprintf("Syntax error in file %s at line %d (char %d). Expected an element.", fn, l, i))
		}

		if i+4 < len(b) && bytes.Equal(b[i:i+4], []byte("<!--")) {
			i, l = skipComment(b, i+4, l)
		} else if i+8 < len(b) && bytes.Equal(b[i:i+8], []byte("<script>")){
			i, l = parseJs(b, i+8, l)
		} else {
			return nil, errors.New(fmt.Sprintf("Unknown element in file %s at line %d (char %d).", fn, l, i))
		}
	}

	return &body, nil
}


func skipComment(b []byte, i, l int) (n_i, n_l int) {
	for ; i < len(b); i++ {
		if r, s := isNewline(b, i); r {
			i += s
			l++
		} else if b[i] == '-' && bytes.Equal(b[i:i+3], []byte("-->")) {
			return i + 2, l
		}
	}
	return i, l
}


func isNewline(b []byte, i int) (r bool, s int){
	if b[i] == '\r' {
		if i+1 < len(b) && b[i+1] == '\n' {
			s = 1
		}
		r = true
	} else if b[i] == '\n' {
		r = true
	} else {
		r = false
	}
	return
}
