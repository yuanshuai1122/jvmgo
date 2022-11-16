package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

/*
*

	new 文件入口
*/
func newDirEntry(path string) *DirEntry {
	// 将path转为绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

/*
*

	读取class内容
*/
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	// 路径和class拼接成完整路径
	fileName := filepath.Join(self.absDir, className)
	// 读取class
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

/*
*

	返回absDir
*/
func (self *DirEntry) String() string {
	// 直接返回
	return self.absDir
}
