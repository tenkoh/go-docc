package docc

import (
	"path/filepath"
	"reflect"
	"testing"
)

var want = []string{
	"Title",
	"Subtitle",
	"Here is a first row.",
	"Here is a second row.",
}

func TestReadAll(t *testing.T) {
	fp := filepath.Clean("./testdata/test.docx")
	r, err := NewReader(fp)
	if err != nil {
		panic(err)
	}
	defer r.Close()
	got, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestRead(t *testing.T) {
	fp := filepath.Clean("./testdata/test.docx")
	r, err := NewReader(fp)
	if err != nil {
		panic(err)
	}
	defer r.Close()
	got, err := r.Read()
	if err != nil {
		panic(err)
	}
	if want[0] != got {
		t.Errorf("want %s, got %s", want[0], got)
	}
}
