package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	baseDir string
}

func newZipEntry(path string) Entry {
	baseDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{baseDir: baseDir}
}

func (z *ZipEntry) readClass(className string) ([]byte, Entry, error) {

	reader, err := zip.OpenReader(z.baseDir)
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	//打开的reader中有一个File []*file属性，遍历file，如果file的name和classname相等
	for _, f := range reader.File {
		if f.Name == className {
			//就打开这个file
			file, err := f.Open()
			if err != nil {
				panic(err)
			}
			defer file.Close()
			//读取数据
			data, e := ioutil.ReadAll(file)
			return data, z, e
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (z *ZipEntry) String() string {
	return z.baseDir
}
