package file

import (
	"testing"
)

func Test_initRootDir(t *testing.T) {
	var result = initRootDir()

	if result.name != "root" {
		t.Errorf("expected name is 'root'. but was %s", result.name)
	}

	if len(result.files) != 0 {
		t.Errorf("expected files length is 0. but was %d", len(result.files))
	}

	if len(result.directories) != 0 {
		t.Errorf("expected directories length is 0. but was %d", len(result.directories))
	}
}

func Test_makeDir(t *testing.T) {
	var testRootDir = directory{name: "test", files: []*file{}, directories: []*directory{}}

	makeDir(&testRootDir, "test2")

	if len(testRootDir.directories) != 1 {
		t.Errorf("expected test root directories length is 1. but was %d", len(testRootDir.directories))
	}

	var testChildDir = *(testRootDir.directories[0])

	if testChildDir.name != "test2" {
		t.Errorf("expected testChildDir name is 'test2'. but was %s", testChildDir.name)
	}

	if len(testChildDir.files) != 0 {
		t.Errorf("expected testChildDir files length is 0. but was %d", len(testChildDir.files))
	}

	if len(testChildDir.directories) != 0 {
		t.Errorf("expected testChildDir directories length is 0. but was %d", len(testChildDir.directories))
	}
}

func Test_removeDir(t *testing.T) {
	var childDirName = "test2"
	var testChildDir = directory{name: childDirName, files: []*file{}, directories: []*directory{}}
	var testRootDir = directory{name: "test", files: []*file{}, directories: []*directory{&testChildDir}}

	removeDir(&testRootDir, childDirName)

	if len(testRootDir.directories) != 0 {
		t.Errorf("expected testRootDir directories length is 0. but was %d", len(testChildDir.directories))
	}
}
