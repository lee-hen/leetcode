package design_in_memory_file_system

import (
	"sort"
	"strings"
)

type FileSystem struct {
	File
}

func Constructor() FileSystem {
	fs := new(FileSystem)
	root := make(map[string]*File)
	root["/"] = newFile()
	fs.children = root
	return *fs
}

func (fs *FileSystem) Ls(path string) []string {
	node, fileName := fs.findPath(path, false)
	if node == nil {
		return []string{}
	}
	if fileName != "" {
		return []string{fileName}
	}

	pathNames := make([]string, 0)
	for name := range node.children {
		pathNames = append(pathNames, name)
	}

	sort.Strings(pathNames)
	return pathNames
}

func (fs *FileSystem) Mkdir(path string)  {
	fs.add(path, "")
}

func (fs *FileSystem) AddContentToFile(filePath string, content string)  {
	fs.add(filePath, content)
}

func (fs *FileSystem) ReadContentFromFile(filePath string) string {
	_, content := fs.findPath(filePath, true)
	return content
}

type File struct {
	children map[string]*File
	content strings.Builder
}

func newFile() *File {
	return &File{children: map[string]*File{}}
}

func (f *File) add(str string, content string) {
	if str == "/" {
		return
	}

	node := f
	strs := strings.Split(str, "/")

	for j := 0; j < len(strs); j++ {
		letter := strs[j]
		if letter == "" {
			letter = "/"
		}
		if _, found := node.children[letter]; !found {
			node.children[letter] = newFile()
		}
		node = node.children[letter]
	}

	node.content.WriteString(content)
}

func (f *File) getChild(str string) *File {
	return f.children[str]
}

func (f *File) findPath(str string, readContent bool) (*File, string) {
	if str == "/" {
		return f.getChild("/"), ""
	}

	strs := strings.Split(str, "/")
	node := f
	parent := node
	var lastName string
	for i := 0; i < len(strs); i++ {
		letter := strs[i]
		if letter == "" {
			letter = "/"
		}

		lastName = letter
		parent = node
		node = node.getChild(letter)
		if node == nil {
			return nil, ""
		}
	}

	if readContent {
		return node, node.content.String()
	} else {
		if node.content.Len() != 0 {
			return parent, lastName
		}

		return node, ""
	}
}
