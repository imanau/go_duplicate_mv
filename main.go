package main

import (
	"duplicate_delete/mvduplicate"
)

func main() {
	config := mvduplicate.ReadConfiguration("./config.json")
	searchPath := config.SearchDir
	moveDir := config.MoveDir
	filepaths := mvduplicate.GetFilepaths(searchPath)
	mvduplicate.MvFiles(filepaths, moveDir)
}
