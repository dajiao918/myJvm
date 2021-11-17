package api_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAbs(t *testing.T) {
	path, err := filepath.Abs("f:\\hello.class")
	if err != nil {
		t.Log("error occur ", err)
	}
	t.Logf("path is %s\n", path)
}

func TestJoin(t *testing.T) {
	fileName := filepath.Join("F:\\goEnvironment\\goProject\\src\\jvmgo\\api_test", "filepath_test.go")
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Log("err occur: ", err)
	}
	t.Log(string(data))
}

func TestSplit(t *testing.T) {
	separator := string(os.PathListSeparator)
	filename := "f:\\work\\hello.class;f:\\a\\hello.class"
	for _, path := range strings.Split(filename, separator) {
		fmt.Println(path)
	}

}
