package main

import (
	"fmt"
	"strings"
	petname "github.com/dustinkirkland/golang-petname"
)

func main() {
	petname.NonDeterministicMode()
	name := petname.Generate(2, "-")
	name = strings.Title(name)
	name = strings.Replace(name, "-", "", -1)
	fmt.Println(name)
}
