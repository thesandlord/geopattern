package main

import (
	"fmt"
	"github.com/pravj/geopattern"
)

// Prints pattern's SVG string for a phrase argument
func main() {
	args := geopattern.Pattern{
		Phrase: "O",
	}
	gp := geopattern.Generate(args)
	fmt.Println(gp)
}
