package mvduplicate

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ReadConfigurationのテスト
func TestReadConfiguration(t *testing.T) {
	dir := ReadConfiguration("../config.json").SearchDir
	assert.Equal(t, "./sample", dir)
}

// getFilepathsのテスト
func TestGetFilepaths(t *testing.T) {
	filepaths := GetFilepaths("../sample/")
	for _, file := range filepaths {
		assert.FileExists(t, file, "")
	}
}

// MvFilesのテスト
func TestMvFiles(t *testing.T) {
	var filepaths []string
	moveDir := "../test"
	filepaths = append(filepaths, "../sample/test1/test1.txt")
	MvFiles(filepaths, moveDir)
	assert.FileExists(t, filepath.Join(moveDir, "test1.txt"))
}
