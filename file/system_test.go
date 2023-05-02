package file

import (
	"testing"
)

func TestHead_GetContent(t *testing.T) {
	var childDirName = "test2"
	var childFileName = "test3"
	var testChildDir = directory{name: childDirName, files: []*file{}, parentDir: nil, childDir: []*directory{}}
	var testChildFile = file{name: childFileName, content: ""}
	var testRootDir = directory{name: "test", files: []*file{&testChildFile}, parentDir: nil, childDir: []*directory{&testChildDir}}
	var testRootDirPtr = &testRootDir
	var testRootDirDptr = &testRootDirPtr
	var head = Head{dir: testRootDirDptr}

	var content = head.GetContent()

	if len(content) != 2 {
		t.Errorf("expected content length is 1. but was %d", len(content))
	}
}

func TestHead_ChangeDir(t *testing.T) {
	var testChildDir = directory{name: "test2", files: []*file{}, parentDir: nil, childDir: []*directory{}}
	var testRootDir = directory{name: "test", files: []*file{}, parentDir: nil, childDir: []*directory{&testChildDir}}
	var testRootDirPtr = &testRootDir
	var testRootDirDptr = &testRootDirPtr
	var head = Head{dir: testRootDirDptr}

	head.ChangeDir("test2")

	if (*head.dir).name != testChildDir.name {
		t.Errorf("expected current directory name is 'test2'. but was %s", (*head.dir).name)
	}
}

func TestHead_ChangeParentDir(t *testing.T) {
	var testRootDir = directory{name: "test", files: []*file{}, parentDir: nil}
	var testChildDir = directory{name: "test2", files: []*file{}, parentDir: &testRootDir}
	var testChildDirPtr = &testChildDir
	var testRootDirDptr = &testChildDirPtr
	var head = Head{dir: testRootDirDptr}

	head.ChangeParentDir()

	if (*head.dir).name != testRootDir.name {
		t.Errorf("expected current directory name is 'test'. but was %s", (*head.dir).name)
	}
}

func Test_initRootDir(t *testing.T) {
	var result = initRootDir()

	if result.name != "root" {
		t.Errorf("expected name is 'root'. but was %s", result.name)
	}

	if len(result.files) != 0 {
		t.Errorf("expected files length is 0. but was %d", len(result.files))
	}

	if result.parentDir != nil {
		t.Errorf("expected parentDir is nil")
	}

	if len(result.childDir) != 0 {
		t.Errorf("expected childDir length is 0. but was %d", len(result.childDir))
	}
}

func Test_makeDir(t *testing.T) {
	var testRootDir = directory{name: "test", files: []*file{}, parentDir: nil, childDir: []*directory{}}

	makeDir(&testRootDir, "test2")

	if len(testRootDir.childDir) != 1 {
		t.Errorf("expected test root childDir length is 1. but was %d", len(testRootDir.childDir))
	}

	var testChildDir = *(testRootDir.childDir[0])

	if testChildDir.name != "test2" {
		t.Errorf("expected testChildDir name is 'test2'. but was %s", testChildDir.name)
	}

	if len(testChildDir.files) != 0 {
		t.Errorf("expected testChildDir files length is 0. but was %d", len(testChildDir.files))
	}

	if testChildDir.parentDir != &testRootDir {
		t.Errorf("expected testParentDir is %p. but was %p", &testRootDir, testChildDir.parentDir)
	}

	if len(testChildDir.childDir) != 0 {
		t.Errorf("expected testChildDir childDir length is 0. but was %d", len(testChildDir.childDir))
	}
}

func Test_removeDir(t *testing.T) {
	var childDirName = "test2"
	var testChildDir = directory{name: childDirName, files: []*file{}, parentDir: nil, childDir: []*directory{}}
	var testRootDir = directory{name: "test", files: []*file{}, parentDir: nil, childDir: []*directory{&testChildDir}}

	removeDir(&testRootDir, childDirName)

	if len(testRootDir.childDir) != 0 {
		t.Errorf("expected testRootDir childDir length is 0. but was %d", len(testChildDir.childDir))
	}
}

func Test_makeFile(t *testing.T) {
	var testRootDir = directory{name: "test", files: []*file{}, childDir: []*directory{}}

	makeFile(&testRootDir, "test2", "test")

	if len(testRootDir.files) != 1 {
		t.Errorf("expected testRootDir files length is 1. but was %d", len(testRootDir.files))
	}

	var testChildFile = *(testRootDir.files[0])

	if testChildFile.name != "test2" {
		t.Errorf("expected testChildFile name is test2. but was %s", testChildFile.name)
	}

	if testChildFile.content != "test" {
		t.Errorf("expected testChildFile content is ''. but was %s", testChildFile.content)
	}
}

func Test_removeFile(t *testing.T) {
	var childFileName = "test2"
	var testChildFile = file{name: childFileName, content: ""}
	var testRootDir = directory{name: "test", files: []*file{&testChildFile}, childDir: []*directory{}}

	removeFile(&testRootDir, childFileName)

	if len(testRootDir.files) != 0 {
		t.Errorf("expected testRootDir files length is 0. but was %d", len(testRootDir.files))
	}
}
