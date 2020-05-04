package main

import (
	"duplicate_delete/mvduplicate"
	"fmt"
)

func main() {
	fmt.Println(mvduplicate.ReadConfiguration("./config.json").Directory)
}
