# docc
Simple ".docx" converter implemented by Go. Convert ".docx" to plain text.

## License
MIT

## Features
- Less dependency.
- No need for Microsoft Office.
- Only on limited environment, also ".doc" could be converted.
  - Windows in which MS Office has been installed.

## Usage

### As a package
This is a simple example to read all paragraphs.

```go
package main

import "github.com/tenkoh/go-docc"

func main(){
    fp := filepath.Clean("./target.docx")
    r, err := NewReader(fp)
    if err != nil {
        panic(err)
    }
    defer r.Close()
    ps, _ := r.ReadAll()
    // do something with ps:[]string
}
```

If you want read the document by a line, the below example is useful.

```go
package main

import "github.com/tenkoh/go-docc"

func main(){
    fp := filepath.Clean("./target.docx")
    r, err := NewReader(fp)
    if err != nil {
        panic(err)
    }
    defer r.Close()
    
    for {
        p, err := r.Read()
        if err == io.EOF {
            return
        } else if err != nil {
            panic(err)
        }
        // do something with p:string
    }
}
```


Before compiling, you shall execute `go mod tidy` to get this package.

### As a binary
`go install` is available.

```shell
go install github.com/tenkoh/go-docc/cmd/docc@latest
```

Then, `docc` command could be used. This is a simple example.

```shell
docc target.docx > plain.txt
```

## Contribution
Your contribution is really welcomed!

## Author
tenkoh