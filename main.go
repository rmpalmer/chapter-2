package main

import (
	"fmt"
	"log"
	"os"

	_ "matchers"
	"search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	fmt.Println("Welcome to chapter 2")
	search.Run("president")
}
