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
	filepaths := GetFilepaths("../sample/")
	for _, file := range filepaths {
		assert.FileExists(t, file, "")
	}
}
