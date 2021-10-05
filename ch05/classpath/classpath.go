package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry // 启动类路径, jdk 里 lib 目录
	extClasspath  Entry // 扩展类路径, ext 目录
	userClasspath Entry // 用户类的目录
}

// Parse 解析 -Xjre 选项的指定的启动类和扩展类路径,
// -claspath/-cp 指定的用户类路径
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)

	return cp
}

// 查找和加载 class 文件
func (c *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := c.bootClasspath.readClass(className); err == nil {
		return data, entry, nil
	}

	if data, entry, err := c.extClasspath.readClass(className); err == nil {
		return data, entry, nil
	}

	return c.userClasspath.readClass(className)
}

// 相当 java 中 toString 方法
func (c *Classpath) String() string {
	return c.userClasspath.String()
}

// parseBootAndExtClasspath 解析启动类和扩展类路径
func (c *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	c.bootClasspath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	c.extClasspath = newWildcardEntry(jreExtPath)
}

// parseUserClasspath 解析用户类路径, 如果没传入使用当前目录
func (c *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}

	c.userClasspath = newEntry(cpOption)
}

// getJreDir 查找 jre 目录
func getJreDir(jreOption string) string {
	// 如果特定传入路径, 查找
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	// 当前目录下的 jre 目录
	if exists("./jre") {
		return "./jre"
	}

	// 查找 java home
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("Cannot find jre folder!")
}

// exists 判断目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}
