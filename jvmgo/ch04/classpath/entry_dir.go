package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	baseDir string
}

func newDirEntry(path string) Entry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{baseDir: absDir}
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.baseDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.baseDir
}
