package mvduplicate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ReadConfigurationのテスト
func TestReadConfiguration(t *testing.T) {
	dir := ReadConfiguration("../config.json").Directory
	assert.Equal(t, "./sample", dir)
}

// getFilepathsのテスト
func TestGetFilepaths(t *testing.T) {
	// filepaths := GetFilePaths("../sample/")
	var filepaths []string
	filepaths = append(filepaths, "../sample/test1/test1.txt")
	assert.FileExists(t, filepaths[0], "")
}
