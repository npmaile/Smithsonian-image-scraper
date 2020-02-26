package main

import (
	"flag"
	"fmt"
)

var key string
var terms []string

func init() {
	flag.StringVar(&key, "key", "none", "API key for open data access. please read the readme for more info")
	flag.Parse()
	terms = flag.Args()
}

func main() {
	fmt.Println(key)
	fmt.Println(terms)
}
