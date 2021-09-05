package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	// 去掉 *
	baseDir := path[:len(path)-1]
	compositeEntry := CompositeEntry{}
	// 遍历该目录
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过子目录, 使用通配符不会递归查找子目录下的 jar 包
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}

		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}

		return nil
	}

	_ = filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
