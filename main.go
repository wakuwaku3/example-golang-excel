package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("missing mode")
	}
	mode := os.Args[1]
	var f func() error
	switch mode {
	case "createTemplate":
		f = createTemplate
	case "add":
		f = add
	}
	if err := f(); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s is done", mode)
}
