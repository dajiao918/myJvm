package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {

	classpath := &Classpath{}
	classpath.parseBootAndExtClasspath(jreOption)
	classpath.parseUserClasspath(cpOption)
	return classpath
}

func (cp *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	//将jre/lib/所有jar文件都读入到
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	cp.bootClasspath = newWildCardEntry(jreLibPath)

	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	cp.extClasspath = newWildCardEntry(jreExtPath)
}

func (cp *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	cp.userClasspath = newEntry(cpOption)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	//如果当前目录下有jre
	if exists("./jre") {
		return "./jre"
	}
	//从环境变量中获取
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return jh + "\\jre"
	}
	panic("can not find the jre folder")
}

func exists(jreOption string) bool {
	//Stat 返会文件的基本信息
	if _, err := os.Stat(jreOption); err != nil {
		//IsNotExist中的参数，ErrNotExist和一些系统调用错误会使它返回真
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (cp Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := cp.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := cp.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return cp.userClasspath.readClass(className)
}
