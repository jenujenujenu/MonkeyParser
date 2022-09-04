package main

import (
	"fmt"
)

func main() {
	fmt.Println("moin dikka")

	doc := Document{}

	doc.Elements = append(doc.Elements, "Hello")
	doc.Elements = append(doc.Elements, 420)
	
	fmt.Printf("%#v", doc)
}

type Document struct {
	Elements []interface{}
}