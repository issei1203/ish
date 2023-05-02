package file

type file struct {
	name    string
	content string
}

type directory struct {
	name      string
	files     []*file
	parentDir *directory
	childDir  []*directory
}

type Content struct {
	IsDirectory bool
	Name        string
}
