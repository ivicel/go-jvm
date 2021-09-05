package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// 目录形式的类路径
type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &DirEntry{absDir}
}

// 查找和加载 class 文件
func (d *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(d.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, d, err
}

// 相当 java 中 toString 方法
func (d *DirEntry) String() string {
	return d.absDir
}
