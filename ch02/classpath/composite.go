package classpath

import (
	"errors"
	"strings"
)

// 混合 Entry
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := CompositeEntry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}

	return compositeEntry
}

// 查找和加载 class 文件
func (e CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range e {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

// 相当 java 中 toString 方法
func (e CompositeEntry) String() string {
	strs := make([]string, len(e))
	for _, entry := range e {
		strs = append(strs, entry.String())
	}

	return strings.Join(strs, pathListSeparator)
}
