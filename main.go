package main

import (
	"fmt"
	"os"

	"github.com/miguelrcborges/svelte-compiler/lexer"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("svelte-compiler\nUsage:\n\tsvelte-compiler <entry_point>.svelte <output_dir>")
		os.Exit(1)
	}

	tree, err := lexer.Parse(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(tree)
	}
}
