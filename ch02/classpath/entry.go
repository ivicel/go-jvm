package classpath

import (
	"os"
	"strings"
)

// 系统分隔符
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// 查找和加载 class 文件
	readClass(className string) ([]byte, Entry, error)
	// 相当 java 中 toString 方法
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		// 混合类型的 classpath 路径, 使用系统(win->分号, linux->冒号)分隔
		return newCompositeEntry(path)
	} else if strings.HasSuffix(path, "*") {
		// 通配符指定某个目录下所有的 jar 包
		return newWildcardEntry(path)
	} else if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasPrefix(path, ".ZIP") {
		// jar 包和 zip 包
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
