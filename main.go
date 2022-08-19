package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/moby/moby/pkg/namesgenerator"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	name := namesgenerator.GetRandomName(0)
	name = strings.Replace(
		strings.Title(
			strings.Replace(
				name, "_", " ", -1,
			),
		), " ", "", -1,
	)

	fmt.Println(name)
}
