package mvduplicate

import (
	"encoding/json"
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
		panic(err)
	}
	var config Configuration
	json.Unmarshal(raw, &config)
	config.Directory = filepath.FromSlash(config.Directory)
	return &config
}

// GetFilepaths 指定するディレクトリ以下のファイル一覧をスライスで返却（サブフォルダも再帰検索）
func GetFilepaths(searchPath string) []string {
	files, err := ioutil.ReadDir(searchPath)
	if err != nil {
		panic(err)
	}
	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, GetFilepaths(filepath.Join(searchPath, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(searchPath, file.Name()))
	}
	return paths
}
