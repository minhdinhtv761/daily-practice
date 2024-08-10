package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		log.Fatal("no arguments found")
	}
	switch args[1] {
	case "hello world":
		fmt.Println("Hi minh")
	}
	return nil
}
