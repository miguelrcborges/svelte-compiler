package main

import (
	"fmt"
	"os"
	"sync"

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
		os.Exit(1)
	}

	err = os.MkdirAll(os.Args[2], 0755)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		path := os.Args[2] + "/index.html"
		html, err := os.Create(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create %s.\n", path)
			return
		}
		defer html.Close()

		_, err = html.WriteString(tree.RenderHTML())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write to %s.\n", path)
			return
		}
	}()
	go func() {
		defer wg.Done()
		path := os.Args[2] + "/style.css"
		html, err := os.Create(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create %s.\n", path)
			return
		}
		defer html.Close()

		_, err = html.WriteString(tree.RenderCSS())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write to %s.\n", path)
			return
		}
	}()
	go func() {
		defer wg.Done()
		path := os.Args[2] + "/main.js"
		html, err := os.Create(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create %s.\n", path)
			return
		}
		defer html.Close()

		_, err = html.WriteString(tree.RenderJS())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write to %s.\n", path)
			return
		}
	}()
	wg.Wait()
}

