package file

type file struct {
	name    string
	content string
}

type directory struct {
	name        string
	files       []*file
	directories []*directory
}
