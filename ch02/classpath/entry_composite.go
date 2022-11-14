package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

/*
newCompositeEntry
*/
func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	// 遍历, 按照分隔符分成小路径，转换成新的Entry实例
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

/*
读取class
*/
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	// 依次调用每一个子路径的readClass（）方法，如果成功读取到class数据
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

/*
*

	String方法
*/
func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	// 遍历子Entry的string
	for i, entry := range self {
		strs[i] = entry.String()
	}
	// 用分隔符拼接返回
	return strings.Join(strs, pathListSeparator)
}
