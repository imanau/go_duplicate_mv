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
