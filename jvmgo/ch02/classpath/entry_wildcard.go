package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildCardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1]
	compositeEntry := []Entry{}
	walkFun := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			entry := newZipEntry(path)
			compositeEntry = append(compositeEntry, entry)
		}
		return nil
	}
	//Walk 函数递归的遍历baseDir中的目录和文件
	//并且执行将每个文件的路径，文件信息传入到walkFun方法中并且执行此方法
	filepath.Walk(baseDir, walkFun)
	return compositeEntry
}
