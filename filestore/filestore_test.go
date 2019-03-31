package filestore

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestInfo(t *testing.T) {
	fs := NewFileStore("./")
	id := "xxx"

	info := &FileInfo{}
	info.ID = id
	info.Offset = 0
	info.Size = 100
	info.Type = ".txt"

	if fs.WriteInfo(id, info) != nil {
		t.Fatal()
	}

	rinfo, err := fs.GetInfo(id)
	if err != nil {
		t.Fatal()
	}
	if !reflect.DeepEqual(rinfo, info) {
		t.Fatal()
	}

	os.Remove(id + ".info")
}

func TestUpload(t *testing.T) {
	fs := NewFileStore("./")
	id := "yyy"

	info := &FileInfo{}
	info.ID = id
	info.Offset = 0
	info.Size = 10
	info.Type = ".jpg"

	if offset, isCompleted, _ := fs.NewUpload(info); offset != 0 || isCompleted {
		t.Fatal()
	}

	if n, isCompleted, _ := fs.WriteChunk(id, 0, strings.NewReader("hello")); n != 5 || isCompleted {
		t.Fatal()
	}

	if offset, isCompleted, _ := fs.NewUpload(info); offset != 5 || isCompleted {
		t.Fatal()
	}

	if n, isCompleted, _ := fs.WriteChunk(id, 5, strings.NewReader("world")); n != 5 || !isCompleted {
		t.Fatal()
	}

	if n, isCompleted, _ := fs.WriteChunk(id, 5, strings.NewReader("abc")); n != 0 || !isCompleted {
		t.Fatal()
	}

	os.Remove(id + ".bin")
	os.Remove(id + ".info")
}
