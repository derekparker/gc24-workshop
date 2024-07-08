package main

import (
	"bufio"
	"fmt"
	"github.com/derekparker/trie"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("must provide prefix word")
	}
	t := trie.New()
	addFromFile(t, "/usr/share/dict/words")
	results := t.PrefixSearch(os.Args[1])
	fmt.Println("Found:")
	for _, result := range results {
		fmt.Println(result)
	}
}

func addFromFile(t *trie.Trie, path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewScanner(file)

	i := 0
	for reader.Scan() {
		i++
		if i > 100 {
			break
		}
		t.Add(reader.Text(), nil)
	}

	if reader.Err() != nil {
		log.Fatal(err)
	}
}
