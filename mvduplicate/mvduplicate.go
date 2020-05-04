package mvduplicate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// Configuration 設定ファイル読み込み用のstruct
type Configuration struct {
	Directory string `json:"directory"`
}

// ReadConfiguration configファイルを読み込む
func ReadConfiguration(configPath string) *Configuration {
	raw, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println(err.Error())

	}
	var config Configuration
	json.Unmarshal(raw, &config)
	config.Directory = filepath.FromSlash(config.Directory)
	return &config
}
