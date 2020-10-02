package testutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func Test_CreateFileTree(t *testing.T) {
	root := CreateFileTree(
		Dir("xyz",
			FileAndData("config.json", []byte("text")),
			Dir("loc",
				Dir("eng", File("string.xml")),
				Dir("fre", File("string.xml")),
				Dir("ger", File("xyz")),
				Dir("xxz", File("string.xml")))))

	files, err := ioutil.ReadDir(root)

	if err != nil {
		t.Errorf("Failed to list root: %s\n", err)
	}

	if len(files) != 2 {
		t.Errorf("Root does not contain both children", files)
	}

	data, err2 := ioutil.ReadFile(filepath.Join(root, "config.json"))

	if err2 != nil {
		t.Errorf("Unable to read config data: %s", err2)
	}

	if string(data) != "text" {
		t.Errorf("Wrong data read from config file: %s", string(data))
	}

	files, err = ioutil.ReadDir(filepath.Join(root, "loc"))

	if err != nil {
		t.Errorf("Failed to list loc dir: %s\n", err)
	}

	if len(files) != 4 {
		t.Errorf("Loc does not contain all expected children", files)
	}

	names := map[string]bool{}

	for _, f := range files {
		names[f.Name()] = false
	}

	contains(t, names, "eng", "fre", "ger", "xxz")
	existsFile(t, "eng", "string.xml", root)
	existsFile(t, "fre", "string.xml", root)
	existsFile(t, "ger", "xyz", root)
	existsFile(t, "xxz", "string.xml", root)

}

func existsFile(
	t *testing.T,
	lang, file, root string) {
	if _, err := os.Stat(filepath.Join(root, "loc", lang, file)); os.IsNotExist(err) {
		t.Errorf("'loc%s/%s' does not exist", lang, file)
	}

}
func contains(t *testing.T, names map[string]bool, expected ...string) {
	for _, n := range expected {
		if _, has := names[n]; !has {
			t.Errorf("Missing dir %s: %v", n, names)
		}

	}
}
