package file

var rootDir directory

const root = "root"

type Head struct {
	dir *directory
}

func (head Head) GetContent() []Content {
	var contents []Content
	for _, dir := range head.dir.childDir {
		contents = append(contents, Content{IsDirectory: true, Name: dir.name})
	}

	for _, file := range head.dir.files {
		contents = append(contents, Content{IsDirectory: false, Name: file.name})
	}

	return contents
}

func (head Head) MakeDir(name string) {
	makeDir(head.dir, name)
}

func (head Head) RemoveDir(name string) {
	removeDir(head.dir, name)
}

func (head Head) MakeFile(name, content string) {
	makeFile(head.dir, name, content)
}

func (head Head) RemoveFile(name string) {
	removeFile(head.dir, name)
}

func InitHead() Head {
	var rootDir = initRootDir()
	return Head{dir: &rootDir}
}

func initRootDir() directory {
	rootDir = directory{name: root, files: []*file{}, parentDir: nil, childDir: []*directory{}}
	return rootDir
}

func makeDir(parent *directory, name string) {
	var child = directory{name: name, files: []*file{}, parentDir: parent, childDir: []*directory{}}
	var newChildren = append(parent.childDir, &child)
	parent.childDir = newChildren
}

func removeDir(parent *directory, name string) {
	var children = parent.childDir

	var targetIndex int = -1

	for i := 0; i < len(children); i++ {
		var child = *(children[i])
		if child.name == name {
			targetIndex = i
			break
		}
	}

	if targetIndex < 0 {
		return
	}

	var newChildren []*directory
	for i := 0; i < len(children); i++ {
		if i != targetIndex {
			newChildren = append(newChildren, children[i])
		}
	}

	parent.childDir = newChildren
}

func makeFile(parent *directory, name string, content string) {
	var child = file{name: name, content: content}
	var newChildren = append(parent.files, &child)
	parent.files = newChildren
}

func removeFile(parent *directory, name string) {
	var children = parent.files

	var targetIndex int = -1

	for i := 0; i < len(children); i++ {
		var child = *(children[i])
		if child.name == name {
			targetIndex = i
			break
		}
	}

	if targetIndex < 0 {
		return
	}

	var newChildren []*file
	for i := 0; i < len(children); i++ {
		if i != targetIndex {
			newChildren = append(newChildren, children[i])
		}
	}

	parent.files = newChildren
}
