package design_in_memory_file_system

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFileSystem(t *testing.T) {
	fs := Constructor()
	require.Equal(t, []string{}, fs.Ls("/"))
	fs.Mkdir("/a/b/c")
	fs.AddContentToFile("/a/b/c/d", "hello")
	require.Equal(t, []string{"a"}, fs.Ls("/"))
	require.Equal(t, "hello", fs.ReadContentFromFile("/a/b/c/d"))
}

func TestFileSystem1(t *testing.T) {
	fs := Constructor()
	fs.Mkdir("/goowmfn")
	fs.Ls("/goowmfn")
	fs.Ls("/")
	fs.Ls("/z")
	fs.Ls("/")
	fs.Ls("/")
	fs.AddContentToFile("/goowmfn/c","shetopcy")

	fs.Ls("z")
	require.Equal(t, []string{"c"}, fs.Ls("/goowmfn/c"))
	fs.Ls("/goowmfn")
}

func TestFileSystem2(t *testing.T) {
	fs := Constructor()
	fs.Mkdir("/m")
	fs.Ls("/m")
	fs.Mkdir("/w")

	fs.Ls("/")
	fs.Ls("/w")
	fs.Ls("/")
	fs.AddContentToFile("/dycete","emer")
	fs.Ls("/")
	fs.Ls("/")
	require.Equal(t, []string{"dycete"}, fs.Ls("/dycete"))
}

func TestFileSystem3(t *testing.T) {
	fs := Constructor()
	fs.Ls("/")
	fs.Mkdir("/a/b/c")
	fs.AddContentToFile("/a/b/c/d","hello world")
	fs.Ls("/")
	fs.ReadContentFromFile("/a/b/c/d")
	fs.AddContentToFile("/a/b/c/d"," hello hello world")
	require.Equal(t, "hello world hello hello world", fs.ReadContentFromFile("/a/b/c/d"))
}

func TestFileSystem4(t *testing.T) {
	fs := Constructor()
	fs.Ls("/")
	fs.Mkdir("/a/b/c")
	fs.Ls("/")
	fs.Mkdir("/a/b")
	fs.Mkdir("/a/b/a")
	require.Equal(t, []string{"a", "c"}, fs.Ls("/a/b"))
}
