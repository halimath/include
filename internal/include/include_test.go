package include

import (
	"strings"
	"testing"
)

const src = `//go:build include

//go:generate go-include --out main_gen.go $GOFILE

package main

import (
	"fmt"
	"net/http"

	"github.com/halimath/include"
)

var (
	html    = include.Bytes("./index.html")
	htmlStr = include.String("./index.html")
)

func main() {
	fmt.Printf("Listening on :8080...\n")

	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(html)
	}))
}`

var want = strings.TrimSpace(`
//Code generated by go-include. DO NOT EDIT.

//go:build !include

//go:generate go-include --out main_gen.go $GOFILE

package main

import (
	"fmt"
	"net/http"
)

var (
	html = []byte{
		60, 33, 68, 79, 67, 84, 89, 80, 69, 32, 104, 116, 109, 108, 62, 10, 60, 104, 116, 109, 108,
		62, 10, 32, 32, 32, 32, 60, 104, 101, 97, 100, 62, 10, 32, 32, 32, 32, 32, 32, 32,
		32, 60, 116, 105, 116, 108, 101, 62, 72, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 32,
		111, 102, 32, 105, 110, 108, 105, 110, 101, 100, 32, 102, 105, 108, 101, 115, 60, 47, 116, 105,
		116, 108, 101, 62, 10, 32, 32, 32, 32, 60, 47, 104, 101, 97, 100, 62, 10, 32, 32, 32,
		32, 60, 98, 111, 100, 121, 62, 10, 32, 32, 32, 32, 32, 32, 32, 32, 60, 104, 49, 62,
		72, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 32, 111, 102, 32, 105, 110, 108, 105, 110,
		101, 100, 32, 102, 105, 108, 101, 115, 60, 47, 104, 49, 62, 10, 32, 32, 32, 32, 60, 47,
		98, 111, 100, 121, 62, 10, 60, 47, 104, 116, 109, 108, 62,
	}
	htmlStr = ` + "`" + `<!DOCTYPE html>
<html>
    <head>
        <title>Hello world of inlined files</title>
    </head>
    <body>
        <h1>Hello world of inlined files</h1>
    </body>
</html>` + "`" + `
)

func main() {
	fmt.Printf("Listening on :8080...\n")

	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(html)
	}))
}
`)

func Test_Package(t *testing.T) {
	gotBytes, err := Include("test.go", []byte(src), Options{
		WorkingDir: "../../example",
	})
	if err != nil {
		t.Fatal(err)
	}

	got := strings.TrimSpace(string(gotBytes))

	if got != want {
		t.Error(got)
	}
}
