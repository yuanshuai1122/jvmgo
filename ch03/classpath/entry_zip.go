package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

/*
*

	new zip压缩包的入口
*/
func newZipEntry(path string) *ZipEntry {
	// 将path转为绝对路径
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}
func (self *ZipEntry) String() string {
	// 返回绝对路径
	return self.absPath
}

/*
*

	从zip文件中提取class文件
*/
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// 打开绝对路径下的zip文件
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	// close zip
	defer r.Close()
	// 遍历zip压缩包里的文件
	for _, f := range r.File {
		// 如果符合class文件 就打开并读取
		if f.Name == className {
			// 打开
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			// close open
			defer rc.Close()
			// 读取
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	// 如果zip没有对应的class文件 返回未找到
	return nil, nil, errors.New("class not found: " + className)
}
