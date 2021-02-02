package main

import (
	"fmt"

	"github.com/rtanx/caesarcipher/cacipher"
)

func main() {
	str := cacipher.Encode("Twitter21050!", 3)
	fmt.Println(str)
}
