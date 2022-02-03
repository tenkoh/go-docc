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
This is a simple example.

```go
package main

import "github.com/tenkoh/go-docc"

func main(){
    ps, _ := docc.Decode("target.docx")
    for _, p := range ps{
        fmt.Println(p)
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