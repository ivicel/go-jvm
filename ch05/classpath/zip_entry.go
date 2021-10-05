package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// 解析 zip 或者 jar 包
type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry{absPath}
}

// 查找和加载 class 文件
// TODO: 优化每次都要打开关闭 zip 文件性能 ch02/classpath/entry_zip2.go
func (e *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(e.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer r.Close()

	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}

			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}

			return data, e, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

// 相当 java 中 toString 方法
func (e *ZipEntry) String() string {
	return e.absPath
}
