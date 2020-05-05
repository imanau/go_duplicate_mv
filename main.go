package main

import (
	"duplicate_delete/mvduplicate"
	"fmt"
)

func main() {
	searchPath := mvduplicate.ReadConfiguration("./config.json").Directory
	fmt.Println(mvduplicate.GetFilepaths(searchPath))
}
