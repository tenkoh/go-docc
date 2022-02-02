package docc

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"path/filepath"
)

var (
	ErrDocumentsNotFound = errors.New("the docx file does not contain word/document.xml")
)

type Document struct {
	XMLName xml.Name `xml:"document"`
	Body    struct {
		P []struct {
			R []struct {
				T struct {
					Text  string `xml:",chardata"`
					Space string `xml:"space,attr"`
				} `xml:"t"`
			} `xml:"r"`
		} `xml:"p"`
	} `xml:"body"`
}

// func Decode(fp string) ([]string, error){
// 	ext := strings.ToLower(filepath.Ext(fp))
// 	if ext == ".docx"{

// 	}
// }

func extractXML(fp string) (io.Reader, error) {
	archive, err := zip.OpenReader(fp)
	if err != nil {
		return nil, err
	}
	defer archive.Close()

	for _, f := range archive.File {
		target := filepath.Clean("word/document.xml")
		if n := filepath.Clean(f.Name); n != target {
			continue
		}

		fd, err := f.Open()
		if err != nil {
			return nil, err
		}
		defer fd.Close()

		b := bytes.NewBuffer(nil)
		if _, err := io.Copy(b, fd); err != nil {
			return nil, err
		}
		return b, nil
	}

	return nil, ErrDocumentsNotFound
}

func decodeXML(r io.Reader) ([]string, error) {
	doc := new(Document)
	if err := xml.NewDecoder(r).Decode(doc); err != nil {
		return nil, fmt.Errorf("could not decode the document: %w", err)
	}
	ps := []string{}
	for _, p := range doc.Body.P {
		t := ""
		for _, r := range p.R {
			t = t + r.T.Text
		}
		ps = append(ps, t)
	}
	return ps, nil
}
