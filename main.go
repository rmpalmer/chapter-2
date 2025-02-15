package main

import (
	"fmt"
	"log"
	"os"

	_ "chapter-2/matchers"
	"chapter-2/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	fmt.Println("Welcome to chapter 2")
	search.Run("president")
	fmt.Println("Goodbye")
}
