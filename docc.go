package docc

import (
	"archive/zip"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var ErrNotSupportFormat = errors.New("the file is not supported")

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

type Paragraph struct {
	R []struct {
		T struct {
			Text  string `xml:",chardata"`
			Space string `xml:"space,attr"`
		} `xml:"t"`
	} `xml:"r"`
}

type Reader struct {
	docxPath string
	fromDoc  bool
	docx     *zip.ReadCloser
	xml      io.ReadCloser
	dec      *xml.Decoder
}

func NewReader(docxPath string) (*Reader, error) {
	r := new(Reader)
	r.docxPath = docxPath
	ext := strings.ToLower(filepath.Ext(docxPath))
	if ext == ".doc" {
		xfp, err := docToX(docxPath)
		if err != nil {
			return nil, fmt.Errorf("fail or not support .doc file: %w", err)
		}
		r.docxPath = xfp
		r.fromDoc = true
	} else if ext != ".docx" {
		return nil, ErrNotSupportFormat
	}

	a, err := zip.OpenReader(r.docxPath)
	if err != nil {
		return nil, err
	}
	r.docx = a

	f, err := a.Open("word/document.xml")
	if err != nil {
		return nil, err
	}
	r.xml = f
	r.dec = xml.NewDecoder(f)

	return r, nil
}

func (r *Reader) Read() (string, error) {
	err := seekNextTag(r.dec, "p")
	if err != nil {
		return "", err
	}
	p, err := seekParagraph(r.dec)
	if err != nil {
		return "", err
	}
	return p, nil
}

func (r *Reader) ReadAll() ([]string, error) {
	ps := []string{}
	for {
		p, err := r.Read()
		if err == io.EOF {
			return ps, nil
		} else if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
}

func (r *Reader) Close() error {
	r.xml.Close()
	r.docx.Close()
	if r.fromDoc {
		os.Remove(r.docxPath)
	}
	return nil
}

func seekParagraph(dec *xml.Decoder) (string, error) {
	var t string
	for {
		token, err := dec.Token()
		if err != nil {
			return "", err
		}
		switch tt := token.(type) {
		case xml.EndElement:
			if tt.Name.Local == "p" {
				return t, nil
			}
		case xml.StartElement:
			if tt.Name.Local == "t" {
				text, err := seekText(dec)
				if err != nil {
					return "", err
				}
				t = t + text
			}
		}
	}
}

func seekText(dec *xml.Decoder) (string, error) {
	for {
		token, err := dec.Token()
		if err != nil {
			return "", err
		}
		switch tt := token.(type) {
		case xml.CharData:
			return string(tt), nil
		case xml.EndElement:
			return "", nil
		}
	}
}

func seekNextTag(dec *xml.Decoder, tag string) error {
	for {
		token, err := dec.Token()
		if err != nil {
			return err
		}
		t, ok := token.(xml.StartElement)
		if !ok {
			continue
		}
		if t.Name.Local != tag {
			continue
		}
		break
	}
	return nil
}
