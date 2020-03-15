package main

import (
    "io/ioutil"
    "os"
    "fmt"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
	"github.com/yuin/goldmark/text"
    "github.com/lobre/mdfmt/renderer/markdown"
)

func main() {
    if len(os.Args) <= 1 {
        fmt.Println("missing filename as argument")
        os.Exit(1)
    }

    b, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    r := renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(markdown.NewRenderer(), 1100))) 
	md := goldmark.New(
		goldmark.WithRenderer(r),
	)
    p := md.Parser()
    reader := text.NewReader(b)
    node := p.Parse(reader)
    node.Dump(b, 0)
}
