package testutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

type Node struct {
	Name     string
	Data     []byte
	file     bool
	Children []Node
}

func Dir(name string, children ...Node) Node {
	return Node{name, []byte{}, false, children}
}
func File(name string) Node {
	return Node{name, []byte{}, true, []Node{}}
}
func FileAndData(name string, data []byte) Node {
	return Node{name, data, true, []Node{}}
}

func AssertEquals(msg, expected, actual string, t *testing.T) {
	if actual != expected {
		t.Errorf("%s: Expected/Actual \n%q\n%q", msg, expected, actual)
	}
}
func AssertEqualsInt(msg string, expected, actual int, t *testing.T) {
	if actual != expected {
		t.Errorf("%s: Expected/Actual \n%q\n%q", msg, expected, actual)
	}
}

func CreateFileTree(tree Node) string {
	if !tree.file {
		dir, err := ioutil.TempDir("", tree.Name)
		if err != nil {
			panic("Unable to create temporary directory")
		}
		tree.Name = filepath.Base(dir)
		for _, n := range tree.Children {
			createFileTreeInternal(n, dir)
		}
		return dir
	}

	file, err := ioutil.TempFile("", tree.Name)
	if err != nil {
		panic("Unable to create temporary file")
	}
	if err := ioutil.WriteFile(file.Name(), tree.Data, 644); err != nil {
		panic("Unable to create data to temporary file")
	}

	tree.Name = filepath.Base(file.Name())
	return file.Name()
}

func createFileTreeInternal(tree Node, parent string) {
	if !tree.file {
		path := filepath.Join(parent, tree.Name)
		if err := os.Mkdir(path, 744); err != nil {
			panic("Failed to create directory: " + path + "\n" + err.Error())
		}
		for _, n := range tree.Children {
			createFileTreeInternal(n, path)
		}
	} else {
		path := filepath.Join(parent, tree.Name)
		if file, err := os.Create(path); err != nil {
			panic("Failed to create file: " + path + "\n" + err.Error())
		} else {
			defer file.Close()
			file.Write(tree.Data)
		}
	}
}
